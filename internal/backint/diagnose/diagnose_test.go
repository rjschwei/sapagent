/*
Copyright 2023 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package diagnose

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"testing"
	"time"

	store "cloud.google.com/go/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/fsouza/fake-gcs-server/fakestorage"
	"github.com/GoogleCloudPlatform/sapagent/internal/configuration"
	bpb "github.com/GoogleCloudPlatform/sapagent/protos/backint"
)

var (
	fakeServer = fakestorage.NewServer([]fakestorage.Object{
		{
			ObjectAttrs: fakestorage.ObjectAttrs{
				BucketName: "test-bucket",
				Name:       "object.txt",
			},
			Content: []byte("test content"),
		},
		{
			ObjectAttrs: fakestorage.ObjectAttrs{
				BucketName: "test-bucket",
				// The backup object name is in the format <userID>/<fileName>/<externalBackupID>.bak
				Name:    "test@TST/object.txt/12345.bak",
				Created: time.UnixMilli(12345),
			},
			Content: []byte("test content"),
		},
		{
			ObjectAttrs: fakestorage.ObjectAttrs{
				BucketName: "test-bucket",
				Name:       "test@TST/object2.txt/12345.bak",
			},
			Content: []byte("test content"),
		},
	})
	defaultBucketHandle = fakeServer.Client().Bucket("test-bucket")
	defaultConfig       = &bpb.BackintConfiguration{UserId: "test@TST"}
	defaultOptions      = diagnoseOptions{
		execute: func(ctx context.Context, config *bpb.BackintConfiguration, bucketHandle *store.BucketHandle, input io.Reader, output io.Writer) bool {
			return false
		},
		config:       defaultConfig,
		bucketHandle: defaultBucketHandle,
		output:       bytes.NewBufferString(""),
		files: []*diagnoseFile{
			{
				fileSize:         12345,
				fileName:         "object.txt",
				externalBackupID: "12345",
			},
		},
	}
	defaultRemoveFunc = func(name string) error { return nil }
)

// defaultExecute utilizes a closure to write the input line by line
// to the output. execute() can then be called multiple times per test case.
func defaultExecute(lines []string) func(ctx context.Context, config *bpb.BackintConfiguration, bucketHandle *store.BucketHandle, input io.Reader, output io.Writer) bool {
	return func(ctx context.Context, config *bpb.BackintConfiguration, bucketHandle *store.BucketHandle, input io.Reader, output io.Writer) bool {
		if len(lines) == 0 {
			return false
		}
		output.Write([]byte(lines[0]))
		lines = lines[1:]
		return true
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name         string
		bucketHandle *store.BucketHandle
		want         bool
	}{
		{
			name:         "FailNoBucket",
			bucketHandle: nil,
			want:         false,
		},
		{
			name:         "Success",
			bucketHandle: defaultBucketHandle,
			want:         true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Set the file sizes to be much smaller for the test.
			defer func(prevOneGB, prevSixteenGB int64) {
				oneGB = prevOneGB
				sixteenGB = prevSixteenGB
			}(oneGB, sixteenGB)
			oneGB /= 8
			sixteenGB = oneGB

			got := Execute(context.Background(), defaultConfig, test.bucketHandle, bytes.NewBufferString(""))
			if got != test.want {
				t.Errorf("Execute() = %v, want: %v", got, test.want)
			}
		})
	}
}

func TestCreateFiles(t *testing.T) {
	tests := []struct {
		name      string
		fileName1 string
		fileSize1 int64
		fileName2 string
		fileSize2 int64
		want      int
		wantError error
	}{
		{
			name:      "CreateFile1Fail",
			wantError: cmpopts.AnyError,
		},
		{
			name:      "TruncateFile1Fail",
			fileName1: t.TempDir() + "/object1.txt",
			fileSize1: -1,
			wantError: cmpopts.AnyError,
		},
		{
			name:      "CreateFile2Fail",
			fileName1: t.TempDir() + "/object1.txt",
			wantError: cmpopts.AnyError,
		},
		{
			name:      "TruncateFile2Fail",
			fileName1: t.TempDir() + "/object1.txt",
			fileName2: t.TempDir() + "/object2.txt",
			fileSize2: -1,
			wantError: cmpopts.AnyError,
		},
		{
			name:      "Success",
			fileName1: t.TempDir() + "/object1.txt",
			fileName2: t.TempDir() + "/object2.txt",
			want:      3,
			wantError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotError := createFiles(context.Background(), test.fileName1, test.fileName2, test.fileSize1, test.fileSize2)
			if len(got) != test.want {
				t.Errorf("createFiles(%s, %s) returned %d files, want: %d", test.fileName1, test.fileName2, len(got), test.want)
			}
			if !cmp.Equal(gotError, test.wantError, cmpopts.EquateErrors()) {
				t.Errorf("createFiles(%s, %s) = %v, wantError: %v", test.fileName1, test.fileName2, gotError, test.wantError)
			}
		})
	}
}

func TestRemoveFiles(t *testing.T) {
	tests := []struct {
		name   string
		opts   diagnoseOptions
		remove removeFunc
		want   bool
	}{
		{
			name: "FailNonTmpFolder",
			opts: diagnoseOptions{
				files: []*diagnoseFile{
					{fileName: "/"},
				},
			},
			remove: defaultRemoveFunc,
			want:   false,
		},
		{
			name: "FailRemoveLocal",
			opts: diagnoseOptions{
				files: []*diagnoseFile{
					{fileName: "/tmp/object.txt"},
				},
			},
			remove: func(name string) error {
				return errors.New("remove error")
			},
			want: false,
		},
		{
			name: "FailRemoveBucket",
			opts: diagnoseOptions{
				files: []*diagnoseFile{
					{fileName: "/tmp/object.txt"},
				},
				bucketHandle: &store.BucketHandle{},
			},
			remove: defaultRemoveFunc,
			want:   false,
		},
		{
			name: "Success",
			opts: diagnoseOptions{
				bucketHandle: defaultBucketHandle,
				files: []*diagnoseFile{
					{fileName: "/tmp/object.txt"},
				},
			},
			remove: defaultRemoveFunc,
			want:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := removeFiles(context.Background(), test.opts, test.remove)
			if !cmp.Equal(got, test.want) {
				t.Errorf("removeFiles(%#v) = %v, want: %v", test.opts, got, test.want)
			}
		})
	}
}

func TestDiagnoseBackup(t *testing.T) {
	tests := []struct {
		name  string
		opts  diagnoseOptions
		lines []string
		want  error
	}{
		{
			name: "NoFiles",
			opts: diagnoseOptions{
				files: []*diagnoseFile{},
			},
			want: cmpopts.AnyError,
		},
		{
			name: "BackupFailed",
			opts: defaultOptions,
			want: cmpopts.AnyError,
		},
		{
			name:  "BackupErrorFailed",
			opts:  defaultOptions,
			lines: []string{`#SAVED "12345" "/object.txt" 12345`},
			want:  cmpopts.AnyError,
		},
		{
			name: "BackupSuccess",
			opts: defaultOptions,
			lines: []string{
				`#SAVED "12345" "/object.txt" 12345`,
				`#ERROR "/object.txt"`,
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.opts.execute = defaultExecute(test.lines)
			got := diagnoseBackup(context.Background(), test.opts)
			if !cmp.Equal(got, test.want, cmpopts.EquateErrors()) {
				t.Errorf("diagnoseBackup(%#v) = %v, want: %v", test.opts, got, test.want)
			}
		})
	}
}

func TestDiagnoseInquire(t *testing.T) {
	tests := []struct {
		name  string
		opts  diagnoseOptions
		lines []string
		want  error
	}{
		{
			name: "NoFiles",
			opts: diagnoseOptions{
				files: []*diagnoseFile{},
			},
			want: cmpopts.AnyError,
		},
		{
			name: "InquireFailed",
			opts: defaultOptions,
			want: cmpopts.AnyError,
		},
		{
			name:  "InquireNullFailed",
			opts:  defaultOptions,
			lines: []string{`#BACKUP "12345" "/object.txt"`},
			want:  cmpopts.AnyError,
		},
		{
			name: "InquireTimestampsOutOfOrder",
			opts: defaultOptions,
			lines: []string{
				`#BACKUP "12345" "/object.txt"`,
				fmt.Sprintf(`#SOFTWAREID "backint 1.50" "Google %s %s"`, configuration.AgentName, configuration.AgentVersion) + "\n" + `#BACKUP "12345" "/object.txt" "2006-01-02T15:04:05.999Z07:00"` + "\n" + `#BACKUP "12345" "/object.txt" "2026-01-02T15:04:05.999Z07:00"`,
			},
			want: cmpopts.AnyError,
		},
		{
			name: "InquireNotFoundFailed",
			opts: defaultOptions,
			lines: []string{
				`#BACKUP "12345" "/object.txt"`,
				fmt.Sprintf(`#SOFTWAREID "backint 1.50" "Google %s %s"`, configuration.AgentName, configuration.AgentVersion) + "\n" + `#BACKUP "12345" "/object.txt" "2006-01-02T15:04:05.999Z07:00"` + "\n" + `#BACKUP "12345" "/object.txt" "2006-01-02T15:04:05.999Z07:00"`,
			},
			want: cmpopts.AnyError,
		},
		{
			name: "InquireErrorFailed",
			opts: defaultOptions,
			lines: []string{
				`#BACKUP "12345" "/object.txt"`,
				fmt.Sprintf(`#SOFTWAREID "backint 1.50" "Google %s %s"`, configuration.AgentName, configuration.AgentVersion) + "\n" + `#BACKUP "12345" "/object.txt" "2006-01-02T15:04:05.999Z07:00"` + "\n" + `#BACKUP "12345" "/object.txt" "2006-01-02T15:04:05.999Z07:00"`,
				`#NOTFOUND "/object.txt"`,
			},
			want: cmpopts.AnyError,
		},
		{
			name: "InquireSuccess",
			opts: defaultOptions,
			lines: []string{
				`#BACKUP "12345" "/object.txt"`,
				fmt.Sprintf(`#SOFTWAREID "backint 1.50" "Google %s %s"`, configuration.AgentName, configuration.AgentVersion) + "\n" + `#BACKUP "12345" "/object.txt" "2006-01-02T15:04:05.999Z07:00"` + "\n" + `#BACKUP "12345" "/object.txt" "2006-01-02T15:04:05.999Z07:00"`,
				`#NOTFOUND "/object.txt"`,
				`#ERROR "/object.txt"`,
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.opts.execute = defaultExecute(test.lines)
			got := diagnoseInquire(context.Background(), test.opts)
			if !cmp.Equal(got, test.want, cmpopts.EquateErrors()) {
				t.Errorf("diagnoseInquire(%#v) = %v, want: %v", test.opts, got, test.want)
			}
		})
	}
}

func TestDiagnoseRestore(t *testing.T) {
	tests := []struct {
		name  string
		opts  diagnoseOptions
		lines []string
		want  error
	}{
		{
			name: "NoFiles",
			opts: diagnoseOptions{
				files: []*diagnoseFile{},
			},
			want: cmpopts.AnyError,
		},
		{
			name: "RestoreFailed",
			opts: defaultOptions,
			want: cmpopts.AnyError,
		},
		{
			name:  "RestoreNullFailed",
			opts:  defaultOptions,
			lines: []string{`#RESTORED "12345" "/object.txt"`},
			want:  cmpopts.AnyError,
		},
		{
			name: "RestoreNotFoundFailed",
			opts: defaultOptions,
			lines: []string{
				`#RESTORED "12345" "/object.txt"`,
				`#RESTORED "12345" "/object.txt"`,
			},
			want: cmpopts.AnyError,
		},
		{
			name: "RestoreErrorFailed",
			opts: defaultOptions,
			lines: []string{
				`#RESTORED "12345" "/object.txt"`,
				`#RESTORED "12345" "/object.txt"`,
				`#NOTFOUND "/object.txt"`,
			},
			want: cmpopts.AnyError,
		},
		{
			name: "RestoreSuccess",
			opts: defaultOptions,
			lines: []string{
				`#RESTORED "12345" "/object.txt"`,
				`#RESTORED "12345" "/object.txt"`,
				`#NOTFOUND "/object.txt"`,
				`#ERROR "/object.txt"`,
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.opts.execute = defaultExecute(test.lines)
			got := diagnoseRestore(context.Background(), test.opts)
			if !cmp.Equal(got, test.want, cmpopts.EquateErrors()) {
				t.Errorf("diagnoseRestore(%#v) = %v, want: %v", test.opts, got, test.want)
			}
		})
	}
}

func TestDiagnoseDelete(t *testing.T) {
	tests := []struct {
		name  string
		opts  diagnoseOptions
		lines []string
		want  error
	}{
		{
			name: "NoFiles",
			opts: diagnoseOptions{
				files: []*diagnoseFile{},
			},
			want: cmpopts.AnyError,
		},
		{
			name: "DeleteFailed",
			opts: defaultOptions,
			want: cmpopts.AnyError,
		},
		{
			name:  "DeleteNotFoundFailed",
			opts:  defaultOptions,
			lines: []string{`#DELETED "12345" "/object.txt"`},
			want:  cmpopts.AnyError,
		},
		{
			name: "DeleteErrorFailed",
			opts: defaultOptions,
			lines: []string{
				`#DELETED "12345" "/object.txt"`,
				`#NOTFOUND "12345" "/object.txt"`,
			},
			want: cmpopts.AnyError,
		},
		{
			name: "DeleteSuccess",
			opts: defaultOptions,
			lines: []string{
				`#DELETED "12345" "/object.txt"`,
				`#NOTFOUND "12345" "/object.txt"`,
				`#ERROR "12345" "/object.txt"`,
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.opts.execute = defaultExecute(test.lines)
			got := diagnoseDelete(context.Background(), test.opts)
			if !cmp.Equal(got, test.want, cmpopts.EquateErrors()) {
				t.Errorf("diagnoseDelete(%#v) = %v, want: %v", test.opts, got, test.want)
			}
		})
	}
}

func TestPerformDiagnostic(t *testing.T) {
	tests := []struct {
		name           string
		lines          []string
		wantExecuteOut string
		want           [][]string
		wantError      error
	}{
		{
			name:      "ExecuteFailed",
			wantError: cmpopts.AnyError,
		},
		{
			name:           "WantLinesFailed",
			lines:          []string{`#RESTORED "12345" "/object.txt"`},
			wantExecuteOut: "#RESTORED\n#RESTORED",
			wantError:      cmpopts.AnyError,
		},
		{
			name:           "WantPrefixFailed",
			lines:          []string{`#RESTORED "12345" "/object.txt"`},
			wantExecuteOut: "#BACKUP",
			wantError:      cmpopts.AnyError,
		},
		{
			name:           "WantSplitFailed",
			lines:          []string{`#RESTORED "12345" "/object.txt"`},
			wantExecuteOut: "#RESTORED <external_backup_id>",
			wantError:      cmpopts.AnyError,
		},
		{
			name:           "SuccessOneLine",
			lines:          []string{`#RESTORED "12345" "/object.txt"`},
			wantExecuteOut: "#RESTORED <external_backup_id> <filename>",
			want:           [][]string{{`#RESTORED`, `"12345"`, `"/object.txt"`}},
			wantError:      nil,
		},
		{
			name:           "SuccessMultipleLines",
			lines:          []string{`#RESTORED "12345" "/object.txt"` + "\n" + `#RESTORED "12345" "/object2.txt"`},
			wantExecuteOut: "#RESTORED <external_backup_id> <filename>\n#RESTORED <external_backup_id> <filename>",
			want:           [][]string{{`#RESTORED`, `"12345"`, `"/object.txt"`}, {`#RESTORED`, `"12345"`, `"/object2.txt"`}},
			wantError:      nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := defaultOptions
			opts.execute = defaultExecute(test.lines)
			opts.want = test.wantExecuteOut
			got, gotError := performDiagnostic(context.Background(), opts)
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("performDiagnostic(%#v) had unexpected diff (-want +got):\n%s", opts, diff)
			}
			if !cmp.Equal(gotError, test.wantError, cmpopts.EquateErrors()) {
				t.Errorf("performDiagnostic(%#v) = %v, wantError: %v", opts, gotError, test.wantError)
			}
		})
	}
}

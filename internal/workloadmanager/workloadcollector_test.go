/*
Copyright 2022 Google LLC

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

package workloadmanager

import (
	"context"
	"embed"
	"errors"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	metricpb "google.golang.org/genproto/googleapis/api/metric"
	monitoredresourcepb "google.golang.org/genproto/googleapis/api/monitoredres"
	monitoringresourcespb "google.golang.org/genproto/googleapis/monitoring/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/GoogleCloudPlatform/sapagent/internal/cloudmonitoring"
	"github.com/GoogleCloudPlatform/sapagent/internal/cloudmonitoring/fake"
	"github.com/GoogleCloudPlatform/sapagent/internal/commandlineexecutor"

	workloadmanager "google.golang.org/api/workloadmanager/v1"
	"github.com/GoogleCloudPlatform/sapagent/internal/heartbeat"
	"github.com/GoogleCloudPlatform/sapagent/internal/instanceinfo"
	cmpb "github.com/GoogleCloudPlatform/sapagent/protos/configurablemetrics"
	cfgpb "github.com/GoogleCloudPlatform/sapagent/protos/configuration"
	iipb "github.com/GoogleCloudPlatform/sapagent/protos/instanceinfo"
	wlmpb "github.com/GoogleCloudPlatform/sapagent/protos/wlmvalidation"
)

var (
	defaultConfiguration = &cfgpb.Configuration{
		CloudProperties: &iipb.CloudProperties{
			InstanceName: "test-instance-name",
			InstanceId:   "test-instance-id",
			Zone:         "test-region-zone",
			ProjectId:    "test-project-id",
		},
		AgentProperties: &cfgpb.AgentProperties{Name: "sapagent", Version: "1.0"},
		CollectionConfiguration: &cfgpb.CollectionConfiguration{
			CollectWorkloadValidationMetrics: true,
		},
	}
	bmConfiguration = &cfgpb.Configuration{
		BareMetal: true,
		CloudProperties: &iipb.CloudProperties{
			ProjectId:  "bm-project-id",
			InstanceId: "bm-instance-id",
			Region:     "us-central1",
		},
		AgentProperties: &cfgpb.AgentProperties{Name: "sapagent", Version: "1.0"},
		CollectionConfiguration: &cfgpb.CollectionConfiguration{
			CollectWorkloadValidationMetrics: true,
		},
	}
	defaultConfigurationDBMetrics = &cfgpb.Configuration{
		CloudProperties: &iipb.CloudProperties{
			InstanceName: "test-instance-name",
			InstanceId:   "test-instance-id",
			Zone:         "test-region-zone",
			ProjectId:    "test-project-id",
		},
		AgentProperties: &cfgpb.AgentProperties{Name: "sapagent", Version: "1.0"},
		CollectionConfiguration: &cfgpb.CollectionConfiguration{
			CollectWorkloadValidationMetrics: true,
			WorkloadValidationDbMetricsConfig: &cfgpb.HANAMetricsConfig{
				HanaDbUser:     "SYSTEM",
				HanaDbPassword: "dummy-pass",
				Hostname:       "test-hostname",
				Port:           "30015",
			},
		},
	}
	//go:embed test_data/metricoverride.yaml
	sampleOverride string
	//go:embed test_data/metricoverride.yaml test_data/os-release.txt test_data/os-release-bad.txt test_data/os-release-empty.txt
	testFS                  embed.FS
	defaultBackOffIntervals = cloudmonitoring.NewBackOffIntervals(time.Millisecond, time.Millisecond)
	DefaultTestReader       = ConfigFileReader(func(data string) (io.ReadCloser, error) {
		return io.NopCloser(strings.NewReader(data)), nil
	})
)

type WriteInsightArgs struct {
	Project  string
	Location string
	Req      *workloadmanager.WriteInsightRequest
}

type testWLMInterface struct {
	T                     *testing.T
	WriteInsightArgs      []WriteInsightArgs
	WriteInsightErrs      []error
	WriteInsightCallCount int
}

func validationDetailSort(a, b *workloadmanager.SapValidationValidationDetail) bool {
	return a.SapValidationType < b.SapValidationType
}

func (t *testWLMInterface) WriteInsight(project string, location string, req *workloadmanager.WriteInsightRequest) error {
	defer func() { t.WriteInsightCallCount++ }()

	if diff := cmp.Diff(t.WriteInsightArgs[t.WriteInsightCallCount], WriteInsightArgs{project, location, req}, cmpopts.SortSlices(validationDetailSort)); diff != "" {
		t.T.Errorf("WriteInsight() arguments diff (-want +got):\n%s", diff)
	}

	return t.WriteInsightErrs[t.WriteInsightCallCount]
}

func TestSetOSReleaseInfo(t *testing.T) {
	defaultFileReader := ConfigFileReader(func(path string) (io.ReadCloser, error) {
		file, err := testFS.Open(path)
		var f io.ReadCloser = file
		return f, err
	})

	tests := []struct {
		name        string
		filePath    string
		reader      ConfigFileReader
		wantID      string
		wantVersion string
	}{
		{
			name:        "Success",
			filePath:    "test_data/os-release.txt",
			reader:      defaultFileReader,
			wantID:      "debian",
			wantVersion: "11",
		},
		{
			name:        "ConfigFileReaderNil",
			filePath:    "test_data/os-release.txt",
			reader:      nil,
			wantID:      "",
			wantVersion: "",
		},
		{
			name:        "OSReleaseFilePathEmpty",
			filePath:    "",
			reader:      defaultFileReader,
			wantID:      "",
			wantVersion: "",
		},
		{
			name:     "FileReadError",
			filePath: "test_data/os-release.txt",
			reader: ConfigFileReader(func(path string) (io.ReadCloser, error) {
				return nil, errors.New("File Read Error")
			}),
			wantID:      "",
			wantVersion: "",
		},
		{
			name:        "FileParseError",
			filePath:    "test_data/os-release-bad.txt",
			reader:      defaultFileReader,
			wantID:      "",
			wantVersion: "",
		},
		{
			name:        "FieldsEmpty",
			filePath:    "test_data/os-release-empty.txt",
			reader:      defaultFileReader,
			wantID:      "",
			wantVersion: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			params := Parameters{
				OSReleaseFilePath: test.filePath,
				ConfigFileReader:  test.reader,
			}
			params.SetOSReleaseInfo()
			if params.osVendorID != test.wantID {
				t.Errorf("SetOSReleaseInfo() unexpected osVendorID, got %q want %q", params.osVendorID, test.wantID)
			}
			if params.osVersion != test.wantVersion {
				t.Errorf("SetOSReleaseInfo() unexpected osVersion, got %q want %q", params.osVersion, test.wantVersion)
			}
		})
	}
}

func TestCollectMetricsFromConfig(t *testing.T) {
	tests := []struct {
		name     string
		params   Parameters
		override string
		want     WorkloadMetrics
	}{
		{
			name: "HasMetricOverride",
			params: Parameters{
				Config: defaultConfiguration,
				OSStatReader: func(data string) (os.FileInfo, error) {
					f, err := testFS.Open(data)
					if err != nil {
						return nil, err
					}
					return f.Stat()
				},
				ConfigFileReader: ConfigFileReader(func(path string) (io.ReadCloser, error) {
					file, err := testFS.Open(path)
					var f io.ReadCloser = file
					return f, err
				}),
			},
			override: "test_data/metricoverride.yaml",
			want: WorkloadMetrics{Metrics: append(createTimeSeries(
				"workload.googleapis.com/sap/validation/system",
				map[string]string{"blank_metric": "", "metric_with_colons": "val1:val2:val3", "os": "rhel-8.4"},
				1.0,
				defaultConfiguration,
			), createTimeSeries(
				"workload.googleapis.com/sap/validation/hana",
				map[string]string{"hana_metric": "/hana/log"},
				0.0,
				defaultConfiguration,
			)...)},
		},
		{
			name: "HasWorkloadConfig",
			params: Parameters{
				Config: defaultConfiguration,
				WorkloadConfig: &wlmpb.WorkloadValidation{
					ValidationSystem: &wlmpb.ValidationSystem{
						SystemMetrics: []*wlmpb.SystemMetric{
							{
								MetricInfo: &cmpb.MetricInfo{
									Type:  "workload.googleapis.com/sap/validation/system",
									Label: "agent",
								},
								Value: wlmpb.SystemVariable_AGENT_NAME,
							},
						},
					},
				},
				Exists: func(string) bool { return false },
				Execute: func(context.Context, commandlineexecutor.Params) commandlineexecutor.Result {
					return commandlineexecutor.Result{}
				},
				OSStatReader:       func(data string) (os.FileInfo, error) { return nil, nil },
				InstanceInfoReader: *instanceinfo.New(&fakeDiskMapper{}, defaultGCEService),
			},
			want: WorkloadMetrics{Metrics: []*monitoringresourcespb.TimeSeries{
				createTimeSeries(
					"workload.googleapis.com/sap/validation/system",
					map[string]string{"agent": "sapagent"},
					1.0,
					defaultConfiguration,
				)[0],
				createTimeSeries(
					"workload.googleapis.com/sap/validation/corosync",
					map[string]string{"agent": "sapagent"},
					0.0,
					defaultConfiguration,
				)[0],
				createTimeSeries(
					"workload.googleapis.com/sap/validation/hana",
					map[string]string{"agent": "sapagent"},
					0.0,
					defaultConfiguration,
				)[0],
				createTimeSeries(
					"workload.googleapis.com/sap/validation/netweaver",
					map[string]string{"agent": "sapagent"},
					0.0,
					defaultConfiguration,
				)[0],
				createTimeSeries(
					"workload.googleapis.com/sap/validation/pacemaker",
					map[string]string{"agent": "sapagent"},
					0.0,
					defaultConfiguration,
				)[0],
				createTimeSeries(
					"workload.googleapis.com/sap/validation/custom",
					map[string]string{"agent": "sapagent"},
					1.0,
					defaultConfiguration,
				)[0],
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := collectMetricsFromConfig(context.Background(), test.params, test.override)
			if diff := cmp.Diff(test.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("collectMetricsFromConfig() returned unexpected diff (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOverrideMetrics(t *testing.T) {
	tests := []struct {
		file   string
		reader ConfigFileReader
		want   WorkloadMetrics
	}{
		{
			file:   "",
			reader: DefaultTestReader,
			want:   WorkloadMetrics{},
		},
		{
			file: sampleOverride,
			reader: ConfigFileReader(func(data string) (io.ReadCloser, error) {
				return nil, errors.New("failed to read file")
			}),
			want: WorkloadMetrics{},
		},
		{
			file:   sampleOverride,
			reader: DefaultTestReader,
			want: WorkloadMetrics{Metrics: append(createTimeSeries(
				"workload.googleapis.com/sap/validation/system",
				map[string]string{"blank_metric": "", "metric_with_colons": "val1:val2:val3", "os": "rhel-8.4"},
				1.0,
				&cfgpb.Configuration{},
			), createTimeSeries(
				"workload.googleapis.com/sap/validation/hana",
				map[string]string{"hana_metric": "/hana/log"},
				.0,
				&cfgpb.Configuration{},
			)...)},
		},
	}

	for _, test := range tests {
		got := collectOverrideMetrics(&cfgpb.Configuration{}, test.reader, test.file)

		if diff := cmp.Diff(test.want, got, protocmp.Transform()); diff != "" {
			t.Errorf("collectOverrideMetrics returned unexpected metrics diff (-want +got):\n%s", diff)
		}
	}
}

func TestSendMetrics(t *testing.T) {
	tests := []struct {
		name            string
		wlmInterface    *testWLMInterface
		params          sendMetricsParams
		wantMetricCount int
	}{
		{
			name:         "succeedsWithZeroMetrics",
			wlmInterface: &testWLMInterface{},
			params: sendMetricsParams{
				wm:                WorkloadMetrics{},
				cp:                defaultConfiguration.GetCloudProperties(),
				timeSeriesCreator: &fake.TimeSeriesCreator{},
				backOffIntervals:  defaultBackOffIntervals,
			},
			wantMetricCount: 0,
		},
		{
			name: "succeedsWithMetrics",
			wlmInterface: &testWLMInterface{
				WriteInsightArgs: []WriteInsightArgs{
					{
						Project:  "test-project-id",
						Location: "test-region",
						Req: &workloadmanager.WriteInsightRequest{
							Insight: &workloadmanager.Insight{
								InstanceId: "test-instance-id",
								SapValidation: &workloadmanager.SapValidation{
									ValidationDetails: []*workloadmanager.SapValidationValidationDetail{
										&workloadmanager.SapValidationValidationDetail{SapValidationType: "SAP_VALIDATION_TYPE_UNSPECIFIED"},
									},
								},
							},
						},
					}},
				WriteInsightErrs: []error{nil},
			},
			params: sendMetricsParams{
				wm: WorkloadMetrics{Metrics: []*monitoringresourcespb.TimeSeries{{
					Metric:   &metricpb.Metric{},
					Resource: &monitoredresourcepb.MonitoredResource{},
					Points:   []*monitoringresourcespb.Point{},
				}}},
				cp:                defaultConfiguration.GetCloudProperties(),
				timeSeriesCreator: &fake.TimeSeriesCreator{},
				backOffIntervals:  defaultBackOffIntervals,
			},
			wantMetricCount: 1,
		},
		{
			name: "succeedsBareMetal",
			wlmInterface: &testWLMInterface{
				WriteInsightArgs: []WriteInsightArgs{
					{
						Project:  "bm-project-id",
						Location: "us-central1",
						Req: &workloadmanager.WriteInsightRequest{
							Insight: &workloadmanager.Insight{
								InstanceId: "bm-instance-id",
								SapValidation: &workloadmanager.SapValidation{
									ValidationDetails: []*workloadmanager.SapValidationValidationDetail{
										&workloadmanager.SapValidationValidationDetail{SapValidationType: "SAP_VALIDATION_TYPE_UNSPECIFIED"},
									},
								},
							},
						},
					},
				},
				WriteInsightErrs: []error{nil},
			},
			params: sendMetricsParams{
				wm: WorkloadMetrics{Metrics: []*monitoringresourcespb.TimeSeries{{
					Metric:   &metricpb.Metric{},
					Resource: &monitoredresourcepb.MonitoredResource{},
					Points:   []*monitoringresourcespb.Point{},
				}}},
				cp:                bmConfiguration.GetCloudProperties(),
				bareMetal:         true,
				timeSeriesCreator: &fake.TimeSeriesCreator{},
				backOffIntervals:  defaultBackOffIntervals,
			},
			wantMetricCount: 1,
		},
		{
			name: "failsSendToCloudMonitoring",
			wlmInterface: &testWLMInterface{
				WriteInsightArgs: []WriteInsightArgs{
					{
						Project:  "test-project-id",
						Location: "test-region",
						Req: &workloadmanager.WriteInsightRequest{
							Insight: &workloadmanager.Insight{
								InstanceId: "test-instance-id",
								SapValidation: &workloadmanager.SapValidation{
									ValidationDetails: []*workloadmanager.SapValidationValidationDetail{
										&workloadmanager.SapValidationValidationDetail{SapValidationType: "SAP_VALIDATION_TYPE_UNSPECIFIED"},
									},
								},
							},
						},
					},
				},
				WriteInsightErrs: []error{nil},
			},
			params: sendMetricsParams{
				wm: WorkloadMetrics{Metrics: []*monitoringresourcespb.TimeSeries{{
					Metric:   &metricpb.Metric{},
					Resource: &monitoredresourcepb.MonitoredResource{},
					Points:   []*monitoringresourcespb.Point{},
				}}},
				cp:                defaultConfiguration.GetCloudProperties(),
				timeSeriesCreator: &fake.TimeSeriesCreator{Err: cmpopts.AnyError},
				backOffIntervals:  defaultBackOffIntervals,
			},
			wantMetricCount: 0,
		},
		{
			wlmInterface: &testWLMInterface{
				WriteInsightArgs: []WriteInsightArgs{{
					Project:  "test-project-id",
					Location: "test-region",
					Req: &workloadmanager.WriteInsightRequest{
						Insight: &workloadmanager.Insight{
							InstanceId: "test-instance-id",
							SapValidation: &workloadmanager.SapValidation{
								ValidationDetails: []*workloadmanager.SapValidationValidationDetail{
									&workloadmanager.SapValidationValidationDetail{SapValidationType: "SAP_VALIDATION_TYPE_UNSPECIFIED"},
								},
							},
						},
					},
				}},
				WriteInsightErrs: []error{cmpopts.AnyError},
			},
			name: "failsSendToDataWarehouse",
			params: sendMetricsParams{
				wm: WorkloadMetrics{Metrics: []*monitoringresourcespb.TimeSeries{{
					Metric:   &metricpb.Metric{},
					Resource: &monitoredresourcepb.MonitoredResource{},
					Points:   []*monitoringresourcespb.Point{},
				}}},
				cp:                defaultConfiguration.GetCloudProperties(),
				timeSeriesCreator: &fake.TimeSeriesCreator{},
				backOffIntervals:  defaultBackOffIntervals,
			},
			wantMetricCount: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.wlmInterface.T = t
			test.params.wlmService = test.wlmInterface
			got := sendMetrics(context.Background(), test.params)
			if got != test.wantMetricCount {
				t.Errorf("sendMetrics returned unexpected metric count for %s, got: %d, want: %d",
					test.name, got, test.wantMetricCount)
			}
		})
	}
}

func TestAppendLabels(t *testing.T) {
	s := make(map[string]string)
	s["skey"] = "svalue"
	got := make(map[string]string)
	got["dkey"] = "dvalue"
	want := make(map[string]string)
	want["dkey"] = "dvalue"
	want["skey"] = "svalue"

	appendLabels(got, s)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("appendLabels maps are not equal, got: %v, want: %v", got, want)
	}
}

func TestStartMetricsCollection(t *testing.T) {
	tests := []struct {
		name         string
		params       Parameters
		os           string
		wlmInterface *testWLMInterface
		want         bool
	}{
		{
			name: "succeedsForLocal",
			params: Parameters{
				Config: defaultConfiguration,
				Execute: func(context.Context, commandlineexecutor.Params) commandlineexecutor.Result {
					return commandlineexecutor.Result{
						StdOut: "",
						StdErr: "",
					}
				},
				Exists:            func(string) bool { return true },
				ConfigFileReader:  DefaultTestReader,
				OSStatReader:      func(data string) (os.FileInfo, error) { return nil, nil },
				TimeSeriesCreator: &fake.TimeSeriesCreator{},
				OSType:            "linux",
				Remote:            false,
				BackOffs:          defaultBackOffIntervals,
			},
			wlmInterface: &testWLMInterface{
				WriteInsightArgs: []WriteInsightArgs{{}},
				WriteInsightErrs: []error{nil},
			},
			want: true,
		},
		{
			name: "succeedsForRemote",
			params: Parameters{
				Config: defaultConfiguration,
				Execute: func(context.Context, commandlineexecutor.Params) commandlineexecutor.Result {
					return commandlineexecutor.Result{
						StdOut: "",
						StdErr: "",
					}
				},
				Exists:            func(string) bool { return true },
				ConfigFileReader:  DefaultTestReader,
				OSStatReader:      func(data string) (os.FileInfo, error) { return nil, nil },
				TimeSeriesCreator: &fake.TimeSeriesCreator{},
				OSType:            "linux",
				Remote:            true,
				BackOffs:          defaultBackOffIntervals,
			},
			wlmInterface: &testWLMInterface{
				WriteInsightArgs: []WriteInsightArgs{{}},
				WriteInsightErrs: []error{nil},
			},
			want: true,
		},
		{
			name: "succeedsForLocalWithDBMetrics",
			params: Parameters{
				Config: defaultConfigurationDBMetrics,
				Execute: func(context.Context, commandlineexecutor.Params) commandlineexecutor.Result {
					return commandlineexecutor.Result{
						StdOut: "",
						StdErr: "",
					}
				},
				Exists:            func(string) bool { return true },
				ConfigFileReader:  DefaultTestReader,
				OSStatReader:      func(data string) (os.FileInfo, error) { return nil, nil },
				TimeSeriesCreator: &fake.TimeSeriesCreator{},
				OSType:            "linux",
				Remote:            false,
				BackOffs:          defaultBackOffIntervals,
			},
			wlmInterface: &testWLMInterface{
				WriteInsightArgs: []WriteInsightArgs{{}},
				WriteInsightErrs: []error{nil},
			},
			want: true,
		},
		{
			name: "failsDueToParams",
			params: Parameters{
				Config: &cfgpb.Configuration{
					CollectionConfiguration: &cfgpb.CollectionConfiguration{
						CollectWorkloadValidationMetrics: false,
					}},
				OSType:   "linux",
				BackOffs: defaultBackOffIntervals,
			},
			wlmInterface: &testWLMInterface{
				WriteInsightArgs: []WriteInsightArgs{{}},
				WriteInsightErrs: []error{nil},
			},
			want: false,
		},
		{
			name: "failsDueToOS",
			params: Parameters{
				Config:   defaultConfiguration,
				OSType:   "windows",
				BackOffs: defaultBackOffIntervals,
			},
			wlmInterface: &testWLMInterface{
				WriteInsightArgs: []WriteInsightArgs{{}},
				WriteInsightErrs: []error{nil},
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.wlmInterface.T = t
			test.params.WLMService = test.wlmInterface
			ctx, cancel := context.WithCancel(context.Background())
			t.Cleanup(cancel)
			got := StartMetricsCollection(ctx, test.params)
			if got != test.want {
				t.Errorf("StartMetricsCollection(%#v) returned unexpected result, got: %t, want: %t", test.params, got, test.want)
			}
		})
	}
}

func TestCollectAndSend_shouldBeatAccordingToHeartbeatSpec(t *testing.T) {
	testData := []struct {
		name         string
		beatInterval time.Duration
		timeout      time.Duration
		want         int
	}{
		{
			name:         "CancelBeforeInitialCollection",
			beatInterval: time.Second,
			timeout:      time.Second * 0,
			want:         0,
		},
		{
			name:         "CancelBeforeBeat",
			beatInterval: time.Second * 1,
			timeout:      time.Millisecond * 50,
			want:         1,
		},
		{
			name:         "Cancel1Beat",
			beatInterval: time.Millisecond * 75,
			timeout:      time.Millisecond * 140,
			want:         2,
		},
		{
			name:         "Cancel2Beats",
			beatInterval: time.Millisecond * 45,
			timeout:      time.Millisecond * 125,
			want:         3,
		},
	}
	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, test.timeout)
			defer cancel()
			got := 0
			lock := sync.Mutex{}
			params := Parameters{
				Config: defaultConfiguration,
				Execute: func(context.Context, commandlineexecutor.Params) commandlineexecutor.Result {
					return commandlineexecutor.Result{
						StdOut: "",
						StdErr: "",
					}
				},
				Exists:            func(string) bool { return true },
				ConfigFileReader:  DefaultTestReader,
				OSStatReader:      func(data string) (os.FileInfo, error) { return nil, nil },
				TimeSeriesCreator: &fake.TimeSeriesCreator{},
				OSType:            "linux",
				Remote:            false,
				BackOffs:          defaultBackOffIntervals,
				WLMService: &testWLMInterface{
					T:                t,
					WriteInsightErrs: []error{nil},
					WriteInsightArgs: []WriteInsightArgs{{
						Project:  "test-project-id",
						Location: "test-region",
						Req: &workloadmanager.WriteInsightRequest{Insight: &workloadmanager.Insight{
							InstanceId: "test-instance-id",
							SapValidation: &workloadmanager.SapValidation{
								ValidationDetails: []*workloadmanager.SapValidationValidationDetail{
									{
										SapValidationType: "SYSTEM",
										Details:           map[string]string{},
									},
									{
										SapValidationType: "NETWEAVER",
										Details:           map[string]string{},
									},
									{
										SapValidationType: "HANA",
										Details:           map[string]string{},
									},
									{
										SapValidationType: "PACEMAKER",
										Details:           map[string]string{},
									},
									{
										SapValidationType: "COROSYNC",
										Details:           map[string]string{},
									},
									{
										SapValidationType: "CUSTOM",
										Details:           map[string]string{},
									},
								},
							}},
						},
					}},
				},
				HeartbeatSpec: &heartbeat.Spec{
					BeatFunc: func() {
						lock.Lock()
						defer lock.Unlock()
						got++
					},
					Interval: test.beatInterval,
				},
			}

			StartMetricsCollection(ctx, params)

			<-ctx.Done()
			lock.Lock()
			defer lock.Unlock()
			if got != test.want {
				t.Errorf("collectAndSend() heartbeat mismatch got %d, want %d", got, test.want)
			}
		})
	}
}

func TestReadHANAInsightsRules(t *testing.T) {
	params := Parameters{Config: defaultConfiguration}
	params.ReadHANAInsightsRules()
	if len(params.HANAInsightRules) != 20 {
		t.Errorf("ReadHANAInsightsRules() got: %d rules, want: %d.", len(params.HANAInsightRules), 20)
	}
}

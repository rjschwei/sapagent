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

package heartbeat

import (
	"context"
	"fmt"
	"math"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	cfgpb "github.com/GoogleCloudPlatform/sapagent/protos/configuration"
)

func createMonitor(t *testing.T, params Parameters) *Monitor {
	t.Helper()
	monitor, err := NewMonitor(params)
	if err != nil {
		t.Fatal(err)
	}
	return monitor
}

func register(t *testing.T, monitor *Monitor, name string) *Spec {
	t.Helper()
	spec, err := monitor.Register(name)
	if err != nil {
		t.Fatal(err)
	}
	return spec
}

func defaultMonitor(t *testing.T) *Monitor {
	params := Parameters{
		Config: &cfgpb.Configuration{
			CollectionConfiguration: &cfgpb.CollectionConfiguration{
				HeartbeatFrequency:       1,
				MissedHeartbeatThreshold: 5,
			},
		},
	}
	return createMonitor(t, params)
}

func TestNewMonitor_shouldIdentifyBadParameters(t *testing.T) {
	testData := []struct {
		name             string
		collectionConfig *cfgpb.CollectionConfiguration
		want             error
	}{
		{
			name: "valid parameters",
			collectionConfig: &cfgpb.CollectionConfiguration{
				MissedHeartbeatThreshold: 5,
				HeartbeatFrequency:       1,
			},
			want: nil,
		},
		{
			name: "negative threshold",
			collectionConfig: &cfgpb.CollectionConfiguration{
				MissedHeartbeatThreshold: -5,
				HeartbeatFrequency:       1,
			},
			want: cmpopts.AnyError,
		},
		{
			name: "negative interval",
			collectionConfig: &cfgpb.CollectionConfiguration{
				MissedHeartbeatThreshold: 10,
				HeartbeatFrequency:       -1,
			},
			want: cmpopts.AnyError,
		},
		{
			name: "negative threshold",
			collectionConfig: &cfgpb.CollectionConfiguration{
				MissedHeartbeatThreshold: -5,
				HeartbeatFrequency:       1,
			},
			want: cmpopts.AnyError,
		},
		{
			name:             "null configuration collection",
			collectionConfig: nil,
			want:             cmpopts.AnyError,
		},
	}
	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			params := Parameters{
				Config: &cfgpb.Configuration{
					CollectionConfiguration: d.collectionConfig,
				},
			}
			_, got := NewMonitor(params)
			if !cmp.Equal(got, d.want, cmpopts.EquateErrors()) {
				t.Errorf("validateParameters(%v) = %v, want %v", params, got, d.want)
			}
		})
	}
}

func TestRegister_shouldNotAllowDuplicateNames(t *testing.T) {
	testData := []struct {
		name       string
		names      []string
		wantErrors int
	}{
		{
			name:       "1a",
			names:      []string{"a"},
			wantErrors: 0,
		},
		{
			name:       "1a 1b",
			names:      []string{"a", "b"},
			wantErrors: 0,
		},
		{
			name:       "2a",
			names:      []string{"a", "a"},
			wantErrors: 1,
		},
		{
			name:       "2a 1b",
			names:      []string{"a", "a", "b"},
			wantErrors: 1,
		},
		{
			name:       "2a 2b",
			names:      []string{"a", "a", "b", "b"},
			wantErrors: 2,
		},
		{
			name: "20a 20b",
			names: []string{
				"a", "a", "a", "a", "a", "a", "a", "a", "a", "a",
				"a", "a", "a", "a", "a", "a", "a", "a", "a", "a",
				"b", "b", "b", "b", "b", "b", "b", "b", "b", "b",
				"b", "b", "b", "b", "b", "b", "b", "b", "b", "b",
			},
			wantErrors: 38,
		},
	}
	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			monitor := defaultMonitor(t)
			numErrors := 0
			lock := sync.Mutex{}
			wg := sync.WaitGroup{}
			wg.Add(len(d.names))
			for _, n := range d.names {
				clone := n
				go func() {
					_, err := monitor.Register(clone)
					if err != nil {
						lock.Lock()
						numErrors++
						lock.Unlock()
					}
					wg.Done()
				}()
			}
			wg.Wait()
			if numErrors != d.wantErrors {
				t.Errorf("Register() error count = %v, want %v", numErrors, d.wantErrors)
			}
		})
	}
}

func TestRun_shouldRespectCancellation(t *testing.T) {
	testData := []struct {
		name              string
		numRegistered     int
		timeout           time.Duration
		frequencyOverride time.Duration
		params            Parameters
		want              int64
	}{
		{
			name:          "cancel before any loops",
			numRegistered: 1,
			timeout:       time.Millisecond * 10,
			params: Parameters{
				Config: &cfgpb.Configuration{
					CollectionConfiguration: &cfgpb.CollectionConfiguration{
						MissedHeartbeatThreshold: 1000,
						HeartbeatFrequency:       1,
					}},
			},
			frequencyOverride: time.Millisecond * 100,
			want:              0,
		},
		{
			name:          "cancel 10ms after first loop",
			numRegistered: 1,
			timeout:       time.Millisecond * 110,
			params: Parameters{
				Config: &cfgpb.Configuration{
					CollectionConfiguration: &cfgpb.CollectionConfiguration{
						MissedHeartbeatThreshold: 1000,
						HeartbeatFrequency:       1,
					}},
			},
			frequencyOverride: time.Millisecond * 100,
			want:              1,
		},
	}

	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			monitor := createMonitor(t, d.params)
			monitor.frequency = d.frequencyOverride
			for i := 0; i < d.numRegistered; i++ {
				name := fmt.Sprintf("%d", i)
				register(t, monitor, name)
			}
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, d.timeout)
			defer cancel()
			monitor.Run(ctx)
			<-ctx.Done()
			got := monitor.registrations["0"].missedHeartbeats
			if got != d.want {
				t.Errorf("missed heart beats = %v, want %v", got, d.want)
			}
		})
	}
}

func TestNullMonitor_shouldNotErrorAndReturnNoops(t *testing.T) {
	testData := []struct {
		name  string
		names []string
	}{
		{
			name:  "1 registration",
			names: []string{"a"},
		},
		{
			name:  "2 registrations",
			names: []string{"a", "b"},
		},
		{
			name:  "3 registrations",
			names: []string{"a", "b", "c"},
		},
	}
	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			monitor := NullMonitor{}
			for _, name := range d.names {
				spec, err := monitor.Register(name)
				if err != nil {
					t.Errorf("Register() = (spec, %v), want (spec, nil)", err)
				}
				ticker := spec.CreateTicker()
				if ticker == nil {
					t.Error("CreateTicker() = nil, want non-nil")
				}
				spec.Beat()
			}
			got := monitor.GetStatuses()
			if got == nil || len(got) != 0 {
				t.Errorf("GetStatuses() = %v, want empty", got)
			}
		})
	}
}

func TestGetStatuses_shouldReturnTrueOnlyWhenMissedHeartbeatsIsLessThanThreshold(t *testing.T) {
	testData := []struct {
		name             string
		registrantNames  []string
		registrantMissed []int64
		threshold        int64
		want             map[string]bool
	}{
		{
			name:             "all under threshold",
			registrantNames:  []string{"a", "b"},
			registrantMissed: []int64{0, 0},
			threshold:        1,
			want:             map[string]bool{"a": true, "b": true},
		},
		{
			name:             "all at or above threshold",
			registrantNames:  []string{"a", "b"},
			registrantMissed: []int64{1, 2},
			threshold:        1,
			want:             map[string]bool{"a": false, "b": false},
		},
	}
	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			monitor := defaultMonitor(t)
			monitor.threshold = d.threshold
			numRegistrants := len(d.registrantNames)
			for regNum := 0; regNum < numRegistrants; regNum++ {
				regName := d.registrantNames[regNum]
				register(t, monitor, regName)
				monitor.registrations[regName].missedHeartbeats = d.registrantMissed[regNum]
			}
			got := monitor.GetStatuses()
			if diff := cmp.Diff(d.want, got); diff != "" {
				t.Errorf("GetStatuses() has unexpected diff (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIncrementAll_shouldIncrementAllMissedHeartbeatsToAtMostThreshold(t *testing.T) {
	testData := []struct {
		name            string
		threshold       int64 // must be < 100
		registrantNames []string
	}{
		{
			name:            "1 threshold",
			threshold:       1,
			registrantNames: []string{"a", "b"},
		},
	}
	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			monitor := defaultMonitor(t)
			monitor.threshold = d.threshold
			numRegistrants := len(d.registrantNames)
			for regNum := 0; regNum < numRegistrants; regNum++ {
				register(t, monitor, d.registrantNames[regNum])
			}
			for i := 0; i < 100; i++ {
				monitor.incrementAll()
			}
			for regNum := 0; regNum < numRegistrants; regNum++ {
				regName := d.registrantNames[regNum]
				got := monitor.registrations[regName].missedHeartbeats
				if got != d.threshold {
					t.Errorf("incrementAll() error: registrants[%s].missedHeartbeats = %v, want %v", regName, got, d.threshold)
				}
			}
		})
	}
}

func TestRegister_shouldReturnFunctionThatResetsRegistrantMissedHeartbeats(t *testing.T) {
	testData := []struct {
		name             string
		registrantNames  []string
		missedHeartbeats []int64
		threshold        int64
	}{
		{
			name:             "2 registrants at 0",
			registrantNames:  []string{"a", "b"},
			missedHeartbeats: []int64{0, 0},
			threshold:        1,
		},
		{
			name:             "2 registrants at threshold",
			registrantNames:  []string{"a", "b"},
			missedHeartbeats: []int64{5, 5},
			threshold:        5,
		},
		{
			name:             "2 registrants less than threshold",
			registrantNames:  []string{"a", "b"},
			missedHeartbeats: []int64{2, 1},
			threshold:        3,
		},
	}
	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			monitor := defaultMonitor(t)
			monitor.threshold = d.threshold
			numRegistrants := len(d.registrantNames)
			specs := make([]*Spec, numRegistrants)
			for i := 0; i < numRegistrants; i++ {
				name := d.registrantNames[i]
				spec := register(t, monitor, name)
				specs[i] = spec
				monitor.registrations[name].missedHeartbeats = d.missedHeartbeats[i]
			}
			for i := 0; i < numRegistrants; i++ {
				name := d.registrantNames[i]
				registrant := monitor.registrations[name]
				spec := specs[i]
				spec.Beat()
				if registrant.missedHeartbeats != 0 {
					t.Errorf("spec.Beat() error, registrants[%v].missedHeartbeats = %v, want 0", name, registrant.missedHeartbeats)
				}
			}
		})
	}
}

func TestSpec_shouldCreateTickerWithTheConfiguredInterval(t *testing.T) {
	testData := []struct {
		name     string
		interval time.Duration
	}{
		{
			name:     "25ms interval",
			interval: time.Millisecond * 25,
		},
		{
			name:     "10ms interval",
			interval: time.Millisecond * 10,
		},
	}
	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			spec := &Spec{
				BeatFunc: func() {},
				Interval: d.interval,
			}
			start := time.Now()
			ticker := spec.CreateTicker()
			<-ticker.C
			end := time.Now()
			defer ticker.Stop()
			delta := end.Sub(start)
			deltaMs := delta.Milliseconds()
			toleranceMs := int64(1)
			absDiff := int64(math.Abs(float64(d.interval.Milliseconds() - deltaMs)))
			if absDiff > toleranceMs {
				t.Errorf("spec.CreateTicker() ms difference from expected = %v, want < %v", deltaMs, toleranceMs)
			}
		})
	}
}

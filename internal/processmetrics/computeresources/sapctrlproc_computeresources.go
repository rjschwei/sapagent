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

package computeresources

import (
	"context"

	mrpb "google.golang.org/genproto/googleapis/monitoring/v3"
	"github.com/GoogleCloudPlatform/sapagent/internal/cloudmonitoring"
	"github.com/GoogleCloudPlatform/sapagent/internal/commandlineexecutor"
	cnfpb "github.com/GoogleCloudPlatform/sapagent/protos/configuration"
	"github.com/GoogleCloudPlatform/sapagent/shared/log"
)

const (
	sapCTRLCPUPath    = "/sap/control/cpu/utilization"
	sapCtrlMemoryPath = "/sap/control/memory/utilization"
)

type (
	// SAPControlProcInstanceProperties have the required context for collecting metrics for cpu
	// and memory per process for SAPControl processes.
	SAPControlProcInstanceProperties struct {
		Config        *cnfpb.Configuration
		Client        cloudmonitoring.TimeSeriesCreator
		Executor      commandlineexecutor.Execute
		NewProcHelper newProcessWithContextHelper
	}
)

// Collect SAP additional metrics like per process CPU and per process memory
// utilization of SAP Control Processes.
func (p *SAPControlProcInstanceProperties) Collect(ctx context.Context) []*mrpb.TimeSeries {
	params := parameters{
		executor:         p.Executor,
		config:           p.Config,
		client:           p.Client,
		cpuMetricPath:    sapCTRLCPUPath,
		memoryMetricPath: sapCtrlMemoryPath,
		newProc:          p.NewProcHelper,
	}
	processes := collectControlProcesses(ctx, params)
	if len(processes) == 0 {
		log.Logger.Debug("Cannot collect CPU and memory per process for Netweaver, empty process list.")
		return nil
	}
	return append(collectCPUPerProcess(ctx, params, processes), collectMemoryPerProcess(ctx, params, processes)...)
}

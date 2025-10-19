package usemetrics

import (
	"testing"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func TestUSEMetricsSLOs(t *testing.T) {
	sloTests := []v1.SLO{
		// Utilization metrics
		ExampleCPUUtilizationSLO(),
		ExampleMemoryUtilizationSLO(),
		ExampleDiskUtilizationSLO(),
		// Saturation metrics
		ExampleCPUSaturationSLO(),
		ExampleMemorySaturationSLO(),
		ExampleDiskSaturationSLO(),
		ExampleNetworkSaturationSLO(),
		// Error metrics
		ExampleDiskErrorsSLO(),
		ExampleNetworkErrorsSLO(),
		ExampleMemoryErrorsSLO(),
		ExampleCPUThrottlingErrorsSLO(),
	}
	for _, tt := range sloTests {
		if err := tt.Validate(); err != nil {
			t.Errorf("slo.Validate() error for %s: (%s)", tt.Metadata.Name, err.Error())
		}
	}
}

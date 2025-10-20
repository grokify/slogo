package usemetrics

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func SLOs() []v1.SLO {
	return []v1.SLO{
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
}

func SLOsBySetSlug() map[string][]v1.SLO {
	return map[string][]v1.SLO{
		"errors": {
			ExampleDiskErrorsSLO(),
			ExampleNetworkErrorsSLO(),
			ExampleMemoryErrorsSLO(),
			ExampleCPUThrottlingErrorsSLO()},
		"saturation": {
			ExampleCPUSaturationSLO(),
			ExampleMemorySaturationSLO(),
			ExampleDiskSaturationSLO(),
			ExampleNetworkSaturationSLO()},
		"utilization": {
			ExampleCPUUtilizationSLO(),
			ExampleMemoryUtilizationSLO(),
			ExampleDiskUtilizationSLO()},
	}
}

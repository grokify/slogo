package redmetrics

import v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"

func SLOs() []v1.SLO {
	return []v1.SLO{
		ExampleRateSLO(),
		ExampleErrorRateSLO(),
		ExampleDurationSLO(),
		ExampleDurationP99SLO(),
		ExampleAvailabilitySLO(),
	}
}

func SLOsBySetSlug() map[string][]v1.SLO {
	return map[string][]v1.SLO{
		"availability": {
			ExampleAvailabilitySLO()},
		"duration": {
			ExampleDurationSLO(),
			ExampleDurationP99SLO()},
		"error": {
			ExampleErrorRateSLO()},
		"rate": {
			ExampleRateSLO()},
	}
}

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

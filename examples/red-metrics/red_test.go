package redmetrics

import (
	"testing"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func TestREDMetricsSLOs(t *testing.T) {
	sloTests := []v1.SLO{
		ExampleRateSLO(),
		ExampleErrorRateSLO(),
		ExampleDurationSLO(),
		ExampleDurationP99SLO(),
		ExampleAvailabilitySLO(),
	}
	for _, tt := range sloTests {
		if err := tt.Validate(); err != nil {
			t.Errorf("slo.Validate() error for %s: (%s)", tt.Metadata.Name, err.Error())
		}
	}
}

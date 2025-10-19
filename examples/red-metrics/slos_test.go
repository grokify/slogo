package redmetrics

import (
	"testing"
)

func TestREDMetricsSLOs(t *testing.T) {
	sloTests := SLOs()

	for _, tt := range sloTests {
		if err := tt.Validate(); err != nil {
			t.Errorf("slo.Validate() error for %s: (%s)", tt.Metadata.Name, err.Error())
		}
	}
}

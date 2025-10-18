package treatlowtrafficasequallyimportant

import (
	"testing"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
)

func TestSLOs(t *testing.T) {
	sloTests := []v1.SLO{
		ExampleLowTrafficSLO(),
	}
	for _, tt := range sloTests {
		if err := tt.Validate(); err != nil {
			t.Errorf("slo.Validate() error: (%s)", err.Error())
		}
	}
}

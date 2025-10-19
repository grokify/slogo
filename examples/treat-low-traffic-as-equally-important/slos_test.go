package treatlowtrafficasequallyimportant

import (
	"testing"
)

func TestSLOs(t *testing.T) {
	sloTests := SLOs()

	for _, tt := range sloTests {
		if err := tt.Validate(); err != nil {
			t.Errorf("slo.Validate() error: (%s)", err.Error())
		}
	}
}

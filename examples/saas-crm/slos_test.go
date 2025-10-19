package saascrm

import (
	"testing"
)

func TestSaaSCRMSLOs(t *testing.T) {
	sloTests := SLOs()

	for _, tt := range sloTests {
		if err := tt.Validate(); err != nil {
			t.Errorf("slo.Validate() error for %s: (%s)", tt.Metadata.Name, err.Error())
		}
	}
}

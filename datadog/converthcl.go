package datadog

/*
import (
	"errors"
	"fmt"
	"io"

	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// ToHCL takes an OpenSLO SLO and writes the equivalent
// Datadog Terraform HCL to the provided io.Writer.
func ToHCL(slo *v1.SLO, w io.Writer) error {
	if slo == nil {
		return errors.New("`v1.SLO` cannot be nil")
	}
	f := hclwrite.NewEmptyFile()
	body := f.Body()

	// Normalize resource name
	resourceName := slo.Metadata.Name
	if resourceName == "" {
		resourceName = "slo_auto_generated"
	}

	// Create a resource block
	r := body.AppendNewBlock("resource", []string{"datadog_service_level_objective", resourceName})
	rb := r.Body()

	// Basic attributes
	rb.SetAttributeValue("name", cty.StringVal(slo.Metadata.DisplayName))
	rb.SetAttributeValue("type", cty.StringVal("metric"))

	// Convert OpenSLO target (e.g. 0.99 → 99)
	if slo.Spec.Target != nil {
		targetPercent := float64(*slo.Spec.Target) * 100
		rb.SetAttributeValue("target", cty.NumberFloatVal(targetPercent))
	}

	// Handle metric indicator
	if slo.Spec.Indicator != nil && slo.Spec.Indicator.MetricSource != nil {
		ms := slo.Spec.Indicator.MetricSource
		q := rb.AppendNewBlock("query", nil).Body()

		// Default example queries
		numerator := "sum:http.requests{status:2xx}.as_count()"
		denominator := "sum:http.requests{*}.as_count()"

		if ms.Spec != nil {
			if num, ok := ms.Spec["numerator"].(string); ok {
				numerator = num
			}
			if denom, ok := ms.Spec["denominator"].(string); ok {
				denominator = denom
			}
		}

		q.SetAttributeValue("numerator", cty.StringVal(numerator))
		q.SetAttributeValue("denominator", cty.StringVal(denominator))
	}

	// Write to the provided writer
	if _, err := w.Write(f.Bytes()); err != nil {
		return fmt.Errorf("write hcl: %w", err)
	}

	return nil
}
*/

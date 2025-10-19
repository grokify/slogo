package redmetrics

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleErrorRateSLO is a SLO that measures the error rate of requests.
// Error rate is the "E" in RED metrics - tracking the proportion of failed requests.
func ExampleErrorRateSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "api-error-rate",
			DisplayName: "API Error Rate",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkRED,
				ontology.LabelLayer:      ontology.LayerService,
				ontology.LabelScope:      ontology.ScopeCustomerFacing,
				ontology.LabelAudience:   ontology.AudienceSRE,
				ontology.LabelCategory:   ontology.CategoryQuality,
				ontology.LabelSeverity:   ontology.SeverityCritical,
				ontology.LabelTier:       ontology.TierP0,
				ontology.LabelMetricType: ontology.MetricTypeErrors,
			})},
		v1.SLOSpec{
			Description: "Track error rate to ensure service reliability - successful responses (2xx, 3xx) vs errors (4xx, 5xx)",
			Service:     "api-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "error-rate",
					DisplayName: "Percentage of successful requests",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(http_requests_total{status=~\"2..|3..\"}[5m]))", // Successful requests
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(http_requests_total[5m]))", // Total requests
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(4, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "High reliability",
					Target:          pointer.Pointer(float64(0.999)),      // 99.9% success rate
					TimeSliceTarget: pointer.Pointer(float64(0.95)),       // 95% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
				{
					DisplayName:     "Acceptable reliability",
					Target:          pointer.Pointer(float64(0.99)),       // 99% success rate
					TimeSliceTarget: pointer.Pointer(float64(0.90)),       // 90% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

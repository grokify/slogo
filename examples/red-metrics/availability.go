package redmetrics

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleAvailabilitySLO is a derived SLO that measures service availability.
// Availability is calculated from Rate (R) and Error (E) metrics in RED - combining
// whether the service is handling requests AND those requests are successful.
func ExampleAvailabilitySLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "api-availability",
			DisplayName: "API Service Availability",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkRED,
				ontology.LabelLayer:      ontology.LayerService,
				ontology.LabelScope:      ontology.ScopeCustomerFacing,
				ontology.LabelAudience:   ontology.AudienceSRE,
				ontology.LabelCategory:   ontology.CategoryAvailability,
				ontology.LabelSeverity:   ontology.SeverityCritical,
				ontology.LabelTier:       ontology.TierP0,
				ontology.LabelMetricType: ontology.MetricTypeAvailability,
			})},
		v1.SLOSpec{
			Description: "Track overall service availability derived from request rate and error rate - ensures service is both reachable and responding successfully",
			Service:     "api-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "service-availability",
					DisplayName: "Service availability percentage",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									// Successful requests (2xx, 3xx) from error metric
									slogo.AttrQuery: "sum(rate(http_requests_total{status=~\"2..|3..\"}[5m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									// Total requests from rate metric
									slogo.AttrQuery: "sum(rate(http_requests_total[5m]))",
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
					DisplayName:     "Four nines availability",
					Target:          pointer.Pointer(float64(0.9999)), // 99.99% availability
					TimeSliceTarget: pointer.Pointer(float64(0.99)),   // 99% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
				{
					DisplayName:     "Three nines availability",
					Target:          pointer.Pointer(float64(0.999)), // 99.9% availability
					TimeSliceTarget: pointer.Pointer(float64(0.95)),  // 95% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

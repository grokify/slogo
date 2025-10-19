package treatlowtrafficasequallyimportant

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleLowTrafficSLO is a SLO that reats low traffic as equally important.
func ExampleLowTrafficSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "service-availability",
			DisplayName: "SLO for this important service",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerService,
				ontology.LabelScope:      ontology.ScopeCustomerFacing,
				ontology.LabelAudience:   ontology.AudienceSRE,
				ontology.LabelCategory:   ontology.CategoryAvailability,
				ontology.LabelSeverity:   ontology.SeverityCritical,
				ontology.LabelTier:       ontology.TierP0,
				ontology.LabelMetricType: ontology.MetricTypeAvailability,
			})},
		v1.SLOSpec{
			Description: "SLO to see if we are above our SLA and to react before.",
			Service:     "web-availability",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "web-availability",
					DisplayName: "Time when our service was running.",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Any", // Any service with needed metrics.
								Spec: map[string]any{ // Fields necessary to query service for the data.
									slogo.AttrQuery: "Any", // 'query' is just an example field.
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Any", // Any service with needed metrics.
								Spec: map[string]any{ // Fields necessary to query service for the data.
									slogo.AttrQuery: "Any", // 'query' is just an example field.
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(1, v1.DurationShorthandUnitWeek),
				IsRolling: false,
				Calendar: &v1.SLOCalendar{
					StartTime: "2022-01-01 12:00:00",
					TimeZone:  "America/New_York",
				},
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "Objective",
					Target:          pointer.Pointer(float64(0.9995)),
					TimeSliceTarget: pointer.Pointer(float64(0.95)),
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitMinute)),
					// Operator:        v1.OperatorGT,
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

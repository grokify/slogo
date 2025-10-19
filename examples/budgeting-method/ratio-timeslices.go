package budgetingmethod

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleRatioTimeSlicesSLO is a SLO that compares success requests with response 2xx to total requests to the main page of our site.
func ExampleRatioTimeSlicesSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "main-page",
			DisplayName: "Main page availability",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerService,
				ontology.LabelScope:      ontology.ScopeCustomerFacing,
				ontology.LabelAudience:   ontology.AudienceSRE,
				ontology.LabelCategory:   ontology.CategoryAvailability,
				ontology.LabelSeverity:   ontology.SeverityCritical,
				ontology.LabelTier:       ontology.TierP0,
				ontology.LabelDomain:     ontology.DomainEcommerce,
				ontology.LabelMetricType: ontology.MetricTypeAvailability,
			})},
		v1.SLOSpec{
			Description: "Our main page should be available always",
			Service:     "web-shop",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "response-code",
					DisplayName: "Response codes of requests to main page",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{ // In most cases it is easier to think about timeslices in ratio metrics.
							MetricSource: v1.SLIMetricSource{
								Type: "Any", // Here put any service that holds information you need.
								Spec: map[string]any{ // Fields necessary to query service for the data.
									slogo.AttrQuery: "Any", // 'query' is just an example field.
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Any", // Here put any service that holds information you need.
								Spec: map[string]any{ // Fields necessary to query service for the data.
									slogo.AttrQuery: "Any", // 'query' is just an example field.
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(2, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodRatioTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "Good",
					Target:          pointer.Pointer(float64(0.99)),
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitMinute)),
					// Operator:        v1.OperatorGT,
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

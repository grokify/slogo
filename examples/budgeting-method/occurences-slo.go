package budgetingmethod

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleOccurencesSLO is a SLO that measures the time of searching for an item in an online shop.
func ExampleOccurencesSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "search-slo",
			DisplayName: "Searching time",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerService,
				ontology.LabelScope:      ontology.ScopeCustomerFacing,
				ontology.LabelAudience:   ontology.AudienceSRE,
				ontology.LabelCategory:   ontology.CategoryLatency,
				ontology.LabelSeverity:   ontology.SeverityHigh,
				ontology.LabelTier:       ontology.TierP1,
				ontology.LabelDomain:     ontology.DomainEcommerce,
				ontology.LabelMetricType: ontology.MetricTypeDuration,
			})},
		v1.SLOSpec{
			Description: "Regardless of the number of parallel searches it never should be more than 500ms.",
			Service:     "web-shop",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "search-latency",
					DisplayName: "Latency of user searches",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Any", // Any service with needed metrics.
							Spec: map[string]any{ // Fields necessary to query service for the data.
								slogo.AttrQuery: "Any", // 'query' is just an example field.
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(2, v1.DurationShorthandUnitWeek), // Two weeks is a good choice in most cases.
				IsRolling: true,                                                     // Rolling timewindow give us better view if service starts to working better.
			}},
			BudgetingMethod: v1.SLOBudgetingMethodOccurrences, // As said in README.md, occurrences are great for measuring all searches the same way.
			Objectives: []v1.SLOObjective{
				{
					DisplayName: "Good experience",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(500)),
					Target:      pointer.Pointer(float64(0.99)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

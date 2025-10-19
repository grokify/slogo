package aiagents

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleAgentResponseQualitySLO measures the quality of agent responses based on user feedback.
func ExampleAgentResponseQualitySLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "agent-response-quality",
			DisplayName: "AI Agent Response Quality",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerBusiness,
				ontology.LabelScope:      ontology.ScopeCustomerFacing,
				ontology.LabelAudience:   ontology.AudienceProduct,
				ontology.LabelCategory:   ontology.CategoryQuality,
				ontology.LabelSeverity:   ontology.SeverityHigh,
				ontology.LabelTier:       ontology.TierP0,
				ontology.LabelDomain:     ontology.DomainAIML,
				ontology.LabelMetricType: ontology.MetricTypeSatisfaction,
			})},
		v1.SLOSpec{
			Description: "Ensure AI agents provide high-quality responses based on user satisfaction ratings",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "response-satisfaction-rate",
					DisplayName: "Percentage of responses rated positively by users",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_response_ratings_total{rating=~\"good|excellent\"}[10m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_response_ratings_total[10m]))",
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
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "High user satisfaction",
					Target:          pointer.Pointer(float64(0.85)), // 85% positive ratings
					TimeSliceTarget: pointer.Pointer(float64(0.90)), // 90% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(10, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExamplePerUserResponseQualitySLO tracks response quality on a per-user basis.
func ExamplePerUserResponseQualitySLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "per-user-response-quality",
			DisplayName: "Per-User Response Quality",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerBusiness,
				ontology.LabelScope:      ontology.ScopeCustomerFacing,
				ontology.LabelAudience:   ontology.AudienceProduct,
				ontology.LabelCategory:   ontology.CategoryQuality,
				ontology.LabelSeverity:   ontology.SeverityHigh,
				ontology.LabelTier:       ontology.TierP0,
				ontology.LabelDomain:     ontology.DomainAIML,
				ontology.LabelMetricType: ontology.MetricTypeSatisfaction,
			})},
		v1.SLOSpec{
			Description: "Monitor response quality per user to identify users with poor experience",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "user-satisfaction-rate",
					DisplayName: "Percentage of users with good response quality",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(agent_response_ratings_total{rating=~\"good|excellent\"}[1d])) / sum by(user_id) (rate(agent_response_ratings_total[1d])) >= 0.80)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(agent_response_ratings_total[1d])))",
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(1, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "Consistent quality across users",
					Target:          pointer.Pointer(float64(0.90)), // 90% of users have >=80% positive ratings
					TimeSliceTarget: pointer.Pointer(float64(0.85)), // 85% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleAgentAccuracySLO measures the accuracy of agent responses.
func ExampleAgentAccuracySLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "agent-accuracy",
			DisplayName: "AI Agent Response Accuracy",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerBusiness,
				ontology.LabelScope:      ontology.ScopeCustomerFacing,
				ontology.LabelAudience:   ontology.AudienceProduct,
				ontology.LabelCategory:   ontology.CategoryQuality,
				ontology.LabelSeverity:   ontology.SeverityCritical,
				ontology.LabelTier:       ontology.TierP0,
				ontology.LabelDomain:     ontology.DomainAIML,
				ontology.LabelMetricType: ontology.MetricTypeSatisfaction,
			})},
		v1.SLOSpec{
			Description: "Track the accuracy of agent responses by measuring hallucination rate and factual correctness",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "response-accuracy-rate",
					DisplayName: "Percentage of accurate responses",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_responses_total[10m])) - sum(rate(agent_hallucinations_total[10m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_responses_total[10m]))",
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(1, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "High accuracy",
					Target:          pointer.Pointer(float64(0.95)), // 95% accurate responses
					TimeSliceTarget: pointer.Pointer(float64(0.90)), // 90% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(10, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

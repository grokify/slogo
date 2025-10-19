package aiagents

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleTaskCompletionRateSLO measures how often agents successfully complete user tasks.
func ExampleTaskCompletionRateSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "task-completion-rate",
			DisplayName: "Agent Task Completion Rate",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerBusiness,
				ontology.LabelScope:      ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:   ontology.AudienceProduct,
				ontology.LabelCategory:   ontology.CategoryQuality,
				ontology.LabelSeverity:   ontology.SeverityHigh,
				ontology.LabelTier:       ontology.TierP0,
				ontology.LabelDomain:     ontology.DomainAIML,
				ontology.LabelMetricType: ontology.MetricTypeEfficiency,
			})},
		v1.SLOSpec{
			Description: "Track the percentage of user tasks that agents successfully complete without errors or abandonment",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "task-success-rate",
					DisplayName: "Percentage of successfully completed tasks",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_tasks_total{status=\"completed\"}[10m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_tasks_total[10m]))",
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
					DisplayName:     "High task completion",
					Target:          pointer.Pointer(float64(0.90)), // 90% task completion rate
					TimeSliceTarget: pointer.Pointer(float64(0.85)), // 85% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(10, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExamplePerUserTaskCompletionSLO tracks task completion on a per-user basis.
func ExamplePerUserTaskCompletionSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "per-user-task-completion",
			DisplayName: "Per-User Task Completion Rate",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerBusiness,
				ontology.LabelScope:      ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:   ontology.AudienceProduct,
				ontology.LabelCategory:   ontology.CategoryQuality,
				ontology.LabelSeverity:   ontology.SeverityHigh,
				ontology.LabelTier:       ontology.TierP1,
				ontology.LabelDomain:     ontology.DomainAIML,
				ontology.LabelMetricType: ontology.MetricTypeEfficiency,
			})},
		v1.SLOSpec{
			Description: "Monitor task completion per user to identify users struggling with the agent",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "user-task-success",
					DisplayName: "Percentage of users with high task completion",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(agent_tasks_total{status=\"completed\"}[1d])) / sum by(user_id) (rate(agent_tasks_total[1d])) >= 0.85)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(agent_tasks_total[1d])))",
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
					DisplayName:     "Consistent user success",
					Target:          pointer.Pointer(float64(0.90)), // 90% of users have >=85% completion rate
					TimeSliceTarget: pointer.Pointer(float64(0.85)), // 85% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleTaskAbandonmentRateSLO measures how often users abandon tasks.
func ExampleTaskAbandonmentRateSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "task-abandonment-rate",
			DisplayName: "Agent Task Abandonment Rate",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerBusiness,
				ontology.LabelScope:      ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:   ontology.AudienceProduct,
				ontology.LabelCategory:   ontology.CategoryQuality,
				ontology.LabelSeverity:   ontology.SeverityCritical,
				ontology.LabelTier:       ontology.TierP0,
				ontology.LabelDomain:     ontology.DomainAIML,
				ontology.LabelMetricType: ontology.MetricTypeEfficiency,
			})},
		v1.SLOSpec{
			Description: "Track task abandonment rate to identify user frustration and agent limitations",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "task-retention-rate",
					DisplayName: "Percentage of tasks not abandoned",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_tasks_total[10m])) - sum(rate(agent_tasks_total{status=\"abandoned\"}[10m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_tasks_total[10m]))",
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
					DisplayName:     "Low abandonment rate",
					Target:          pointer.Pointer(float64(0.95)), // 95% of tasks not abandoned (5% abandonment)
					TimeSliceTarget: pointer.Pointer(float64(0.90)), // 90% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(10, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleMultiStepTaskSuccessSLO measures success rate for complex multi-step tasks.
func ExampleMultiStepTaskSuccessSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "multi-step-task-success",
			DisplayName: "Multi-Step Task Success Rate",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:  ontology.FrameworkCustom,
				ontology.LabelLayer:      ontology.LayerBusiness,
				ontology.LabelScope:      ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:   ontology.AudienceProduct,
				ontology.LabelCategory:   ontology.CategoryQuality,
				ontology.LabelSeverity:   ontology.SeverityHigh,
				ontology.LabelTier:       ontology.TierP1,
				ontology.LabelDomain:     ontology.DomainAIML,
				ontology.LabelMetricType: ontology.MetricTypeEfficiency,
			})},
		v1.SLOSpec{
			Description: "Track success rate for complex tasks requiring multiple agent actions",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "complex-task-success",
					DisplayName: "Multi-step task completion rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_tasks_total{type=\"multi_step\",status=\"completed\"}[10m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_tasks_total{type=\"multi_step\"}[10m]))",
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
					DisplayName:     "Complex task reliability",
					Target:          pointer.Pointer(float64(0.85)), // 85% completion for complex tasks
					TimeSliceTarget: pointer.Pointer(float64(0.80)), // 80% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(10, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

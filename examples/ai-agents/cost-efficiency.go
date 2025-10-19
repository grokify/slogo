package aiagents

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleTokenUsagePerTaskSLO measures token efficiency per completed task.
func ExampleTokenUsagePerTaskSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "token-usage-per-task",
			DisplayName: "Token Usage Efficiency"},
		v1.SLOSpec{
			Description: "Monitor average token usage per completed task to ensure cost efficiency",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "avg-tokens-per-task",
					DisplayName: "Average tokens consumed per task",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "sum(rate(agent_tokens_consumed_total[1h])) / sum(rate(agent_tasks_total{status=\"completed\"}[1h]))",
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(1, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodOccurrences,
			Objectives: []v1.SLOObjective{
				{
					DisplayName: "Efficient token usage",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(5000)), // Less than 5000 tokens per task
					Target:      pointer.Pointer(float64(0.90)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExamplePerUserCostSLO measures cost efficiency on a per-user basis.
func ExamplePerUserCostSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "per-user-cost",
			DisplayName: "Per-User Cost Efficiency"},
		v1.SLOSpec{
			Description: "Track average cost per user to ensure sustainable economics",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "user-cost-distribution",
					DisplayName: "Percentage of users within cost budget",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(agent_cost_usd_total[1d])) < 5)", // Users costing < $5/day
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(agent_cost_usd_total[1d])))",
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
					DisplayName:     "Cost within budget",
					Target:          pointer.Pointer(float64(0.95)), // 95% of users under budget
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleCostPerSuccessfulTaskSLO measures cost efficiency for successful outcomes.
func ExampleCostPerSuccessfulTaskSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "cost-per-successful-task",
			DisplayName: "Cost per Successful Task"},
		v1.SLOSpec{
			Description: "Monitor cost per successfully completed task to optimize ROI",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "successful-task-cost",
					DisplayName: "Average cost per successful task in USD",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "sum(rate(agent_cost_usd_total[1h])) / sum(rate(agent_tasks_total{status=\"completed\"}[1h]))",
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(1, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodOccurrences,
			Objectives: []v1.SLOObjective{
				{
					DisplayName: "Sustainable unit economics",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(0.50)), // Less than $0.50 per successful task
					Target:      pointer.Pointer(float64(0.90)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleCacheHitRateSLO measures efficiency of response caching.
func ExampleCacheHitRateSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "cache-hit-rate",
			DisplayName: "Agent Response Cache Hit Rate"},
		v1.SLOSpec{
			Description: "Track cache hit rate to reduce costs and improve response times",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "cache-efficiency",
					DisplayName: "Percentage of cacheable requests served from cache",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_cache_hits_total[5m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_cache_requests_total[5m]))",
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
					DisplayName:     "High cache efficiency",
					Target:          pointer.Pointer(float64(0.70)), // 70% cache hit rate
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

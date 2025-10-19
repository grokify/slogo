package aiagents

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleAgentResponseTimeSLO measures the response time for agent interactions.
func ExampleAgentResponseTimeSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "agent-response-time",
			DisplayName: "AI Agent Response Time"},
		v1.SLOSpec{
			Description: "Ensure AI agents respond quickly to user queries for good user experience",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "agent-p95-latency",
					DisplayName: "P95 response time for agent queries",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "histogram_quantile(0.95, rate(agent_response_duration_seconds_bucket[5m]))",
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
					DisplayName: "Fast response time",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(3)), // 3 seconds P95 latency
					Target:      pointer.Pointer(float64(0.95)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExamplePerUserResponseTimeSLO tracks response time on a per-user basis.
func ExamplePerUserResponseTimeSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "per-user-response-time",
			DisplayName: "Per-User Response Time"},
		v1.SLOSpec{
			Description: "Monitor response time per user to ensure consistent performance",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "user-response-time",
					DisplayName: "Percentage of users with acceptable response times",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(histogram_quantile(0.95, sum by(user_id, le) (rate(agent_response_duration_seconds_bucket[1h]))) < 3)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(agent_response_duration_seconds_count[1h])))",
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
					DisplayName:     "Consistent performance",
					Target:          pointer.Pointer(float64(0.95)), // 95% of users have P95 < 3s
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitHour)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleAgentFirstTokenLatencySLO measures time to first token for streaming responses.
func ExampleAgentFirstTokenLatencySLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "agent-first-token-latency",
			DisplayName: "AI Agent First Token Latency"},
		v1.SLOSpec{
			Description: "Track time to first token in streaming responses for perceived responsiveness",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "first-token-latency",
					DisplayName: "P95 time to first token",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "histogram_quantile(0.95, rate(agent_first_token_duration_seconds_bucket[5m]))",
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
					DisplayName: "Responsive streaming",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(0.5)), // 500ms P95 to first token
					Target:      pointer.Pointer(float64(0.95)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

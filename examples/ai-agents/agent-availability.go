package aiagents

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleAgentAvailabilitySLO measures the availability of AI agent services for end users.
// This tracks whether users can successfully start and interact with AI agents.
func ExampleAgentAvailabilitySLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "agent-availability",
			DisplayName: "AI Agent Service Availability"},
		v1.SLOSpec{
			Description: "Ensure AI agents are available when users need them - successful agent sessions vs total attempts",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "agent-session-success-rate",
					DisplayName: "Successful agent session start rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_session_starts_total{status=\"success\"}[5m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(agent_session_starts_total[5m]))",
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
					DisplayName:     "High availability",
					Target:          pointer.Pointer(float64(0.999)),      // 99.9% availability
					TimeSliceTarget: pointer.Pointer(float64(0.95)),       // 95% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExamplePerUserAgentAvailabilitySLO measures agent availability for individual users.
func ExamplePerUserAgentAvailabilitySLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "per-user-agent-availability",
			DisplayName: "Per-User Agent Availability"},
		v1.SLOSpec{
			Description: "Track agent availability per user to ensure consistent service across all customers",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "user-agent-success-rate",
					DisplayName: "Per-user agent session success rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(agent_session_starts_total{status=\"success\"}[1h])) / sum by(user_id) (rate(agent_session_starts_total[1h])) >= 0.99)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(agent_session_starts_total[1h])))",
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
					DisplayName:     "Consistent user experience",
					Target:          pointer.Pointer(float64(0.95)),       // 95% of users should have >=99% success rate
					TimeSliceTarget: pointer.Pointer(float64(0.90)),       // 90% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitHour)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

package aiagents

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleDailyActiveUsersSLO measures daily active users engaging with agents.
func ExampleDailyActiveUsersSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "daily-active-users",
			DisplayName: "Daily Active Users",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework: ontology.FrameworkCustom,
				ontology.LabelLayer:     ontology.LayerBusiness,
				ontology.LabelScope:     ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:  ontology.AudienceProduct,
				ontology.LabelCategory:  ontology.CategoryEngagement,
				ontology.LabelSeverity:  ontology.SeverityHigh,
				ontology.LabelTier:      ontology.TierP1,
				ontology.LabelDomain:    ontology.DomainAIML,
			})},
		v1.SLOSpec{
			Description: "Track daily active users to monitor platform adoption and engagement",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "dau-count",
					DisplayName: "Daily active user count",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "count(count_over_time(agent_user_activity_total[24h]) > 0)",
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(1, v1.DurationShorthandUnitMonth),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodOccurrences,
			Objectives: []v1.SLOObjective{
				{
					DisplayName: "Healthy DAU growth",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(10000)), // Minimum 10K daily active users
					Target:      pointer.Pointer(float64(0.90)),  // 90% of days meet threshold
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleUserRetentionSLO measures user retention over time.
func ExampleUserRetentionSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "user-retention",
			DisplayName: "7-Day User Retention",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework: ontology.FrameworkCustom,
				ontology.LabelLayer:     ontology.LayerBusiness,
				ontology.LabelScope:     ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:  ontology.AudienceProduct,
				ontology.LabelCategory:  ontology.CategoryEngagement,
				ontology.LabelSeverity:  ontology.SeverityCritical,
				ontology.LabelTier:      ontology.TierP0,
				ontology.LabelDomain:    ontology.DomainAIML,
			})},
		v1.SLOSpec{
			Description: "Track percentage of users who return within 7 days of first use",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "retention-rate",
					DisplayName: "7-day retention rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM returning_users WHERE days_since_first_use <= 7",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM new_users WHERE created_at >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 7 DAY)",
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
					DisplayName:     "Good user retention",
					Target:          pointer.Pointer(float64(0.60)), // 60% 7-day retention
					TimeSliceTarget: pointer.Pointer(float64(0.90)), // 90% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleSessionDurationSLO measures average user session duration.
func ExampleSessionDurationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "session-duration",
			DisplayName: "User Session Duration",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework: ontology.FrameworkCustom,
				ontology.LabelLayer:     ontology.LayerBusiness,
				ontology.LabelScope:     ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:  ontology.AudienceProduct,
				ontology.LabelCategory:  ontology.CategoryEngagement,
				ontology.LabelSeverity:  ontology.SeverityMedium,
				ontology.LabelTier:      ontology.TierP2,
				ontology.LabelDomain:    ontology.DomainAIML,
			})},
		v1.SLOSpec{
			Description: "Track average session duration to measure user engagement depth",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "avg-session-duration",
					DisplayName: "Average session duration in minutes",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "avg(rate(agent_session_duration_seconds_sum[1h]) / rate(agent_session_duration_seconds_count[1h])) / 60",
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
					DisplayName: "Engaged sessions",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(5)), // Average session >= 5 minutes
					Target:      pointer.Pointer(float64(0.80)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleConversationTurnsSLO measures the number of turns in agent conversations.
func ExampleConversationTurnsSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "conversation-turns",
			DisplayName: "Agent Conversation Depth",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework: ontology.FrameworkCustom,
				ontology.LabelLayer:     ontology.LayerBusiness,
				ontology.LabelScope:     ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:  ontology.AudienceProduct,
				ontology.LabelCategory:  ontology.CategoryEngagement,
				ontology.LabelSeverity:  ontology.SeverityMedium,
				ontology.LabelTier:      ontology.TierP2,
				ontology.LabelDomain:    ontology.DomainAIML,
			})},
		v1.SLOSpec{
			Description: "Track average conversation turns to measure engagement and agent helpfulness",
			Service:     "ai-agent-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "avg-conversation-turns",
					DisplayName: "Average turns per conversation",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "avg(rate(agent_conversation_turns_sum[1h]) / rate(agent_conversation_turns_count[1h]))",
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
					DisplayName: "Meaningful conversations",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(3)), // At least 3 turns per conversation
					Target:      pointer.Pointer(float64(0.75)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

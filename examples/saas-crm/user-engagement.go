package saascrm

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleDailyActiveUsersSLO measures daily active users.
func ExampleDailyActiveUsersSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "daily-active-users",
			DisplayName: "Daily Active Users (DAU)",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkCustom,
				ontology.LabelLayer:        ontology.LayerBusiness,
				ontology.LabelScope:        ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:     ontology.AudienceProduct,
				ontology.LabelCategory:     ontology.CategoryEngagement,
				ontology.LabelSeverity:     ontology.SeverityHigh,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelDomain:       ontology.DomainCRM,
				ontology.LabelJourneyStage: ontology.JourneyStageEngagement,
			})},
		v1.SLOSpec{
			Description: "Track daily active users to monitor platform engagement",
			Service:     "saas-crm-platform",
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
								slogo.AttrQuery: "count(count_over_time(user_activity_total[24h]) > 0)",
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
					DisplayName: "Healthy daily engagement",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(50000)), // Minimum 50K DAU
					Target:      pointer.Pointer(float64(0.90)),  // 90% of days meet threshold
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleMonthlyActiveUsersSLO measures monthly active users.
func ExampleMonthlyActiveUsersSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "monthly-active-users",
			DisplayName: "Monthly Active Users (MAU)",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkCustom,
				ontology.LabelLayer:        ontology.LayerBusiness,
				ontology.LabelScope:        ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:     ontology.AudienceProduct,
				ontology.LabelCategory:     ontology.CategoryEngagement,
				ontology.LabelSeverity:     ontology.SeverityHigh,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelDomain:       ontology.DomainCRM,
				ontology.LabelJourneyStage: ontology.JourneyStageEngagement,
			})},
		v1.SLOSpec{
			Description: "Track monthly active users to monitor overall platform adoption",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "mau-count",
					DisplayName: "Monthly active user count",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "count(count_over_time(user_activity_total[30d]) > 0)",
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(3, v1.DurationShorthandUnitMonth),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodOccurrences,
			Objectives: []v1.SLOObjective{
				{
					DisplayName: "Strong monthly engagement",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(200000)), // Minimum 200K MAU
					Target:      pointer.Pointer(float64(0.90)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleDAUMAURatioSLO measures the DAU/MAU ratio (stickiness).
func ExampleDAUMAURatioSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "dau-mau-ratio",
			DisplayName: "DAU/MAU Ratio (Stickiness)",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkCustom,
				ontology.LabelLayer:        ontology.LayerBusiness,
				ontology.LabelScope:        ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:     ontology.AudienceProduct,
				ontology.LabelCategory:     ontology.CategoryEngagement,
				ontology.LabelSeverity:     ontology.SeverityHigh,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelDomain:       ontology.DomainCRM,
				ontology.LabelMetricType:   ontology.MetricTypeStickiness,
				ontology.LabelJourneyStage: ontology.JourneyStageEngagement,
			})},
		v1.SLOSpec{
			Description: "Track DAU/MAU ratio to measure product stickiness and engagement frequency",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "stickiness-ratio",
					DisplayName: "Ratio of daily to monthly active users",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "count(count_over_time(user_activity_total[24h]) > 0) / count(count_over_time(user_activity_total[30d]) > 0)",
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
					DisplayName: "Good stickiness",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(0.40)), // 40% DAU/MAU ratio
					Target:      pointer.Pointer(float64(0.85)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleWeeklyActiveUsersSLO measures weekly active users.
func ExampleWeeklyActiveUsersSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "weekly-active-users",
			DisplayName: "Weekly Active Users (WAU)"},
		v1.SLOSpec{
			Description: "Track weekly active users for mid-range engagement metrics",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "wau-count",
					DisplayName: "Weekly active user count",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "count(count_over_time(user_activity_total[7d]) > 0)",
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(2, v1.DurationShorthandUnitMonth),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodOccurrences,
			Objectives: []v1.SLOObjective{
				{
					DisplayName: "Healthy weekly engagement",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(120000)), // Minimum 120K WAU
					Target:      pointer.Pointer(float64(0.90)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleDAUWAURatioSLO measures the DAU/WAU ratio.
func ExampleDAUWAURatioSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "dau-wau-ratio",
			DisplayName: "DAU/WAU Ratio"},
		v1.SLOSpec{
			Description: "Track DAU/WAU ratio to measure weekly engagement intensity",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "weekly-stickiness",
					DisplayName: "Ratio of daily to weekly active users",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "count(count_over_time(user_activity_total[24h]) > 0) / count(count_over_time(user_activity_total[7d]) > 0)",
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
					DisplayName: "Strong weekly stickiness",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(0.60)), // 60% DAU/WAU ratio
					Target:      pointer.Pointer(float64(0.85)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExamplePowerUserRatioSLO measures percentage of power users (daily active for 20+ days per month).
func ExamplePowerUserRatioSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "power-user-ratio",
			DisplayName: "Power User Ratio"},
		v1.SLOSpec{
			Description: "Track percentage of users who are highly engaged (active 20+ days per month)",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "power-user-percentage",
					DisplayName: "Percentage of power users among MAU",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM user_activity WHERE days_active_last_30d >= 20",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM user_activity WHERE days_active_last_30d > 0",
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(1, v1.DurationShorthandUnitMonth),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "Healthy power user base",
					Target:          pointer.Pointer(float64(0.25)),       // 25% of MAU are power users
					TimeSliceTarget: pointer.Pointer(float64(0.85)),       // 85% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

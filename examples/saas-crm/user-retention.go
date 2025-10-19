package saascrm

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleDay7RetentionSLO measures 7-day user retention.
func ExampleDay7RetentionSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "day7-retention",
			DisplayName: "7-Day User Retention",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkCustom,
				ontology.LabelLayer:        ontology.LayerBusiness,
				ontology.LabelScope:        ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:     ontology.AudienceProduct,
				ontology.LabelCategory:     ontology.CategoryConversion,
				ontology.LabelSeverity:     ontology.SeverityCritical,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelDomain:       ontology.DomainCRM,
				ontology.LabelMetricType:   ontology.MetricTypeRetention,
				ontology.LabelJourneyStage: ontology.JourneyStageRetention,
			})},
		v1.SLOSpec{
			Description: "Track percentage of users who return on day 7 after signup",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "week-1-retention",
					DisplayName: "Day 7 retention rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM user_retention WHERE days_since_signup = 7 AND was_active = true",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM users WHERE signup_date = DATE_SUB(CURRENT_DATE(), INTERVAL 7 DAY)",
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
					DisplayName:     "Strong early retention",
					Target:          pointer.Pointer(float64(0.50)), // 50% return on day 7
					TimeSliceTarget: pointer.Pointer(float64(0.80)), // 80% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleDay30RetentionSLO measures 30-day user retention.
func ExampleDay30RetentionSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "day30-retention",
			DisplayName: "30-Day User Retention",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkCustom,
				ontology.LabelLayer:        ontology.LayerBusiness,
				ontology.LabelScope:        ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:     ontology.AudienceProduct,
				ontology.LabelCategory:     ontology.CategoryConversion,
				ontology.LabelSeverity:     ontology.SeverityCritical,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelDomain:       ontology.DomainCRM,
				ontology.LabelMetricType:   ontology.MetricTypeRetention,
				ontology.LabelJourneyStage: ontology.JourneyStageRetention,
			})},
		v1.SLOSpec{
			Description: "Track percentage of users who return on day 30 after signup",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "month-1-retention",
					DisplayName: "Day 30 retention rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM user_retention WHERE days_since_signup = 30 AND was_active = true",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM users WHERE signup_date = DATE_SUB(CURRENT_DATE(), INTERVAL 30 DAY)",
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(3, v1.DurationShorthandUnitMonth),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "Good monthly retention",
					Target:          pointer.Pointer(float64(0.40)), // 40% return on day 30
					TimeSliceTarget: pointer.Pointer(float64(0.80)), // 80% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleChurnRateSLO measures monthly customer churn rate.
func ExampleChurnRateSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "monthly-churn-rate",
			DisplayName: "Monthly Customer Churn Rate",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkCustom,
				ontology.LabelLayer:        ontology.LayerBusiness,
				ontology.LabelScope:        ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:     ontology.AudienceProduct,
				ontology.LabelCategory:     ontology.CategoryConversion,
				ontology.LabelSeverity:     ontology.SeverityCritical,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelDomain:       ontology.DomainCRM,
				ontology.LabelMetricType:   ontology.MetricTypeRetention,
				ontology.LabelJourneyStage: ontology.JourneyStageRetention,
			})},
		v1.SLOSpec{
			Description: "Track monthly churn rate to ensure retention goals are met",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "customer-retention-rate",
					DisplayName: "Monthly retention rate (1 - churn)",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "BigQuery",
							Spec: map[string]any{
								slogo.AttrQuery: "SELECT 1 - (churned_customers / total_customers_start_of_month) FROM monthly_churn_metrics WHERE month = CURRENT_DATE()",
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
					DisplayName: "Low churn rate",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(0.97)), // 97% retention (3% churn)
					Target:      pointer.Pointer(float64(0.90)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleResurrectionRateSLO measures the rate of inactive users returning.
func ExampleResurrectionRateSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "user-resurrection-rate",
			DisplayName: "User Resurrection Rate",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkCustom,
				ontology.LabelLayer:        ontology.LayerBusiness,
				ontology.LabelScope:        ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:     ontology.AudienceProduct,
				ontology.LabelCategory:     ontology.CategoryConversion,
				ontology.LabelSeverity:     ontology.SeverityCritical,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelDomain:       ontology.DomainCRM,
				ontology.LabelMetricType:   ontology.MetricTypeRetention,
				ontology.LabelJourneyStage: ontology.JourneyStageRetention,
			})},
		v1.SLOSpec{
			Description: "Track percentage of previously inactive users (30+ days) who return",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "resurrection-percentage",
					DisplayName: "Monthly user resurrection rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM user_activity WHERE last_active_date <= DATE_SUB(CURRENT_DATE(), INTERVAL 60 DAY) AND current_active_date >= DATE_SUB(CURRENT_DATE(), INTERVAL 30 DAY)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM users WHERE last_active_date <= DATE_SUB(CURRENT_DATE(), INTERVAL 60 DAY)",
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
					DisplayName:     "Healthy resurrection rate",
					Target:          pointer.Pointer(float64(0.10)), // 10% of dormant users return
					TimeSliceTarget: pointer.Pointer(float64(0.75)), // 75% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleCohortRetentionSLO measures retention for specific user cohorts.
func ExampleCohortRetentionSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "cohort-retention-90d",
			DisplayName: "90-Day Cohort Retention",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkCustom,
				ontology.LabelLayer:        ontology.LayerBusiness,
				ontology.LabelScope:        ontology.ScopeBusinessOutcome,
				ontology.LabelAudience:     ontology.AudienceProduct,
				ontology.LabelCategory:     ontology.CategoryConversion,
				ontology.LabelSeverity:     ontology.SeverityCritical,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelDomain:       ontology.DomainCRM,
				ontology.LabelMetricType:   ontology.MetricTypeRetention,
				ontology.LabelJourneyStage: ontology.JourneyStageRetention,
			})},
		v1.SLOSpec{
			Description: "Track percentage of users from each cohort still active after 90 days",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "cohort-90d-retention",
					DisplayName: "90-day cohort retention rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM user_cohorts WHERE cohort_date = DATE_SUB(CURRENT_DATE(), INTERVAL 90 DAY) AND days_active_in_last_30d > 0",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM user_cohorts WHERE cohort_date = DATE_SUB(CURRENT_DATE(), INTERVAL 90 DAY)",
								},
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(3, v1.DurationShorthandUnitMonth),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodTimeslices,
			Objectives: []v1.SLOObjective{
				{
					DisplayName:     "Strong long-term retention",
					Target:          pointer.Pointer(float64(0.35)), // 35% still active after 90 days
					TimeSliceTarget: pointer.Pointer(float64(0.80)), // 80% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

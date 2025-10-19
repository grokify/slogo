package saascrm

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleUserActivationSLO measures the percentage of new users who complete activation steps.
func ExampleUserActivationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "user-activation-rate",
			DisplayName: "New User Activation Rate"},
		v1.SLOSpec{
			Description: "Track percentage of new users who complete core activation steps within first 7 days",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "activation-completion",
					DisplayName: "New user activation completion rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM users WHERE activated = true AND activation_date <= DATE_ADD(signup_date, INTERVAL 7 DAY)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM users WHERE signup_date >= DATE_SUB(CURRENT_DATE(), INTERVAL 7 DAY)",
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
					DisplayName:     "Strong activation rate",
					Target:          pointer.Pointer(float64(0.60)), // 60% of new users activate
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleTimeToFirstValueSLO measures how quickly users achieve first value.
func ExampleTimeToFirstValueSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "time-to-first-value",
			DisplayName: "Time to First Value"},
		v1.SLOSpec{
			Description: "Track how quickly new users create their first contact or deal",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "fast-first-value",
					DisplayName: "Percentage of users with first value within 24 hours",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM user_milestones WHERE milestone = 'first_contact_created' AND TIMESTAMP_DIFF(milestone_time, signup_time, HOUR) <= 24",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "BigQuery",
								Spec: map[string]any{
									slogo.AttrQuery: "SELECT COUNT(DISTINCT user_id) FROM users WHERE signup_date >= DATE_SUB(CURRENT_DATE(), INTERVAL 7 DAY)",
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
					DisplayName:     "Fast time to value",
					Target:          pointer.Pointer(float64(0.50)), // 50% create first contact in 24h
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleOnboardingCompletionSLO measures completion of onboarding flow.
func ExampleOnboardingCompletionSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "onboarding-completion",
			DisplayName: "Onboarding Completion Rate"},
		v1.SLOSpec{
			Description: "Track percentage of users who complete the full onboarding checklist",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "onboarding-complete",
					DisplayName: "Onboarding completion rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(onboarding_progress{status=\"completed\"})",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(onboarding_progress)",
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
					DisplayName:     "High onboarding completion",
					Target:          pointer.Pointer(float64(0.70)), // 70% complete onboarding
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

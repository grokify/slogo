package saascrm

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleContactManagementUsageSLO measures adoption of contact management features.
func ExampleContactManagementUsageSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "contact-management-usage",
			DisplayName: "Contact Management Feature Usage"},
		v1.SLOSpec{
			Description: "Track percentage of active users who regularly use contact management features",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "contact-feature-adoption",
					DisplayName: "Weekly contact management adoption rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(feature_usage_total{feature=\"contacts\"}[7d])) > 0)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(count_over_time(user_activity_total[7d]) > 0)",
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
					DisplayName:     "Strong contact feature adoption",
					Target:          pointer.Pointer(float64(0.80)), // 80% of active users use contacts
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleDealPipelineUsageSLO measures adoption of deal pipeline features.
func ExampleDealPipelineUsageSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "deal-pipeline-usage",
			DisplayName: "Deal Pipeline Feature Usage"},
		v1.SLOSpec{
			Description: "Track percentage of active users managing deals in the pipeline",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "deal-feature-adoption",
					DisplayName: "Weekly deal pipeline adoption rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(feature_usage_total{feature=\"deals\"}[7d])) > 0)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(count_over_time(user_activity_total[7d]) > 0)",
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
					DisplayName:     "Good deal pipeline adoption",
					Target:          pointer.Pointer(float64(0.60)), // 60% of active users use deals
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleEmailIntegrationUsageSLO measures adoption of email integration.
func ExampleEmailIntegrationUsageSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "email-integration-usage",
			DisplayName: "Email Integration Usage"},
		v1.SLOSpec{
			Description: "Track percentage of users with active email integration",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "email-integration-adoption",
					DisplayName: "Email integration adoption rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(email_integration_status{status=\"active\"})",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(count_over_time(user_activity_total[7d]) > 0)",
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
					DisplayName:     "Strong email integration",
					Target:          pointer.Pointer(float64(0.70)), // 70% have email connected
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleReportingUsageSLO measures usage of reporting and analytics features.
func ExampleReportingUsageSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "reporting-usage",
			DisplayName: "Reporting & Analytics Usage"},
		v1.SLOSpec{
			Description: "Track percentage of active users who view reports or dashboards",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "reporting-adoption",
					DisplayName: "Weekly reporting feature adoption rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(feature_usage_total{feature=~\"reports|dashboards\"}[7d])) > 0)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(count_over_time(user_activity_total[7d]) > 0)",
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
					DisplayName:     "Good reporting adoption",
					Target:          pointer.Pointer(float64(0.45)), // 45% of users view reports weekly
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleMobileAppUsageSLO measures mobile app adoption.
func ExampleMobileAppUsageSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "mobile-app-usage",
			DisplayName: "Mobile App Usage"},
		v1.SLOSpec{
			Description: "Track percentage of active users who use the mobile app",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "mobile-adoption",
					DisplayName: "Monthly mobile app adoption rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(sum by(user_id) (rate(user_activity_total{platform=\"mobile\"}[30d])) > 0)",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "count(count_over_time(user_activity_total[30d]) > 0)",
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
					DisplayName:     "Growing mobile adoption",
					Target:          pointer.Pointer(float64(0.35)), // 35% use mobile app monthly
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

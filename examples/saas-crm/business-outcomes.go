package saascrm

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleDealsCreatedSLO measures the rate of deal creation.
func ExampleDealsCreatedSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "deals-created-rate",
			DisplayName: "Deal Creation Rate"},
		v1.SLOSpec{
			Description: "Track the rate of new deals being created in the system",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "deal-creation-velocity",
					DisplayName: "Daily deal creation count",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "sum(rate(deals_created_total[24h]))",
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
					DisplayName: "Healthy deal creation",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(1000)), // At least 1000 deals created daily
					Target:      pointer.Pointer(float64(0.85)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleDealWinRateSLO measures the percentage of deals that are won.
func ExampleDealWinRateSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "deal-win-rate",
			DisplayName: "Deal Win Rate"},
		v1.SLOSpec{
			Description: "Track percentage of closed deals that are won (vs lost)",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "win-rate-percentage",
					DisplayName: "Win rate for closed deals",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(deals_closed_total{status=\"won\"}[7d]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "sum(rate(deals_closed_total[7d]))",
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
					DisplayName:     "Healthy win rate",
					Target:          pointer.Pointer(float64(0.30)),       // 30% win rate
					TimeSliceTarget: pointer.Pointer(float64(0.85)),       // 85% of time slices meet target
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(1, v1.DurationShorthandUnitDay)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleSalesCycleLengthSLO measures the average sales cycle duration.
func ExampleSalesCycleLengthSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "sales-cycle-length",
			DisplayName: "Average Sales Cycle Length"},
		v1.SLOSpec{
			Description: "Track average time from deal creation to close (won or lost)",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "avg-sales-cycle-days",
					DisplayName: "Average sales cycle in days",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "avg(rate(deal_cycle_duration_days_sum[7d]) / rate(deal_cycle_duration_days_count[7d]))",
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
					DisplayName: "Efficient sales cycle",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(45)), // Less than 45 days average
					Target:      pointer.Pointer(float64(0.80)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleContactCreationRateSLO measures the rate of new contacts being added.
func ExampleContactCreationRateSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "contact-creation-rate",
			DisplayName: "Contact Creation Rate"},
		v1.SLOSpec{
			Description: "Track the rate of new contacts being added to the CRM",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "contact-growth-velocity",
					DisplayName: "Daily contact creation count",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "sum(rate(contacts_created_total[24h]))",
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
					DisplayName: "Healthy contact growth",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(5000)), // At least 5000 contacts created daily
					Target:      pointer.Pointer(float64(0.85)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleEmailsSentPerUserSLO measures average emails sent per active user.
func ExampleEmailsSentPerUserSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "emails-sent-per-user",
			DisplayName: "Emails Sent Per Active User"},
		v1.SLOSpec{
			Description: "Track average number of emails sent per active user to measure engagement",
			Service:     "saas-crm-platform",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "email-activity-per-user",
					DisplayName: "Average emails sent per DAU",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "sum(rate(emails_sent_total[24h])) / count(count_over_time(user_activity_total[24h]) > 0)",
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
					DisplayName: "Active email engagement",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(10)), // At least 10 emails per user per day
					Target:      pointer.Pointer(float64(0.80)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

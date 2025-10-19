package redmetrics

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleRateSLO is a SLO that measures the request rate (requests per second) for an API service.
// Rate is the "R" in RED metrics - tracking how many requests the service handles.
func ExampleRateSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "api-request-rate",
			DisplayName: "API Request Rate"},
		v1.SLOSpec{
			Description: "Monitor API request rate to ensure the service handles expected traffic volume",
			Service:     "api-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "request-rate",
					DisplayName: "Requests per second",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus", // Common monitoring service for RED metrics
							Spec: map[string]any{
								slogo.AttrQuery: "rate(http_requests_total[5m])", // Example Prometheus query for request rate
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
					DisplayName: "Healthy request rate",
					Operator:    v1.OperatorGTE,
					Value:       pointer.Pointer(float64(100)), // Minimum expected requests per second
					Target:      pointer.Pointer(float64(0.99)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

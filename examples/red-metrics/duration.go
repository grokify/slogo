package redmetrics

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleDurationSLO is a SLO that measures request duration (latency).
// Duration is the "D" in RED metrics - tracking how long requests take to complete.
func ExampleDurationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "api-latency",
			DisplayName: "API Request Duration"},
		v1.SLOSpec{
			Description: "Monitor request latency to ensure fast response times for users",
			Service:     "api-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "request-duration",
					DisplayName: "P95 latency of API requests",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))", // P95 latency
							},
						},
					},
				},
			},
			TimeWindow: []v1.SLOTimeWindow{{
				Duration:  v1.NewDurationShorthand(2, v1.DurationShorthandUnitWeek),
				IsRolling: true,
			}},
			BudgetingMethod: v1.SLOBudgetingMethodOccurrences,
			Objectives: []v1.SLOObjective{
				{
					DisplayName: "Fast response time",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(200)), // 200ms P95 latency
					Target:      pointer.Pointer(float64(0.95)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleDurationP99SLO is a SLO that measures P99 latency for more stringent requirements.
func ExampleDurationP99SLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "api-latency-p99",
			DisplayName: "API P99 Duration"},
		v1.SLOSpec{
			Description: "Monitor P99 latency to catch tail latency issues",
			Service:     "api-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "request-duration-p99",
					DisplayName: "P99 latency of API requests",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m]))", // P99 latency
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
					DisplayName: "P99 latency target",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(500)), // 500ms P99 latency
					Target:      pointer.Pointer(float64(0.99)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

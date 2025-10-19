package usemetrics

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleCPUUtilizationSLO is a SLO that measures CPU utilization.
// Utilization is the "U" in USE metrics - tracking the average time a resource is busy.
func ExampleCPUUtilizationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "cpu-utilization",
			DisplayName: "CPU Utilization"},
		v1.SLOSpec{
			Description: "Monitor CPU utilization to ensure resources are not over-utilized",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "cpu-usage",
					DisplayName: "CPU usage percentage",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "avg(rate(node_cpu_seconds_total{mode!=\"idle\"}[5m])) * 100",
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
					DisplayName: "Healthy CPU utilization",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(80)), // Keep CPU below 80%
					Target:      pointer.Pointer(float64(0.99)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleMemoryUtilizationSLO is a SLO that measures memory utilization.
func ExampleMemoryUtilizationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "memory-utilization",
			DisplayName: "Memory Utilization"},
		v1.SLOSpec{
			Description: "Monitor memory utilization to prevent out-of-memory conditions",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "memory-usage",
					DisplayName: "Memory usage percentage",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100",
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
					DisplayName: "Healthy memory utilization",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(85)), // Keep memory below 85%
					Target:      pointer.Pointer(float64(0.99)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleDiskUtilizationSLO is a SLO that measures disk utilization.
func ExampleDiskUtilizationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "disk-utilization",
			DisplayName: "Disk Utilization"},
		v1.SLOSpec{
			Description: "Monitor disk space utilization to prevent storage exhaustion",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "disk-usage",
					DisplayName: "Disk usage percentage",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "(1 - (node_filesystem_avail_bytes{mountpoint=\"/\"} / node_filesystem_size_bytes{mountpoint=\"/\"})) * 100",
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
					DisplayName: "Healthy disk utilization",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(75)), // Keep disk below 75%
					Target:      pointer.Pointer(float64(0.99)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

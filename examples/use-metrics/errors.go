package usemetrics

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
)

// ExampleDiskErrorsSLO is a SLO that measures disk I/O errors.
// Errors is the "E" in USE metrics - tracking error events for resources.
func ExampleDiskErrorsSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "disk-errors",
			DisplayName: "Disk I/O Errors"},
		v1.SLOSpec{
			Description: "Monitor disk I/O errors to detect failing or degraded storage devices",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "disk-error-rate",
					DisplayName: "Rate of disk I/O errors",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "rate(node_disk_io_errors_total[5m])",
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
					DisplayName: "Minimal disk errors",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(0.01)), // Less than 0.01 errors per second
					Target:      pointer.Pointer(float64(0.99)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleNetworkErrorsSLO is a SLO that measures network errors.
func ExampleNetworkErrorsSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "network-errors",
			DisplayName: "Network Errors"},
		v1.SLOSpec{
			Description: "Monitor network errors including packet drops and transmission errors",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "network-error-rate",
					DisplayName: "Network error rate",
				},
				Spec: v1.SLISpec{
					RatioMetric: &v1.SLIRatioMetric{
						Good: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "rate(node_network_transmit_packets_total[5m]) + rate(node_network_receive_packets_total[5m]) - (rate(node_network_transmit_errs_total[5m]) + rate(node_network_receive_errs_total[5m]) + rate(node_network_transmit_drop_total[5m]) + rate(node_network_receive_drop_total[5m]))",
								},
							},
						},
						Total: &v1.SLIMetricSpec{
							MetricSource: v1.SLIMetricSource{
								Type: "Prometheus",
								Spec: map[string]any{
									slogo.AttrQuery: "rate(node_network_transmit_packets_total[5m]) + rate(node_network_receive_packets_total[5m])",
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
					DisplayName:     "Low network error rate",
					Target:          pointer.Pointer(float64(0.9999)), // 99.99% packets without errors
					TimeSliceWindow: pointer.Pointer(v1.NewDurationShorthand(5, v1.DurationShorthandUnitMinute)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleMemoryErrorsSLO is a SLO that measures memory errors (ECC errors).
func ExampleMemoryErrorsSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "memory-errors",
			DisplayName: "Memory ECC Errors"},
		v1.SLOSpec{
			Description: "Monitor memory ECC errors to detect failing or degraded RAM modules",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "memory-error-rate",
					DisplayName: "Rate of correctable and uncorrectable memory errors",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "rate(node_edac_correctable_errors_total[5m]) + rate(node_edac_uncorrectable_errors_total[5m])",
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
					DisplayName: "Minimal memory errors",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(0.001)), // Less than 0.001 errors per second
					Target:      pointer.Pointer(float64(0.999)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleCPUThrottlingErrorsSLO is a SLO that measures CPU throttling events.
func ExampleCPUThrottlingErrorsSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "cpu-throttling",
			DisplayName: "CPU Throttling Events"},
		v1.SLOSpec{
			Description: "Monitor CPU throttling events in containerized environments to detect resource constraints",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "cpu-throttle-rate",
					DisplayName: "Rate of CPU throttling periods",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "rate(container_cpu_cfs_throttled_periods_total[5m])",
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
					DisplayName: "Minimal CPU throttling",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(1)), // Less than 1 throttle event per second
					Target:      pointer.Pointer(float64(0.95)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

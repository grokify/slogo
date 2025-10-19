package usemetrics

import (
	v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
	"github.com/grokify/mogo/pointer"

	"github.com/grokify/slogo"
	"github.com/grokify/slogo/ontology"
)

// ExampleCPUSaturationSLO is a SLO that measures CPU saturation via load average.
// Saturation is the "S" in USE metrics - tracking the degree to which a resource has extra work it can't service.
func ExampleCPUSaturationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "cpu-saturation",
			DisplayName: "CPU Saturation (Load Average)",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkUSE,
				ontology.LabelLayer:        ontology.LayerInfrastructure,
				ontology.LabelScope:        ontology.ScopeInternal,
				ontology.LabelAudience:     ontology.AudienceSRE,
				ontology.LabelCategory:     ontology.CategoryResource,
				ontology.LabelSeverity:     ontology.SeverityHigh,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelMetricType:   ontology.MetricTypeSaturation,
				ontology.LabelResourceType: ontology.ResourceTypeCPU,
			})},
		v1.SLOSpec{
			Description: "Monitor CPU saturation via load average to detect queuing and resource contention",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "cpu-load",
					DisplayName: "CPU load average per core",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "node_load5 / count(node_cpu_seconds_total{mode=\"idle\"})", // 5-minute load average per CPU core
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
					DisplayName: "Acceptable CPU load",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(1.0)), // Load average should be below 1.0 per core
					Target:      pointer.Pointer(float64(0.95)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleMemorySaturationSLO is a SLO that measures memory saturation via swap usage.
func ExampleMemorySaturationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "memory-saturation",
			DisplayName: "Memory Saturation (Swap Usage)",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkUSE,
				ontology.LabelLayer:        ontology.LayerInfrastructure,
				ontology.LabelScope:        ontology.ScopeInternal,
				ontology.LabelAudience:     ontology.AudienceSRE,
				ontology.LabelCategory:     ontology.CategoryResource,
				ontology.LabelSeverity:     ontology.SeverityCritical,
				ontology.LabelTier:         ontology.TierP0,
				ontology.LabelMetricType:   ontology.MetricTypeSaturation,
				ontology.LabelResourceType: ontology.ResourceTypeMemory,
			})},
		v1.SLOSpec{
			Description: "Monitor memory saturation via swap usage to detect memory pressure",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "swap-usage",
					DisplayName: "Swap usage percentage",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "(1 - (node_memory_SwapFree_bytes / node_memory_SwapTotal_bytes)) * 100",
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
					DisplayName: "Minimal swap usage",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(10)), // Keep swap usage below 10%
					Target:      pointer.Pointer(float64(0.99)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleDiskSaturationSLO is a SLO that measures disk I/O saturation.
func ExampleDiskSaturationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "disk-saturation",
			DisplayName: "Disk I/O Saturation",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkUSE,
				ontology.LabelLayer:        ontology.LayerInfrastructure,
				ontology.LabelScope:        ontology.ScopeInternal,
				ontology.LabelAudience:     ontology.AudienceSRE,
				ontology.LabelCategory:     ontology.CategoryResource,
				ontology.LabelSeverity:     ontology.SeverityHigh,
				ontology.LabelTier:         ontology.TierP1,
				ontology.LabelMetricType:   ontology.MetricTypeSaturation,
				ontology.LabelResourceType: ontology.ResourceTypeDisk,
			})},
		v1.SLOSpec{
			Description: "Monitor disk I/O saturation to detect disk bottlenecks and queuing",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "disk-io-utilization",
					DisplayName: "Disk I/O utilization percentage",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "rate(node_disk_io_time_seconds_total[5m]) * 100", // Percentage of time disk was busy
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
					DisplayName: "Healthy disk I/O",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(80)), // Keep disk I/O utilization below 80%
					Target:      pointer.Pointer(float64(0.95)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

// ExampleNetworkSaturationSLO is a SLO that measures network saturation.
func ExampleNetworkSaturationSLO() v1.SLO {
	return v1.NewSLO(
		v1.Metadata{
			Name:        "network-saturation",
			DisplayName: "Network Saturation",
			Labels: ontology.NewLabels(map[string]string{
				ontology.LabelFramework:    ontology.FrameworkUSE,
				ontology.LabelLayer:        ontology.LayerInfrastructure,
				ontology.LabelScope:        ontology.ScopeInternal,
				ontology.LabelAudience:     ontology.AudienceSRE,
				ontology.LabelCategory:     ontology.CategoryResource,
				ontology.LabelSeverity:     ontology.SeverityMedium,
				ontology.LabelTier:         ontology.TierP1,
				ontology.LabelMetricType:   ontology.MetricTypeSaturation,
				ontology.LabelResourceType: ontology.ResourceTypeNetwork,
			})},
		v1.SLOSpec{
			Description: "Monitor network bandwidth saturation to detect network congestion",
			Service:     "backend-service",
			Indicator: &v1.SLOIndicatorInline{
				Metadata: v1.Metadata{
					Name:        "network-bandwidth-usage",
					DisplayName: "Network bandwidth usage percentage",
				},
				Spec: v1.SLISpec{
					ThresholdMetric: &v1.SLIMetricSpec{
						MetricSource: v1.SLIMetricSource{
							Type: "Prometheus",
							Spec: map[string]any{
								slogo.AttrQuery: "(rate(node_network_transmit_bytes_total[5m]) + rate(node_network_receive_bytes_total[5m])) / 125000000 * 100", // Assuming 1Gbps = 125MB/s
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
					DisplayName: "Healthy network bandwidth",
					Operator:    v1.OperatorLT,
					Value:       pointer.Pointer(float64(70)), // Keep network usage below 70%
					Target:      pointer.Pointer(float64(0.95)),
				},
			},
			AlertPolicies: []v1.SLOAlertPolicy{},
		},
	)
}

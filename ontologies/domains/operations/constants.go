// Package operations provides operations-specific ontology labels and values for SLOs.
//
// This package extends the core ontology with labels specific to IT operations,
// incident management, and infrastructure operations.
package operations

// Label keys specific to operations domain.
const (
	// LabelProcess indicates the operational process.
	LabelProcess = "process"

	// LabelImpact indicates operations-specific business impacts.
	LabelImpact = "impact"

	// LabelCategory provides operations-specific categories.
	LabelCategory = "category"
)

// Labels returns all operations-specific label keys.
func Labels() []string {
	return []string{
		LabelProcess,
		LabelImpact,
		LabelCategory,
	}
}

// Process indicates operational processes.
const (
	ProcessIncidentResponse  = "incident-response"
	ProcessChangeManagement  = "change-management"
	ProcessDeployment        = "deployment"
	ProcessBackup            = "backup"
	ProcessRecovery          = "recovery"
	ProcessScaling           = "scaling"
	ProcessMonitoring        = "monitoring"
	ProcessAlerting          = "alerting"
	ProcessCapacityPlanning  = "capacity-planning"
	ProcessPatchManagement   = "patch-management"
	ProcessConfigManagement  = "config-management"
	ProcessReleaseManagement = "release-management"
)

// Category provides operations-specific categories.
const (
	CategoryIncidentManagement = "incident-management"
	CategoryChangeManagement   = "change-management"
	CategoryCapacityManagement = "capacity-management"
	CategoryAvailability       = "availability"
	CategoryPerformance        = "performance"
	CategoryReliability        = "reliability"
	CategoryRecovery           = "recovery"
)

// Impact indicates operations-specific business impacts.
const (
	// Performance Impact
	ImpactPerformanceDegradation = "performance-degradation"
	ImpactLatencyIncrease        = "latency-increase"
	ImpactThroughputReduction    = "throughput-reduction"
	ImpactScalabilityLimit       = "scalability-limit"
	ImpactResourceExhaustion     = "resource-exhaustion"

	// Integration Impact
	ImpactIntegrationFailure = "integration-failure"
	ImpactDataSyncDelay      = "data-sync-delay"
	ImpactDataInconsistency  = "data-inconsistency"
	ImpactWorkflowDisruption = "workflow-disruption"
	ImpactAPIRateLimit       = "api-rate-limit"

	// Operational Impact
	ImpactOperationalEfficiency = "operational-efficiency"
	ImpactMeanTimeToRepair      = "mean-time-to-repair"
	ImpactMeanTimeToDetect      = "mean-time-to-detect"
	ImpactMeanTimeToRecover     = "mean-time-to-recover"
	ImpactIncidentVolume        = "incident-volume"
	ImpactAlertFatigue          = "alert-fatigue"
	ImpactManualIntervention    = "manual-intervention"
	ImpactToil                  = "toil"

	// Infrastructure Impact
	ImpactInfrastructureHealth = "infrastructure-health"
	ImpactCapacityShortfall    = "capacity-shortfall"
	ImpactBackupFailure        = "backup-failure"
	ImpactRecoveryTime         = "recovery-time"
)

// GetLabelDefinitions returns operations-specific label values for each label type.
func GetLabelDefinitions() map[string][]string {
	return map[string][]string{
		LabelProcess: {
			ProcessIncidentResponse,
			ProcessChangeManagement,
			ProcessDeployment,
			ProcessBackup,
			ProcessRecovery,
			ProcessScaling,
			ProcessMonitoring,
			ProcessAlerting,
			ProcessCapacityPlanning,
			ProcessPatchManagement,
			ProcessConfigManagement,
			ProcessReleaseManagement,
		},
		LabelCategory: {
			CategoryIncidentManagement,
			CategoryChangeManagement,
			CategoryCapacityManagement,
			CategoryAvailability,
			CategoryPerformance,
			CategoryReliability,
			CategoryRecovery,
		},
		LabelImpact: {
			ImpactPerformanceDegradation,
			ImpactLatencyIncrease,
			ImpactThroughputReduction,
			ImpactScalabilityLimit,
			ImpactResourceExhaustion,
			ImpactIntegrationFailure,
			ImpactDataSyncDelay,
			ImpactDataInconsistency,
			ImpactWorkflowDisruption,
			ImpactAPIRateLimit,
			ImpactOperationalEfficiency,
			ImpactMeanTimeToRepair,
			ImpactMeanTimeToDetect,
			ImpactMeanTimeToRecover,
			ImpactIncidentVolume,
			ImpactAlertFatigue,
			ImpactManualIntervention,
			ImpactToil,
			ImpactInfrastructureHealth,
			ImpactCapacityShortfall,
			ImpactBackupFailure,
			ImpactRecoveryTime,
		},
	}
}

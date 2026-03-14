// Package iam provides IAM/IGA-specific ontology labels and values for SLOs.
//
// This package extends the core ontology with labels specific to Identity and
// Access Management (IAM) and Identity Governance and Administration (IGA).
package iam

// Label keys specific to IAM domain.
const (
	// LabelProcess indicates the IAM business process or workflow.
	LabelProcess = "process"

	// LabelAccessModel indicates the access control model.
	LabelAccessModel = "access-model"

	// LabelRiskLevel indicates the level of risk for IAM operations.
	LabelRiskLevel = "risk-level"

	// LabelRiskType indicates the type of risk for IAM operations.
	LabelRiskType = "risk-type"

	// LabelCapability indicates IAM system or process capabilities.
	LabelCapability = "capability"

	// LabelWorkflow indicates specific IAM workflow types.
	LabelWorkflow = "workflow"

	// LabelImpact indicates IAM-specific business impacts.
	LabelImpact = "impact"
)

// Labels returns all IAM-specific label keys.
func Labels() []string {
	return []string{
		LabelProcess,
		LabelAccessModel,
		LabelRiskLevel,
		LabelRiskType,
		LabelCapability,
		LabelWorkflow,
		LabelImpact,
	}
}

// Process indicates IAM-specific business processes.
const (
	// Core IGA processes
	ProcessAccessRequest     = "access-request"
	ProcessApprovalWorkflow  = "approval-workflow"
	ProcessProvisioning      = "provisioning"
	ProcessDeprovisioning    = "deprovisioning"
	ProcessRecertification   = "recertification"
	ProcessAccessReview      = "access-review"
	ProcessAttestation       = "attestation"
	ProcessOnboarding        = "onboarding"
	ProcessOffboarding       = "offboarding"
	ProcessJoinerMoverLeaver = "joiner-mover-leaver"

	// Additional IAM processes
	ProcessDecisionMaking    = "decision-making"
	ProcessTermination       = "termination"
	ProcessAutomation        = "automation"
	ProcessFullLifecycle     = "full-lifecycle"
	ProcessRoleAssignment    = "role-assignment"
	ProcessRBAC              = "rbac"
	ProcessElevatedAccess    = "elevated-access"
	ProcessSecurityReview    = "security-review"
	ProcessSODDetection      = "sod-detection"
	ProcessPolicyEnforcement = "policy-enforcement"
	ProcessCompliance        = "compliance"
	ProcessJoiner            = "joiner"
)

// MetricType provides IAM-specific metric type values.
const (
	MetricTypeResponseTime       = "response-time"
	MetricTypeApprovalTime       = "approval-time"
	MetricTypeProvisioningTime   = "provisioning-time"
	MetricTypeEndToEnd           = "end-to-end"
	MetricTypeCampaignCompletion = "campaign-completion"
	MetricTypeRevocationTime     = "revocation-time"
	MetricTypeDetectionTime      = "detection-time"
)

// AccessModel indicates the access control model.
const (
	AccessModelRoleBased      = "role-based"
	AccessModelRBAC           = "rbac"
	AccessModelABAC           = "abac"
	AccessModelPBAC           = "pbac"
	AccessModelReBAC          = "rebac"
	AccessModelDiscretionary  = "discretionary"
	AccessModelMandatory      = "mandatory"
	AccessModelJustInTime     = "just-in-time"
	AccessModelZeroTrust      = "zero-trust"
	AccessModelLeastPrivilege = "least-privilege"
)

// RiskLevel indicates the level of risk for IAM operations.
const (
	RiskLevelCritical         = "critical"
	RiskLevelHigh             = "high"
	RiskLevelMedium           = "medium"
	RiskLevelLow              = "low"
	RiskLevelPrivilegedAccess = "privileged-access"
)

// RiskType indicates the type of risk for IAM operations.
const (
	RiskTypeSODViolation       = "sod-violation"
	RiskTypeConflictingDuties  = "conflicting-duties"
	RiskTypeExcessiveAccess    = "excessive-access"
	RiskTypeOrphanedAccount    = "orphaned-account"
	RiskTypeStaleAccess        = "stale-access"
	RiskTypePrivilegeCreep     = "privilege-creep"
	RiskTypeUnauthorizedAccess = "unauthorized-access"
)

// Capability indicates IAM system or process capabilities.
const (
	CapabilityAutoProvisioning    = "auto-provisioning"
	CapabilityAutoDeprovisioning  = "auto-deprovisioning"
	CapabilityJustInTime          = "just-in-time"
	CapabilitySelfService         = "self-service"
	CapabilityDelegatedAdmin      = "delegated-admin"
	CapabilityRoleMining          = "role-mining"
	CapabilityAccessCertification = "access-certification"
)

// Workflow indicates specific IAM workflow types.
const (
	WorkflowRequestAcknowledgment = "request-acknowledgment"
	WorkflowApproval              = "approval"
	WorkflowEscalation            = "escalation"
	WorkflowRemediation           = "remediation"
	WorkflowNotification          = "notification"
)

// Impact indicates IAM-specific business impacts.
const (
	ImpactProductivity        = "productivity"
	ImpactFraudPrevention     = "fraud-prevention"
	ImpactInsiderThreat       = "insider-threat"
	ImpactUnauthorizedAccess  = "unauthorized-access"
	ImpactPrivilegeEscalation = "privilege-escalation"
	ImpactComplianceViolation = "compliance-violation"
	ImpactAuditFinding        = "audit-finding"
	ImpactAccessDelay         = "access-delay"
	ImpactSecurityExposure    = "security-exposure"
)

// GetLabelDefinitions returns IAM-specific label values for each label type.
func GetLabelDefinitions() map[string][]string {
	return map[string][]string{
		LabelProcess: {
			ProcessAccessRequest,
			ProcessApprovalWorkflow,
			ProcessProvisioning,
			ProcessDeprovisioning,
			ProcessRecertification,
			ProcessAccessReview,
			ProcessAttestation,
			ProcessOnboarding,
			ProcessOffboarding,
			ProcessJoinerMoverLeaver,
			ProcessDecisionMaking,
			ProcessTermination,
			ProcessAutomation,
			ProcessFullLifecycle,
			ProcessRoleAssignment,
			ProcessRBAC,
			ProcessElevatedAccess,
			ProcessSecurityReview,
			ProcessSODDetection,
			ProcessPolicyEnforcement,
			ProcessCompliance,
			ProcessJoiner,
		},
		LabelAccessModel: {
			AccessModelRoleBased,
			AccessModelRBAC,
			AccessModelABAC,
			AccessModelPBAC,
			AccessModelReBAC,
			AccessModelDiscretionary,
			AccessModelMandatory,
			AccessModelJustInTime,
			AccessModelZeroTrust,
			AccessModelLeastPrivilege,
		},
		LabelRiskLevel: {
			RiskLevelCritical,
			RiskLevelHigh,
			RiskLevelMedium,
			RiskLevelLow,
			RiskLevelPrivilegedAccess,
		},
		LabelRiskType: {
			RiskTypeSODViolation,
			RiskTypeConflictingDuties,
			RiskTypeExcessiveAccess,
			RiskTypeOrphanedAccount,
			RiskTypeStaleAccess,
			RiskTypePrivilegeCreep,
			RiskTypeUnauthorizedAccess,
		},
		LabelCapability: {
			CapabilityAutoProvisioning,
			CapabilityAutoDeprovisioning,
			CapabilityJustInTime,
			CapabilitySelfService,
			CapabilityDelegatedAdmin,
			CapabilityRoleMining,
			CapabilityAccessCertification,
		},
		LabelWorkflow: {
			WorkflowRequestAcknowledgment,
			WorkflowApproval,
			WorkflowEscalation,
			WorkflowRemediation,
			WorkflowNotification,
		},
		LabelImpact: {
			ImpactProductivity,
			ImpactFraudPrevention,
			ImpactInsiderThreat,
			ImpactUnauthorizedAccess,
			ImpactPrivilegeEscalation,
			ImpactComplianceViolation,
			ImpactAuditFinding,
			ImpactAccessDelay,
			ImpactSecurityExposure,
		},
	}
}

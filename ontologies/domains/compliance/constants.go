// Package compliance provides compliance-specific ontology labels and values for SLOs.
//
// This package extends the core ontology with labels specific to regulatory
// compliance, audit, and governance requirements.
package compliance

// Label keys specific to compliance domain.
const (
	// LabelFramework indicates the compliance framework.
	LabelFramework = "compliance-framework"

	// LabelRequirement indicates the specific compliance requirement.
	LabelRequirement = "requirement"

	// LabelImpact indicates compliance-specific business impacts.
	LabelImpact = "impact"

	// LabelCategory provides compliance-specific categories.
	LabelCategory = "category"
)

// Labels returns all compliance-specific label keys.
func Labels() []string {
	return []string{
		LabelFramework,
		LabelRequirement,
		LabelImpact,
		LabelCategory,
	}
}

// Framework indicates the compliance framework.
const (
	FrameworkGDPR     = "gdpr"
	FrameworkSOX      = "sox"
	FrameworkHIPAA    = "hipaa"
	FrameworkPCIDSS   = "pci-dss"
	FrameworkSOC2     = "soc2"
	FrameworkISO27001 = "iso-27001"
	FrameworkNIST     = "nist"
	FrameworkFedRAMP  = "fedramp"
	FrameworkCCPA     = "ccpa"
	FrameworkCOBIT    = "cobit"
)

// Category provides compliance-specific categories.
const (
	CategoryDataPrivacy      = "data-privacy"
	CategoryDataProtection   = "data-protection"
	CategoryAccessControl    = "access-control"
	CategoryAuditLogging     = "audit-logging"
	CategoryIncidentResponse = "incident-response"
	CategoryRiskManagement   = "risk-management"
	CategoryVendorManagement = "vendor-management"
	CategoryChangeManagement = "change-management"
	CategorySecurityControls = "security-controls"
)

// Impact indicates compliance-specific business impacts.
const (
	ImpactComplianceGap       = "compliance-gap"
	ImpactRegulatoryViolation = "regulatory-violation"
	ImpactAuditFinding        = "audit-finding"
	ImpactDataPrivacy         = "data-privacy"
	ImpactGDPRCompliance      = "gdpr-compliance"
	ImpactSOXCompliance       = "sox-compliance"
	ImpactHIPAACompliance     = "hipaa-compliance"
	ImpactPCIDSSCompliance    = "pci-dss-compliance"
	ImpactPenalty             = "penalty"
	ImpactLegalLiability      = "legal-liability"
	ImpactReputationalDamage  = "reputational-damage"
)

// GetLabelDefinitions returns compliance-specific label values for each label type.
func GetLabelDefinitions() map[string][]string {
	return map[string][]string{
		LabelFramework: {
			FrameworkGDPR,
			FrameworkSOX,
			FrameworkHIPAA,
			FrameworkPCIDSS,
			FrameworkSOC2,
			FrameworkISO27001,
			FrameworkNIST,
			FrameworkFedRAMP,
			FrameworkCCPA,
			FrameworkCOBIT,
		},
		LabelCategory: {
			CategoryDataPrivacy,
			CategoryDataProtection,
			CategoryAccessControl,
			CategoryAuditLogging,
			CategoryIncidentResponse,
			CategoryRiskManagement,
			CategoryVendorManagement,
			CategoryChangeManagement,
			CategorySecurityControls,
		},
		LabelImpact: {
			ImpactComplianceGap,
			ImpactRegulatoryViolation,
			ImpactAuditFinding,
			ImpactDataPrivacy,
			ImpactGDPRCompliance,
			ImpactSOXCompliance,
			ImpactHIPAACompliance,
			ImpactPCIDSSCompliance,
			ImpactPenalty,
			ImpactLegalLiability,
			ImpactReputationalDamage,
		},
	}
}

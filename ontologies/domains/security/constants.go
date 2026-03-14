// Package security provides security-specific ontology labels and values for SLOs.
//
// This package extends the core ontology with labels specific to security
// monitoring, threat detection, and security operations.
package security

// Label keys specific to security domain.
const (
	// LabelRiskLevel indicates the level of security risk.
	LabelRiskLevel = "risk-level"

	// LabelRiskType indicates the type of security risk.
	LabelRiskType = "risk-type"

	// LabelThreatType indicates the type of security threat.
	LabelThreatType = "threat-type"

	// LabelImpact indicates security-specific business impacts.
	LabelImpact = "impact"

	// LabelCategory provides security-specific categories.
	LabelCategory = "category"
)

// Labels returns all security-specific label keys.
func Labels() []string {
	return []string{
		LabelRiskLevel,
		LabelRiskType,
		LabelThreatType,
		LabelImpact,
		LabelCategory,
	}
}

// RiskLevel indicates the level of security risk.
const (
	RiskLevelCritical = "critical"
	RiskLevelHigh     = "high"
	RiskLevelMedium   = "medium"
	RiskLevelLow      = "low"
	RiskLevelInfo     = "info"
)

// RiskType indicates the type of security risk.
const (
	RiskTypeDataBreach         = "data-breach"
	RiskTypeUnauthorizedAccess = "unauthorized-access"
	RiskTypeDataExfiltration   = "data-exfiltration"
	RiskTypeMalware            = "malware"
	RiskTypePhishing           = "phishing"
	RiskTypeDDoS               = "ddos"
	RiskTypeInsiderThreat      = "insider-threat"
	RiskTypeSupplyChain        = "supply-chain"
	RiskTypeVulnerability      = "vulnerability"
)

// ThreatType indicates the type of security threat.
const (
	ThreatTypeExternal      = "external"
	ThreatTypeInternal      = "internal"
	ThreatTypeAPT           = "apt"
	ThreatTypeOpportunistic = "opportunistic"
	ThreatTypeTargeted      = "targeted"
)

// Category provides security-specific categories.
const (
	CategoryThreatDetection     = "threat-detection"
	CategoryVulnerability       = "vulnerability"
	CategoryIncidentResponse    = "incident-response"
	CategoryAccessControl       = "access-control"
	CategoryDataProtection      = "data-protection"
	CategoryNetworkSecurity     = "network-security"
	CategoryEndpointSecurity    = "endpoint-security"
	CategoryApplicationSecurity = "application-security"
)

// Impact indicates security-specific business impacts.
const (
	ImpactSecurityRisk         = "security-risk"
	ImpactDataBreach           = "data-breach"
	ImpactDataBreachRisk       = "data-breach-risk"
	ImpactInsiderThreat        = "insider-threat"
	ImpactUnauthorizedAccess   = "unauthorized-access"
	ImpactCompromiseRisk       = "compromise-risk"
	ImpactAttackSurface        = "attack-surface"
	ImpactPrivilegeEscalation  = "privilege-escalation"
	ImpactDataLoss             = "data-loss"
	ImpactDataCorruption       = "data-corruption"
	ImpactDataExfiltration     = "data-exfiltration"
	ImpactMalwareInfection     = "malware-infection"
	ImpactCredentialCompromise = "credential-compromise" //nolint:gosec // G101: Not a credential, just a label value
)

// GetLabelDefinitions returns security-specific label values for each label type.
func GetLabelDefinitions() map[string][]string {
	return map[string][]string{
		LabelRiskLevel: {
			RiskLevelCritical,
			RiskLevelHigh,
			RiskLevelMedium,
			RiskLevelLow,
			RiskLevelInfo,
		},
		LabelRiskType: {
			RiskTypeDataBreach,
			RiskTypeUnauthorizedAccess,
			RiskTypeDataExfiltration,
			RiskTypeMalware,
			RiskTypePhishing,
			RiskTypeDDoS,
			RiskTypeInsiderThreat,
			RiskTypeSupplyChain,
			RiskTypeVulnerability,
		},
		LabelThreatType: {
			ThreatTypeExternal,
			ThreatTypeInternal,
			ThreatTypeAPT,
			ThreatTypeOpportunistic,
			ThreatTypeTargeted,
		},
		LabelCategory: {
			CategoryThreatDetection,
			CategoryVulnerability,
			CategoryIncidentResponse,
			CategoryAccessControl,
			CategoryDataProtection,
			CategoryNetworkSecurity,
			CategoryEndpointSecurity,
			CategoryApplicationSecurity,
		},
		LabelImpact: {
			ImpactSecurityRisk,
			ImpactDataBreach,
			ImpactDataBreachRisk,
			ImpactInsiderThreat,
			ImpactUnauthorizedAccess,
			ImpactCompromiseRisk,
			ImpactAttackSurface,
			ImpactPrivilegeEscalation,
			ImpactDataLoss,
			ImpactDataCorruption,
			ImpactDataExfiltration,
			ImpactMalwareInfection,
			ImpactCredentialCompromise,
		},
	}
}

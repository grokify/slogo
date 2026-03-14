// Package saas provides SaaS-specific ontology labels and values for SLOs.
//
// This package extends the core ontology with labels specific to Software as a
// Service (SaaS) platforms, including multi-tenancy and platform operations.
package saas

// Label keys specific to SaaS domain.
const (
	// LabelTenancy indicates the tenancy model or scope.
	LabelTenancy = "tenancy"

	// LabelImpact indicates SaaS-specific business impacts.
	LabelImpact = "impact"

	// LabelCategory provides SaaS-specific categories.
	LabelCategory = "category"
)

// Labels returns all SaaS-specific label keys.
func Labels() []string {
	return []string{
		LabelTenancy,
		LabelImpact,
		LabelCategory,
	}
}

// Tenancy indicates the tenancy model or scope.
const (
	TenancySingleTenant = "single-tenant"
	TenancyMultiTenant  = "multi-tenant"
	TenancyShared       = "shared"
	TenancyDedicated    = "dedicated"
	TenancyHybrid       = "hybrid"
)

// Category provides SaaS-specific categories.
const (
	CategoryTenantHealth    = "tenant-health"
	CategoryPlatformHealth  = "platform-health"
	CategoryResourceSharing = "resource-sharing"
	CategoryIsolation       = "isolation"
)

// Impact indicates SaaS-specific business impacts.
const (
	// Multi-tenancy Impact
	ImpactTenantIsolation          = "tenant-isolation"
	ImpactCrossTenantLeak          = "cross-tenant-leak"
	ImpactSharedResourceContention = "shared-resource-contention"
	ImpactNoisyNeighbor            = "noisy-neighbor"
	ImpactTenantPerformance        = "tenant-performance"
	ImpactTenantAvailability       = "tenant-availability"

	// Platform Impact
	ImpactPlatformStability   = "platform-stability"
	ImpactPlatformScalability = "platform-scalability"
	ImpactPlatformReliability = "platform-reliability"

	// Service Impact
	ImpactServiceAvailability = "service-availability"
	ImpactServiceDowntime     = "service-downtime"
	ImpactServiceOutage       = "service-outage"
	ImpactServiceDegradation  = "service-degradation"
	ImpactAPIAvailability     = "api-availability"
)

// GetLabelDefinitions returns SaaS-specific label values for each label type.
func GetLabelDefinitions() map[string][]string {
	return map[string][]string{
		LabelTenancy: {
			TenancySingleTenant,
			TenancyMultiTenant,
			TenancyShared,
			TenancyDedicated,
			TenancyHybrid,
		},
		LabelCategory: {
			CategoryTenantHealth,
			CategoryPlatformHealth,
			CategoryResourceSharing,
			CategoryIsolation,
		},
		LabelImpact: {
			ImpactTenantIsolation,
			ImpactCrossTenantLeak,
			ImpactSharedResourceContention,
			ImpactNoisyNeighbor,
			ImpactTenantPerformance,
			ImpactTenantAvailability,
			ImpactPlatformStability,
			ImpactPlatformScalability,
			ImpactPlatformReliability,
			ImpactServiceAvailability,
			ImpactServiceDowntime,
			ImpactServiceOutage,
			ImpactServiceDegradation,
			ImpactAPIAvailability,
		},
	}
}

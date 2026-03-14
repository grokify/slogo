// Package product provides product/growth-specific ontology labels and values for SLOs.
//
// This package extends the core ontology with labels specific to product metrics,
// user journeys, and growth frameworks like AARRR (Pirate Metrics).
package product

// Label keys specific to product domain.
const (
	// LabelJourneyStage indicates where in the user journey this metric applies.
	LabelJourneyStage = "journey-stage"

	// LabelMetricType provides product-specific metric type values.
	LabelMetricType = "metric-type"

	// LabelImpact indicates product-specific business impacts.
	LabelImpact = "impact"

	// LabelCategory provides product-specific categories.
	LabelCategory = "category"
)

// Labels returns all product-specific label keys.
func Labels() []string {
	return []string{
		LabelJourneyStage,
		LabelMetricType,
		LabelImpact,
		LabelCategory,
	}
}

// JourneyStage indicates where in the user journey this metric applies.
// Based on AARRR (Pirate Metrics) framework.
const (
	JourneyStageAcquisition = "acquisition" // User acquisition
	JourneyStageActivation  = "activation"  // User activation/onboarding
	JourneyStageEngagement  = "engagement"  // Ongoing engagement
	JourneyStageRetention   = "retention"   // User retention
	JourneyStageRevenue     = "revenue"     // Revenue generation
	JourneyStageReferral    = "referral"    // User referrals
)

// MetricType provides product-specific metric type values.
const (
	MetricTypeSatisfaction = "satisfaction"
	MetricTypeStickiness   = "stickiness"
	MetricTypeRetention    = "retention"
	MetricTypeActivation   = "activation"
	MetricTypeAdoption     = "adoption"
	MetricTypeEfficiency   = "efficiency"
	MetricTypeEngagement   = "engagement"
	MetricTypeConversion   = "conversion"
	MetricTypeChurn        = "churn"
	MetricTypeNPS          = "nps"
	MetricTypeCES          = "ces"
	MetricTypeCSAT         = "csat"
)

// Category provides product-specific categories.
const (
	CategoryEngagement = "engagement"
	CategoryConversion = "conversion"
	CategoryRetention  = "retention"
	CategoryGrowth     = "growth"
)

// Impact indicates product-specific business impacts.
const (
	// User & Customer Impact
	ImpactUserExperience       = "user-experience"
	ImpactCustomerSatisfaction = "customer-satisfaction"
	ImpactUserProductivity     = "user-productivity"
	ImpactEmployeeExperience   = "employee-experience"
	ImpactCustomerRetention    = "customer-retention"
	ImpactCustomerChurn        = "customer-churn"
	ImpactTimeToValue          = "time-to-value"
	ImpactFeatureAdoption      = "feature-adoption"
	ImpactUserOnboarding       = "user-onboarding"
	ImpactProductStickiness    = "product-stickiness"
)

// GetLabelDefinitions returns product-specific label values for each label type.
func GetLabelDefinitions() map[string][]string {
	return map[string][]string{
		LabelJourneyStage: {
			JourneyStageAcquisition,
			JourneyStageActivation,
			JourneyStageEngagement,
			JourneyStageRetention,
			JourneyStageRevenue,
			JourneyStageReferral,
		},
		LabelMetricType: {
			MetricTypeSatisfaction,
			MetricTypeStickiness,
			MetricTypeRetention,
			MetricTypeActivation,
			MetricTypeAdoption,
			MetricTypeEfficiency,
			MetricTypeEngagement,
			MetricTypeConversion,
			MetricTypeChurn,
			MetricTypeNPS,
			MetricTypeCES,
			MetricTypeCSAT,
		},
		LabelCategory: {
			CategoryEngagement,
			CategoryConversion,
			CategoryRetention,
			CategoryGrowth,
		},
		LabelImpact: {
			ImpactUserExperience,
			ImpactCustomerSatisfaction,
			ImpactUserProductivity,
			ImpactEmployeeExperience,
			ImpactCustomerRetention,
			ImpactCustomerChurn,
			ImpactTimeToValue,
			ImpactFeatureAdoption,
			ImpactUserOnboarding,
			ImpactProductStickiness,
		},
	}
}

package ontology

import v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"

// NewLabels creates a v1.Labels map from a simple key-value map.
// OpenSLO v1.Labels expects map[string]Label format where Label is []string.
func NewLabels(labels map[string]string) v1.Labels {
	result := make(v1.Labels, len(labels))
	for k, v := range labels {
		result[k] = v1.Label{v}
	}
	return result
}

// GetLabelDefinitions returns all possible label values for each label type
func GetLabelDefinitions() map[string][]string {
	return map[string][]string{
		LabelFramework: {
			FrameworkRED,
			FrameworkUSE,
			FrameworkCustom,
		},
		LabelLayer: {
			LayerService,
			LayerInfrastructure,
			LayerBusiness,
			LayerApplication,
		},
		LabelScope: {
			ScopeCustomerFacing,
			ScopeInternal,
			ScopeBusinessOutcome,
		},
		LabelAudience: {
			AudienceSRE,
			AudienceEngineering,
			AudienceProduct,
			AudienceExecutive,
			AudienceCustomerSuccess,
		},
		LabelCategory: {
			CategoryAvailability,
			CategoryLatency,
			CategoryThroughput,
			CategoryQuality,
			CategoryResource,
			CategoryEngagement,
			CategoryConversion,
			CategoryCost,
			CategorySecurity,
			CategoryCompliance,
		},
		LabelSeverity: {
			SeverityCritical,
			SeverityHigh,
			SeverityMedium,
			SeverityLow,
			SeverityInformational,
		},
		LabelTier: {
			TierP0,
			TierP1,
			TierP2,
			TierP3,
		},
		LabelMetricType: {
			MetricTypeRate,
			MetricTypeErrors,
			MetricTypeDuration,
			MetricTypeAvailability,
			MetricTypeUtilization,
			MetricTypeSaturation,
			MetricTypeSatisfaction,
			MetricTypeStickiness,
			MetricTypeRetention,
			MetricTypeActivation,
			MetricTypeAdoption,
			MetricTypeEfficiency,
		},
		LabelResourceType: {
			ResourceTypeCPU,
			ResourceTypeMemory,
			ResourceTypeDisk,
			ResourceTypeNetwork,
			ResourceTypeGPU,
		},
		LabelDomain: {
			DomainAIML,
			DomainCRM,
			DomainSaaS,
			DomainEcommerce,
			DomainFintech,
		},
		LabelJourneyStage: {
			JourneyStageAcquisition,
			JourneyStageActivation,
			JourneyStageEngagement,
			JourneyStageRetention,
			JourneyStageRevenue,
			JourneyStageReferral,
		},
	}
}

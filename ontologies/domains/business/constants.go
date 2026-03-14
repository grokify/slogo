// Package business provides business-specific ontology labels and values for SLOs.
//
// This package extends the core ontology with labels specific to business
// outcomes, financial metrics, and organizational impacts.
package business

// Label keys specific to business domain.
const (
	// LabelDomain indicates the business domain or vertical.
	LabelDomain = "domain"

	// LabelImpact indicates business-specific impacts.
	LabelImpact = "impact"

	// LabelCategory provides business-specific categories.
	LabelCategory = "category"
)

// Labels returns all business-specific label keys.
func Labels() []string {
	return []string{
		LabelDomain,
		LabelImpact,
		LabelCategory,
	}
}

// Domain indicates the business domain or vertical.
const (
	DomainAIML       = "ai-ml"
	DomainCRM        = "crm"
	DomainSaaS       = "saas"
	DomainEcommerce  = "ecommerce"
	DomainFintech    = "fintech"
	DomainHealthcare = "healthcare"
	DomainMedia      = "media"
	DomainGaming     = "gaming"
	DomainLogistics  = "logistics"
	DomainEducation  = "education"
)

// Category provides business-specific categories.
const (
	CategoryRevenue    = "revenue"
	CategoryCost       = "cost"
	CategoryGrowth     = "growth"
	CategoryEfficiency = "efficiency"
	CategoryRisk       = "risk"
	CategoryCustomer   = "customer"
	CategoryMarket     = "market"
)

// Impact indicates business-specific impacts.
const (
	// Financial Impact
	ImpactRevenue         = "revenue"
	ImpactRevenueGrowth   = "revenue-growth"
	ImpactRevenueLoss     = "revenue-loss"
	ImpactOperationalCost = "operational-cost"
	ImpactCostReduction   = "cost-reduction"
	ImpactProfitMargin    = "profit-margin"
	ImpactCashFlow        = "cash-flow"

	// Market Impact
	ImpactMarketShare          = "market-share"
	ImpactBrandReputation      = "brand-reputation"
	ImpactCompetitiveAdvantage = "competitive-advantage"
	ImpactMarketPosition       = "market-position"

	// Organizational Impact
	ImpactBusinessAgility       = "business-agility"
	ImpactInnovationVelocity    = "innovation-velocity"
	ImpactTimeToMarket          = "time-to-market"
	ImpactOperationalExcellence = "operational-excellence"

	// Strategic Impact
	ImpactStrategicAlignment    = "strategic-alignment"
	ImpactDigitalTransformation = "digital-transformation"
	ImpactBusinessContinuity    = "business-continuity"
)

// GetLabelDefinitions returns business-specific label values for each label type.
func GetLabelDefinitions() map[string][]string {
	return map[string][]string{
		LabelDomain: {
			DomainAIML,
			DomainCRM,
			DomainSaaS,
			DomainEcommerce,
			DomainFintech,
			DomainHealthcare,
			DomainMedia,
			DomainGaming,
			DomainLogistics,
			DomainEducation,
		},
		LabelCategory: {
			CategoryRevenue,
			CategoryCost,
			CategoryGrowth,
			CategoryEfficiency,
			CategoryRisk,
			CategoryCustomer,
			CategoryMarket,
		},
		LabelImpact: {
			ImpactRevenue,
			ImpactRevenueGrowth,
			ImpactRevenueLoss,
			ImpactOperationalCost,
			ImpactCostReduction,
			ImpactProfitMargin,
			ImpactCashFlow,
			ImpactMarketShare,
			ImpactBrandReputation,
			ImpactCompetitiveAdvantage,
			ImpactMarketPosition,
			ImpactBusinessAgility,
			ImpactInnovationVelocity,
			ImpactTimeToMarket,
			ImpactOperationalExcellence,
			ImpactStrategicAlignment,
			ImpactDigitalTransformation,
			ImpactBusinessContinuity,
		},
	}
}

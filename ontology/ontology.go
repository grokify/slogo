package ontology

import v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"

// Framework indicates which monitoring methodology is being used
const (
	FrameworkRED    = "red"    // Rate, Errors, Duration
	FrameworkUSE    = "use"    // Utilization, Saturation, Errors
	FrameworkCustom = "custom" // Custom business or domain-specific metrics
)

// Layer indicates the observability layer
const (
	LayerService        = "service"        // Request-driven services (APIs, microservices)
	LayerInfrastructure = "infrastructure" // Hardware and system resources
	LayerBusiness       = "business"       // Business outcomes and metrics
	LayerApplication    = "application"    // Application-level metrics
)

// Scope indicates the impact scope of the metric
const (
	ScopeCustomerFacing  = "customer-facing"  // Metrics affecting end users
	ScopeInternal        = "internal"         // Internal infrastructure metrics
	ScopeBusinessOutcome = "business-outcome" // Business performance metrics
)

// Audience indicates the primary stakeholder for this metric
const (
	AudienceSRE             = "sre"              // Site Reliability Engineering
	AudienceEngineering     = "engineering"      // Engineering teams
	AudienceProduct         = "product"          // Product management
	AudienceExecutive       = "executive"        // Executive leadership
	AudienceCustomerSuccess = "customer-success" // Customer success teams
)

// Category indicates the type of metric being measured
const (
	CategoryAvailability = "availability" // Service uptime and availability
	CategoryLatency      = "latency"      // Response times and duration
	CategoryThroughput   = "throughput"   // Request rate and volume
	CategoryQuality      = "quality"      // Error rates, accuracy, satisfaction
	CategoryResource     = "resource"     // CPU, memory, disk, network
	CategoryEngagement   = "engagement"   // User engagement metrics
	CategoryConversion   = "conversion"   // Activation, retention, churn
	CategoryCost         = "cost"         // Operational costs
	CategorySecurity     = "security"     // Security metrics
	CategoryCompliance   = "compliance"   // Compliance metrics
)

// Severity indicates the impact of an SLO breach
const (
	SeverityCritical = "critical" // Immediate action required
	SeverityHigh     = "high"     // Action required soon
	SeverityMedium   = "medium"   // Should be addressed
	SeverityLow      = "low"      // Nice to fix
)

// Tier indicates the service tier or priority
const (
	TierP0 = "p0" // Highest priority
	TierP1 = "p1" // High priority
	TierP2 = "p2" // Medium priority
	TierP3 = "p3" // Low priority
)

// MetricType provides additional categorization within a domain
const (
	// RED metric types
	MetricTypeRate     = "rate"
	MetricTypeErrors   = "errors"
	MetricTypeDuration = "duration"

	// USE metric types
	MetricTypeUtilization = "utilization"
	MetricTypeSaturation  = "saturation"

	// Business metric types
	MetricTypeSatisfaction = "satisfaction"
	MetricTypeStickiness   = "stickiness"
	MetricTypeRetention    = "retention"
	MetricTypeActivation   = "activation"
	MetricTypeAdoption     = "adoption"
	MetricTypeEfficiency   = "efficiency"
)

// ResourceType specifies the type of resource for USE metrics
const (
	ResourceTypeCPU     = "cpu"
	ResourceTypeMemory  = "memory"
	ResourceTypeDisk    = "disk"
	ResourceTypeNetwork = "network"
	ResourceTypeGPU     = "gpu"
)

// Domain provides domain-specific categorization
const (
	DomainAIML      = "ai-ml"     // AI/ML specific metrics
	DomainCRM       = "crm"       // CRM specific metrics
	DomainSaaS      = "saas"      // SaaS specific metrics
	DomainEcommerce = "ecommerce" // E-commerce metrics
	DomainFintech   = "fintech"   // Financial technology metrics
)

// JourneyStage indicates where in the user journey this metric applies
const (
	JourneyStageAcquisition = "acquisition" // User acquisition
	JourneyStageActivation  = "activation"  // User activation/onboarding
	JourneyStageEngagement  = "engagement"  // Ongoing engagement
	JourneyStageRetention   = "retention"   // User retention
	JourneyStageRevenue     = "revenue"     // Revenue generation
	JourneyStageReferral    = "referral"    // User referrals
)

// Environment indicates the deployment environment
const (
	EnvironmentProduction = "production"
	EnvironmentStaging    = "staging"
	EnvironmentDev        = "dev"
)

// Label keys for consistent labeling across SLOs
const (
	LabelFramework    = "framework"
	LabelLayer        = "layer"
	LabelScope        = "scope"
	LabelAudience     = "audience"
	LabelCategory     = "category"
	LabelSeverity     = "severity"
	LabelTier         = "tier"
	LabelMetricType   = "metric-type"
	LabelResourceType = "resource-type"
	LabelDomain       = "domain"
	LabelJourneyStage = "journey-stage"
	LabelService      = "service"
	LabelTeam         = "team"
	LabelEnvironment  = "environment"
)

// NewLabels creates a v1.Labels map from a simple key-value map.
// OpenSLO v1.Labels expects map[string]Label format where Label is []string.
func NewLabels(labels map[string]string) v1.Labels {
	result := make(v1.Labels, len(labels))
	for k, v := range labels {
		result[k] = v1.Label{v}
	}
	return result
}

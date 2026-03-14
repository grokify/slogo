// Package ontology provides a generic labeling taxonomy for SLOs.
//
// This package contains core labels that are universally applicable to any
// SLO implementation. Domain-specific labels are available in the
// ontologies/domains/ subpackages.
package ontology

// Label keys for consistent labeling across SLOs.
// These are the core, generic labels applicable to any domain.
const (
	// LabelFramework indicates which monitoring methodology is being used.
	LabelFramework = "framework"

	// LabelLayer indicates the observability layer.
	LabelLayer = "layer"

	// LabelScope indicates the impact scope of the metric.
	LabelScope = "scope"

	// LabelAudience indicates the primary stakeholder for this metric.
	LabelAudience = "audience"

	// LabelCategory indicates the type of metric being measured.
	LabelCategory = "category"

	// LabelSeverity indicates the impact of an SLO breach.
	LabelSeverity = "severity"

	// LabelTier indicates the service tier or priority.
	LabelTier = "tier"

	// LabelMetricType provides additional categorization within a domain.
	LabelMetricType = "metric-type"

	// LabelResourceType specifies the type of resource for USE metrics.
	LabelResourceType = "resource-type"

	// LabelService identifies the service being measured.
	LabelService = "service"

	// LabelTeam identifies the team responsible for the service.
	LabelTeam = "team"

	// LabelEnvironment indicates the deployment environment.
	LabelEnvironment = "environment"
)

// Labels returns all core label keys.
func Labels() []string {
	return []string{
		LabelFramework,
		LabelLayer,
		LabelScope,
		LabelAudience,
		LabelCategory,
		LabelSeverity,
		LabelTier,
		LabelMetricType,
		LabelResourceType,
		LabelService,
		LabelTeam,
		LabelEnvironment,
	}
}

// Framework indicates which monitoring methodology is being used.
const (
	FrameworkRED    = "red"    // Rate, Errors, Duration
	FrameworkUSE    = "use"    // Utilization, Saturation, Errors
	FrameworkCustom = "custom" // Custom business or domain-specific metrics
)

// Layer indicates the observability layer.
const (
	LayerService        = "service"        // Request-driven services (APIs, microservices)
	LayerInfrastructure = "infrastructure" // Hardware and system resources
	LayerBusiness       = "business"       // Business outcomes and metrics
	LayerApplication    = "application"    // Application-level metrics
)

// Scope indicates the impact scope of the metric.
const (
	ScopeCustomerFacing  = "customer-facing"  // Metrics affecting end users
	ScopeInternal        = "internal"         // Internal infrastructure metrics
	ScopeBusinessOutcome = "business-outcome" // Business performance metrics
)

// Audience indicates the primary stakeholder for this metric.
const (
	AudienceSRE             = "sre"              // Site Reliability Engineering
	AudienceEngineering     = "engineering"      // Engineering teams
	AudienceProduct         = "product"          // Product management
	AudienceExecutive       = "executive"        // Executive leadership
	AudienceCustomerSuccess = "customer-success" // Customer success teams
)

// Category indicates the type of metric being measured.
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

// Severity indicates the impact of an SLO breach.
const (
	SeverityCritical      = "critical" // Immediate action required
	SeverityHigh          = "high"     // Action required soon
	SeverityMedium        = "medium"   // Should be addressed
	SeverityLow           = "low"      // Nice to fix
	SeverityInformational = "info"     // Informational only, no risk
)

// Tier indicates the service tier or priority.
const (
	TierP0 = "p0" // Highest priority
	TierP1 = "p1" // High priority
	TierP2 = "p2" // Medium priority
	TierP3 = "p3" // Low priority
)

// MetricType provides additional categorization within a domain.
const (
	// RED metric types
	MetricTypeRate         = "rate"
	MetricTypeErrors       = "errors"
	MetricTypeDuration     = "duration"
	MetricTypeAvailability = "availability" // Derived from rate and errors

	// USE metric types
	MetricTypeUtilization = "utilization"
	MetricTypeSaturation  = "saturation"

	// Business/Product metric types
	MetricTypeSatisfaction = "satisfaction"
	MetricTypeEfficiency   = "efficiency"
)

// ResourceType specifies the type of resource for USE metrics.
const (
	ResourceTypeCPU     = "cpu"
	ResourceTypeMemory  = "memory"
	ResourceTypeDisk    = "disk"
	ResourceTypeNetwork = "network"
	ResourceTypeGPU     = "gpu"
)

// Environment indicates the deployment environment.
const (
	EnvironmentProduction = "production"
	EnvironmentStaging    = "staging"
	EnvironmentDev        = "dev"
)

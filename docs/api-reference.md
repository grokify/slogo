# API Reference

SLOGo is documented on pkg.go.dev with full API documentation.

## Package Documentation

| Package | Description |
|---------|-------------|
| [slogo](https://pkg.go.dev/github.com/grokify/slogo) | Root package with Objects type and helpers |
| [slogo/ontology](https://pkg.go.dev/github.com/grokify/slogo/ontology) | Core SLO ontology with 12 labeling dimensions |
| [slogo/alert](https://pkg.go.dev/github.com/grokify/slogo/alert) | Alert condition builders for OpenSLO |
| [slogo/requirement](https://pkg.go.dev/github.com/grokify/slogo/requirement) | Requirement types with rate/duration goals |

## Domain Ontology Packages

| Package | Description |
|---------|-------------|
| [ontologies/domains/iam](https://pkg.go.dev/github.com/grokify/slogo/ontologies/domains/iam) | IAM/IGA labels |
| [ontologies/domains/product](https://pkg.go.dev/github.com/grokify/slogo/ontologies/domains/product) | AARRR journey stages |
| [ontologies/domains/business](https://pkg.go.dev/github.com/grokify/slogo/ontologies/domains/business) | Business verticals |
| [ontologies/domains/saas](https://pkg.go.dev/github.com/grokify/slogo/ontologies/domains/saas) | Multi-tenancy labels |
| [ontologies/domains/security](https://pkg.go.dev/github.com/grokify/slogo/ontologies/domains/security) | Security labels |
| [ontologies/domains/compliance](https://pkg.go.dev/github.com/grokify/slogo/ontologies/domains/compliance) | Regulatory frameworks |
| [ontologies/domains/operations](https://pkg.go.dev/github.com/grokify/slogo/ontologies/domains/operations) | Operational processes |

## Key Types

### Objects

The `Objects` type manages collections of OpenSLO objects:

```go
type Objects struct {
    Services        []openslo.Service
    SLOs            []openslo.SLO
    SLIs            []openslo.SLI
    DataSources     []openslo.DataSource
    AlertPolicies   []openslo.AlertPolicy
    AlertConditions []openslo.AlertCondition
    AlertNotificationTargets []openslo.AlertNotificationTarget
}
```

**Methods:**

- `Add(objs ...any)` - Add OpenSLO objects to the collection
- `WriteJSONFile(filepath string)` - Write as JSON
- `WriteYAMLFile(filepath string)` - Write as YAML

### Ontology Labels

Labels are defined as string constants:

```go
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
    LabelService      = "service"
    LabelTeam         = "team"
    LabelEnvironment  = "environment"
)
```

### Alert Conditions

Build alert conditions with the builder pattern:

```go
cond := alert.NewCondition("high-burn-rate").
    WithBurnRate(14.4, "1h").
    WithSeverity(alert.SeverityCritical).
    Build()
```

### Requirements

Define SLO requirements with targets:

```go
req := requirement.Requirement{
    Name: "API Availability",
    Target: requirement.Target{
        RateGoal:     99.9,
        DurationGoal: "30d",
    },
}
```

## Full Documentation

For complete API documentation, visit:

- [pkg.go.dev/github.com/grokify/slogo](https://pkg.go.dev/github.com/grokify/slogo)

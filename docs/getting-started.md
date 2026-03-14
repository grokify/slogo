# Getting Started

## Installation

```bash
go get github.com/grokify/slogo
```

## Reading SLO Files

```go
import "github.com/grokify/slogo"

// Read from file (auto-detects JSON or YAML)
objs, err := slogo.ReadFile("slo-definition.yaml")
if err != nil {
    log.Fatal(err)
}

// Read from io.Reader
objs, err := slogo.Read(reader)
if err != nil {
    log.Fatal(err)
}
```

## Validating SLO Objects

```go
// Validate all objects
if err := objs.Validate(); err != nil {
    log.Fatal(err)
}
```

## Writing SLO Files

```go
// Write as YAML
err := objs.WriteFileYAML("output.yaml")
if err != nil {
    log.Fatal(err)
}

// Write as JSON
err := objs.WriteFileJSON("output.json")
if err != nil {
    log.Fatal(err)
}
```

## Using Predefined Constants

```go
import "github.com/grokify/slogo"

// Severity levels
severity := slogo.SeverityCritical // "critical"
severity := slogo.SeverityHigh     // "high"
severity := slogo.SeverityMedium   // "medium"
severity := slogo.SeverityLow      // "low"
severity := slogo.SeverityInfo     // "info"

// Attributes
attr := slogo.AttrQuery // "query"
```

## Using the Ontology System

The ontology system provides standardized labels for categorizing SLOs:

```go
import (
    v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
    "github.com/grokify/slogo/ontology"
)

// Create labeled SLO metadata
metadata := v1.Metadata{
    Name:        "api-error-rate",
    DisplayName: "API Error Rate",
    Labels: ontology.NewLabels(map[string]string{
        ontology.LabelFramework:  ontology.FrameworkRED,
        ontology.LabelLayer:      ontology.LayerService,
        ontology.LabelScope:      ontology.ScopeCustomerFacing,
        ontology.LabelAudience:   ontology.AudienceSRE,
        ontology.LabelCategory:   ontology.CategoryQuality,
        ontology.LabelSeverity:   ontology.SeverityCritical,
        ontology.LabelTier:       ontology.TierP0,
    }),
}
```

## Using Domain Ontologies

For domain-specific labels, import the appropriate domain package:

```go
import (
    "github.com/grokify/slogo/ontology"
    "github.com/grokify/slogo/ontologies/domains/business"
    "github.com/grokify/slogo/ontologies/domains/product"
)

// Use domain-specific labels
labels := ontology.NewLabels(map[string]string{
    ontology.LabelFramework:    ontology.FrameworkCustom,
    ontology.LabelLayer:        ontology.LayerBusiness,
    business.LabelDomain:       business.DomainCRM,
    product.LabelJourneyStage:  product.JourneyStageActivation,
})
```

## Creating Alert Conditions

```go
import "github.com/grokify/slogo"

// Create an alert condition with burn rate configuration
condition := slogo.NewAlertCondition(
    "critical-burn-rate",
    slogo.SeverityCritical,
    14.4,    // threshold (14.4x burn rate)
    "1h",    // lookback window
    "5m",    // alert after duration
)
```

## Next Steps

- Learn about the [Ontology System](ontology/overview.md)
- Explore [Example SLOs](examples/index.md)
- See [Domain Packages](ontology/domains.md) for specialized use cases

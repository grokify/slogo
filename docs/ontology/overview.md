# Ontology Overview

The `ontology` package provides a comprehensive labeling system for organizing and categorizing SLOs across multiple dimensions.

## Core Labels

The core ontology contains 12 generic labels that apply to any SLO:

| Label | Description | Example Values |
|-------|-------------|----------------|
| `framework` | Monitoring methodology | `red`, `use`, `custom` |
| `layer` | System layer | `service`, `infrastructure`, `business`, `application` |
| `scope` | Impact scope | `customer-facing`, `internal`, `business-outcome` |
| `audience` | Target audience | `sre`, `engineering`, `product`, `executive` |
| `category` | Metric category | `availability`, `latency`, `throughput`, `quality` |
| `severity` | Alert severity | `critical`, `high`, `medium`, `low` |
| `tier` | Priority tier | `p0`, `p1`, `p2`, `p3` |
| `metric-type` | Type of metric | `rate`, `errors`, `duration`, `utilization` |
| `resource-type` | Resource being measured | `cpu`, `memory`, `disk`, `network` |
| `service` | Service name | Custom value |
| `team` | Owning team | Custom value |
| `environment` | Deployment environment | `production`, `staging`, `development` |

## Usage

```go
import (
    v1 "github.com/OpenSLO/go-sdk/pkg/openslo/v1"
    "github.com/grokify/slogo/ontology"
)

metadata := v1.Metadata{
    Name:        "api-latency-p99",
    DisplayName: "API Latency P99",
    Labels: ontology.NewLabels(map[string]string{
        ontology.LabelFramework:  ontology.FrameworkRED,
        ontology.LabelLayer:      ontology.LayerService,
        ontology.LabelScope:      ontology.ScopeCustomerFacing,
        ontology.LabelAudience:   ontology.AudienceSRE,
        ontology.LabelCategory:   ontology.CategoryLatency,
        ontology.LabelSeverity:   ontology.SeverityHigh,
        ontology.LabelTier:       ontology.TierP1,
        ontology.LabelMetricType: ontology.MetricTypeDuration,
    }),
}
```

## Framework Values

### RED (Request-driven)

For services that handle requests:

- **R**ate - Requests per second
- **E**rrors - Failed requests
- **D**uration - Latency distribution

```go
ontology.LabelFramework: ontology.FrameworkRED
```

### USE (Resource-focused)

For infrastructure resources:

- **U**tilization - Percentage of resource used
- **S**aturation - Queue depth, backlog
- **E**rrors - Error counts

```go
ontology.LabelFramework: ontology.FrameworkUSE
```

### Custom

For business metrics and custom implementations:

```go
ontology.LabelFramework: ontology.FrameworkCustom
```

## Layer Values

| Value | Description | Use Case |
|-------|-------------|----------|
| `service` | Application services | APIs, microservices |
| `infrastructure` | Underlying resources | CPU, memory, disk, network |
| `business` | Business metrics | Revenue, conversion, retention |
| `application` | Application-level | User sessions, feature usage |

## Scope Values

| Value | Description | Example |
|-------|-------------|---------|
| `customer-facing` | Affects end users | API latency, error rates |
| `internal` | Internal operations | Build times, CI/CD |
| `business-outcome` | Business KPIs | Conversion rate, churn |

## Audience Values

| Value | Description | Typical Metrics |
|-------|-------------|-----------------|
| `sre` | Site Reliability Engineers | Availability, latency |
| `engineering` | Development teams | Build success, test coverage |
| `product` | Product managers | Engagement, adoption |
| `executive` | Leadership | Revenue impact, growth |
| `customer-success` | Customer-facing teams | Satisfaction, retention |

## Category Values

| Value | Description |
|-------|-------------|
| `availability` | Uptime, success rate |
| `latency` | Response time, duration |
| `throughput` | Requests per second, rate |
| `quality` | Error rate, accuracy |
| `resource` | CPU, memory, disk |
| `engagement` | DAU, MAU, stickiness |
| `conversion` | Signup, activation |
| `cost` | Efficiency, spend |

## Severity & Tier

### Severity Levels

| Value | Description | Response Time |
|-------|-------------|---------------|
| `critical` | Service down | Immediate |
| `high` | Major degradation | < 1 hour |
| `medium` | Minor impact | < 4 hours |
| `low` | Minimal impact | Next business day |

### Priority Tiers

| Value | Description | Error Budget |
|-------|-------------|--------------|
| `p0` | Mission-critical | 99.99% |
| `p1` | High priority | 99.9% |
| `p2` | Standard | 99.5% |
| `p3` | Best effort | 99% |

## Next Steps

- See [Domain Packages](domains.md) for specialized labels
- Explore [Examples](../examples/index.md) using the ontology

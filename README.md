# SLOGo - Go for Service Level Objectives (SLOs)

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

Go utilities for working with [OpenSLO](https://github.com/OpenSLO/OpenSLO) (Service Level Objectives) specifications.

## Overview

`slogo` is a Go library that provides convenient utilities for reading, writing, and validating OpenSLO specification files. It builds on top of the [OpenSLO Go SDK](https://github.com/OpenSLO/go-sdk) to provide a simpler interface for working with SLO definitions in JSON and YAML formats.

## Features

- Read OpenSLO objects from files or readers (auto-detects JSON/YAML format)
- Write OpenSLO objects to files in JSON or YAML format
- Validate SLO definitions against the OpenSLO specification
- Type-safe Go representations of SLO objects
- Predefined severity and attribute constants
- Comprehensive SLO ontology for metric labeling and categorization
- Production-ready examples for RED, USE, AI Agents, and SaaS metrics

## Installation

```bash
go get github.com/grokify/slogo
```

## Usage

### Reading SLO Files

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

### Validating SLO Objects

```go
// Validate all objects
if err := objs.Validate(); err != nil {
    log.Fatal(err)
}
```

### Writing SLO Files

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

### Using Predefined Constants

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

### Using the Ontology System

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

## Ontology

The [`ontology`](ontology/) package provides a comprehensive labeling system for organizing and categorizing SLOs across multiple dimensions:

- **Frameworks**: RED (Rate/Errors/Duration), USE (Utilization/Saturation/Errors), Custom
- **Layers**: Service, Infrastructure, Business, Application
- **Scopes**: Customer-facing, Internal, Business-outcome
- **Audiences**: SRE, Engineering, Product, Executive, Customer-success
- **Categories**: Availability, Latency, Throughput, Quality, Resource, Engagement, Conversion, Cost
- **Severities**: Critical, High, Medium, Low
- **Metric Types**: Rate, Errors, Duration, Utilization, Saturation, Satisfaction, Stickiness, Retention, etc.
- **Domains**: AI-ML, CRM, SaaS, E-commerce, Fintech
- **Journey Stages**: Acquisition, Activation, Engagement, Retention, Revenue, Referral

This multi-dimensional taxonomy enables effective filtering, querying, and organization of SLOs across different teams and use cases.

## Examples

The [`examples`](examples/) directory contains production-ready SLO examples organized by monitoring framework and use case:

### Infrastructure & Service Monitoring

#### [RED Metrics](examples/red-metrics/) (4 SLOs)
Request-driven service monitoring for APIs and microservices:
- **Rate SLO** - Track request throughput (requests per second)
- **Error Rate SLO** - Monitor success/failure ratio with 99.9% reliability target
- **Duration SLOs** - P95 and P99 latency monitoring for response times

#### [USE Metrics](examples/use-metrics/) (11 SLOs)
Infrastructure resource monitoring with Brendan Gregg's USE methodology:
- **Utilization** - CPU, Memory, Disk space usage
- **Saturation** - CPU load average, swap usage, disk I/O, network bandwidth
- **Errors** - Disk I/O errors, network errors, memory ECC errors, CPU throttling

### Business & Product Metrics

#### [AI Agents](examples/ai-agents/) (20 SLOs)
Comprehensive monitoring for AI agent platforms with both aggregated and per-user metrics:
- **Availability** - Service uptime and per-user consistency
- **Quality** - User satisfaction, accuracy, hallucination tracking
- **Performance** - Response time, first-token latency
- **Task Success** - Completion rates, abandonment, multi-step tasks
- **Engagement** - DAU, retention, session duration, conversation depth
- **Cost Efficiency** - Token usage, per-user costs, cache hit rates

#### [SaaS CRM](examples/saas-crm/) (25 SLOs)
End-to-end user journey metrics for CRM platforms (Salesforce, HubSpot):
- **Activation** - User onboarding, time to first value, activation rates
- **Engagement** - **DAU, MAU, DAU/MAU ratio (stickiness)**, WAU, power users
- **Feature Adoption** - Contact management, deal pipeline, email integration, reporting, mobile app
- **Business Outcomes** - Deal creation, win rates, sales cycle length, contact growth
- **Retention** - Day 7/30 retention, churn rate, cohort analysis, user resurrection

### Methodology Examples

#### [Budgeting Methods](examples/budgeting-method/)
Compares different budgeting methods (Occurrences, Timeslices, and RatioTimeslices):
- **Availability SLO** - Internet provider SLA with 99% uptime
- **Occurrences SLO** - Search latency treating all searches equally
- **Ratio Timeslices SLO** - Main page availability based on response codes

#### [Treat Low Traffic as Equally Important](examples/treat-low-traffic-as-equally-important/)
Maintains high reliability regardless of traffic volume, ensuring low-usage services receive equal attention.

### Example Features

Each example includes:
- ✅ Complete, working Go code
- ✅ Detailed descriptions of what is being measured
- ✅ OpenSLO-compliant metadata with ontology labels
- ✅ Prometheus/BigQuery query examples
- ✅ Automated validation tests
- 📚 README with methodology explanations and best practices

## Project Structure

```
slogo/
├── ontology/          # SLO labeling and categorization system
├── examples/
│   ├── red-metrics/       # Rate, Errors, Duration (service monitoring)
│   ├── use-metrics/       # Utilization, Saturation, Errors (infrastructure)
│   ├── ai-agents/         # AI agent platform business metrics
│   ├── saas-crm/          # SaaS CRM user journey metrics
│   ├── budgeting-method/  # Budgeting method comparisons
│   └── treat-low-traffic-as-equally-important/
├── datadog/          # Datadog integration utilities
└── cmd/parse/        # CLI tools for parsing SLO files
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Resources

### OpenSLO
- [OpenSLO Specification](https://github.com/OpenSLO/OpenSLO)
- [OpenSLO Go SDK](https://github.com/OpenSLO/go-sdk)

### SRE & Monitoring Methodologies
- [Site Reliability Engineering (SRE) Book](https://sre.google/books/)
- [The RED Method](https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/) - Tom Wilkie, Grafana Labs
- [The USE Method](https://www.brendangregg.com/usemethod.html) - Brendan Gregg
- [Google's Four Golden Signals](https://sre.google/sre-book/monitoring-distributed-systems/)

### Product & Business Metrics
- [Measuring Product Health](https://www.reforge.com/blog/retention-engagement-growth-silent-killer) - Reforge
- [SaaS Metrics 2.0](https://www.forentrepreneurs.com/saas-metrics-2/) - For Entrepreneurs
- [The Lean Analytics Book](https://leananalyticsbook.com/) - Alistair Croll & Benjamin Yoskovitz

 [build-status-svg]: https://github.com/grokify/slogo/actions/workflows/ci.yaml/badge.svg?branch=main
 [build-status-url]: https://github.com/grokify/slogo/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/grokify/slogo/actions/workflows/lint.yaml/badge.svg?branch=main
 [lint-status-url]: https://github.com/grokify/slogo/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/slogo
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/slogo
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/slogo
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/slogo
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/slogo/blob/master/LICENSE

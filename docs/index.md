# SLOGo

Go utilities for working with [OpenSLO](https://github.com/OpenSLO/OpenSLO) (Service Level Objectives) specifications.

[![Go Report Card](https://goreportcard.com/badge/github.com/grokify/slogo)](https://goreportcard.com/report/github.com/grokify/slogo)
[![GoDoc](https://pkg.go.dev/badge/github.com/grokify/slogo)](https://pkg.go.dev/github.com/grokify/slogo)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/grokify/slogo/blob/master/LICENSE)

## Overview

`slogo` is a Go library that provides convenient utilities for reading, writing, and validating OpenSLO specification files. It builds on top of the [OpenSLO Go SDK](https://github.com/OpenSLO/go-sdk) to provide a simpler interface for working with SLO definitions in JSON and YAML formats.

## Features

- 📖 Read OpenSLO objects from files or readers (auto-detects JSON/YAML format)
- 📝 Write OpenSLO objects to files in JSON or YAML format
- ✅ Validate SLO definitions against the OpenSLO specification
- 🔒 Type-safe Go representations of SLO objects
- 🏷️ Predefined severity and attribute constants
- 📊 Comprehensive SLO ontology for metric labeling and categorization
- 🚀 Production-ready examples for RED, USE, AI Agents, and SaaS metrics

## Quick Start

```bash
go get github.com/grokify/slogo
```

```go
import "github.com/grokify/slogo"

// Read from file (auto-detects JSON or YAML)
objs, err := slogo.ReadFile("slo-definition.yaml")
if err != nil {
    log.Fatal(err)
}

// Validate all objects
if err := objs.Validate(); err != nil {
    log.Fatal(err)
}
```

## What's Included

### Ontology System

A comprehensive [labeling system](ontology/overview.md) for organizing SLOs across multiple dimensions:

- 🔧 **Frameworks**: RED, USE, Custom
- 📦 **Layers**: Service, Infrastructure, Business, Application
- 🎯 **Scopes**: Customer-facing, Internal, Business-outcome
- 👥 **Audiences**: SRE, Engineering, Product, Executive
- 📊 **Categories**: Availability, Latency, Throughput, Quality
- 📈 **Metric Types**: Rate, Errors, Duration, Retention, Stickiness

### Production-Ready Examples

64 SLOs across 6 example sets:

| Example Set | SLOs | Description |
|-------------|------|-------------|
| [RED Metrics](examples/red-metrics.md) | 5 | Rate, Errors, Duration for services |
| [USE Metrics](examples/use-metrics.md) | 11 | Utilization, Saturation, Errors for infrastructure |
| [AI Agents](examples/ai-agents.md) | 20 | Quality, performance, cost efficiency |
| [SaaS CRM](examples/saas-crm.md) | 25 | User journey from activation to retention |
| [Budgeting Methods](examples/budgeting-methods.md) | 3 | Occurrences vs Timeslices comparison |

### Domain Ontologies

Opt-in [domain packages](ontology/domains.md) for specialized use cases:

- **IAM** - Identity & Access Management
- **Product** - AARRR journey stages
- **Business** - Industry verticals
- **Security** - Risk and threat classification
- **Compliance** - Regulatory frameworks
- **Operations** - Incident response, SRE

## Resources

### OpenSLO

- 📜 [OpenSLO Specification](https://github.com/OpenSLO/OpenSLO)
- 📦 [OpenSLO Go SDK](https://github.com/OpenSLO/go-sdk)

### SRE & Monitoring Methodologies

- 📖 [Site Reliability Engineering (SRE) Book](https://sre.google/books/)
- 🔴 [The RED Method](https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/) - Tom Wilkie
- ⚡ [The USE Method](https://www.brendangregg.com/usemethod.html) - Brendan Gregg
- 🔔 [Google's Four Golden Signals](https://sre.google/sre-book/monitoring-distributed-systems/)

### Product & Business Metrics

- 📈 [Measuring Product Health](https://www.reforge.com/blog/retention-engagement-growth-silent-killer) - Reforge
- 💰 [SaaS Metrics 2.0](https://www.forentrepreneurs.com/saas-metrics-2/) - For Entrepreneurs
- 📊 [The Lean Analytics Book](https://leananalyticsbook.com/)

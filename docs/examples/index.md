# Examples Overview

The `examples/` directory contains production-ready SLO definitions organized by monitoring framework and use case.

## Summary

| Example Set | SLOs | Framework | Layer | Description |
|-------------|------|-----------|-------|-------------|
| [RED Metrics](red-metrics.md) | 5 | RED | Service | Request-driven service monitoring |
| [USE Metrics](use-metrics.md) | 11 | USE | Infrastructure | Resource monitoring |
| [AI Agents](ai-agents.md) | 20 | Custom | Business | AI platform metrics |
| [SaaS CRM](saas-crm.md) | 25 | Custom | Business | User journey metrics |
| [Budgeting Methods](budgeting-methods.md) | 3 | Custom | Service | Budget method comparison |

**Total: 64 SLOs**

## Example Features

Each example includes:

- ✅ Complete, working Go code
- 📋 Detailed descriptions of what is being measured
- 🏷️ OpenSLO-compliant metadata with ontology labels
- 🔍 Prometheus/BigQuery query examples
- 🧪 Automated validation tests
- 📚 README with methodology explanations

## By Monitoring Framework

### RED (Rate, Errors, Duration)

Request-driven monitoring for services and APIs:

- [RED Metrics](red-metrics.md) - Core RED implementation
- [AI Agents](ai-agents.md) - AI response time, quality, errors

### USE (Utilization, Saturation, Errors)

Infrastructure and resource monitoring:

- [USE Metrics](use-metrics.md) - CPU, memory, disk, network

### Custom / Business

Business metrics and custom implementations:

- [SaaS CRM](saas-crm.md) - User journey and engagement
- [AI Agents](ai-agents.md) - Cost efficiency, task completion
- [Budgeting Methods](budgeting-methods.md) - SLO budgeting strategies

## By Audience

### SRE / Platform

- [RED Metrics](red-metrics.md) - Service reliability
- [USE Metrics](use-metrics.md) - Infrastructure health

### Product

- [SaaS CRM](saas-crm.md) - User engagement, retention
- [AI Agents](ai-agents.md) - User satisfaction, adoption

### Executive

- [SaaS CRM](saas-crm.md) - Business outcomes, revenue metrics
- [AI Agents](ai-agents.md) - Cost efficiency, ROI

## By Use Case

### Service Monitoring

Monitor API and microservice health:

```go
// RED metrics for API
slo := redmetrics.ExampleAvailabilitySLO()
```

### Infrastructure Monitoring

Monitor underlying resources:

```go
// USE metrics for CPU
slo := usemetrics.ExampleCPUUtilizationSLO()
```

### User Engagement

Track user activity and stickiness:

```go
// DAU/MAU ratio (stickiness)
slo := saascrm.ExampleDAUMAURatioSLO()
```

### AI/ML Platforms

Monitor AI agent performance:

```go
// AI response quality
slo := aiagents.ExampleAgentResponseQualitySLO()
```

## Label Distribution

See the [Metrics Report](metrics-report.md) for detailed analysis of label distribution across all 64 SLOs.

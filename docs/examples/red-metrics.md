# RED Metrics

RED (Rate, Errors, Duration) is a service-level monitoring methodology focused on request-driven services.

## What are RED Metrics?

RED is a monitoring framework introduced by Tom Wilkie at Grafana Labs, designed specifically for monitoring request-driven services (APIs, microservices, web applications). The three pillars are:

- 📈 **Rate**: The number of requests per second your service is handling
- ❌ **Errors**: The number/percentage of those requests that are failing
- ⏱️ **Duration**: The amount of time those requests take (latency percentiles)

## SLOs in this Example

| SLO | Type | Description |
|-----|------|-------------|
| `ExampleRateSLO` | Rate | Tracks requests per second |
| `ExampleErrorRateSLO` | Errors | Success ratio (2xx, 3xx vs 4xx, 5xx) |
| `ExampleDurationSLO` | Duration | P95 latency |
| `ExampleDurationP99SLO` | Duration | P99 latency (tail) |
| `ExampleAvailabilitySLO` | Derived | Availability from Rate + Errors |

## Usage

```go
import redmetrics "github.com/grokify/slogo/examples/red-metrics"

// Get individual SLOs
rateSLO := redmetrics.ExampleRateSLO()
errorSLO := redmetrics.ExampleErrorRateSLO()
durationSLO := redmetrics.ExampleDurationSLO()

// Get all SLOs
slos := redmetrics.SLOs()
```

## When to Use RED Metrics

RED metrics are ideal for:

- **Request-driven services**: REST APIs, GraphQL APIs, microservices
- **Web applications**: Backend services handling HTTP requests
- **Service-to-service communication**: Internal API calls
- **User-facing endpoints**: Any service where users make requests

## Prometheus Queries

### Rate

```promql
rate(http_requests_total[5m])
```

### Errors

```promql
sum(rate(http_requests_total{status=~"2..|3.."}[5m])) / sum(rate(http_requests_total[5m]))
```

### Duration (P95)

```promql
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))
```

### Duration (P99)

```promql
histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m]))
```

## Ontology Labels

All RED metric SLOs use these labels:

```go
ontology.LabelFramework:  ontology.FrameworkRED,
ontology.LabelLayer:      ontology.LayerService,
ontology.LabelScope:      ontology.ScopeCustomerFacing,
ontology.LabelAudience:   ontology.AudienceSRE,
```

## Complementary Frameworks

| Framework | Focus | Use Case |
|-----------|-------|----------|
| [USE Metrics](use-metrics.md) | Infrastructure | CPU, memory, disk, network |
| Four Golden Signals | Hybrid | RED + Saturation |
| DORA Metrics | Deployment | DevOps performance |

## References

- 🔴 [The RED Method](https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/) - Tom Wilkie, Grafana Labs
- 📖 [Monitoring Distributed Systems](https://sre.google/sre-book/monitoring-distributed-systems/) - Google SRE Book
- 📜 [OpenSLO Specification](https://github.com/OpenSLO/OpenSLO)

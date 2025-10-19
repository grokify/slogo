# RED Metrics Examples

This directory contains OpenSLO examples for RED (Rate, Errors, Duration) metrics, a service-level monitoring methodology focused on request-driven services.

## What are RED Metrics?

RED is a monitoring framework introduced by Tom Wilkie at Grafana Labs, designed specifically for monitoring request-driven services (APIs, microservices, web applications). The three pillars are:

- **Rate**: The number of requests per second your service is handling
- **Errors**: The number/percentage of those requests that are failing
- **Duration**: The amount of time those requests take (typically measured as latency percentiles)

## Files in this Directory

### rate.go
Measures the request rate (throughput) for an API service:
- `ExampleRateSLO` - Tracks requests per second to ensure the service handles expected traffic volume

### error.go
Monitors the error rate of requests:
- `ExampleErrorRateSLO` - Tracks the ratio of successful responses (2xx, 3xx) vs errors (4xx, 5xx)

### duration.go
Tracks request latency and response times:
- `ExampleDurationSLO` - Measures P95 latency to ensure fast response times
- `ExampleDurationP99SLO` - Measures P99 latency for tail latency monitoring

### red_test.go
Validation tests for all RED metric SLOs

## When to Use RED Metrics

RED metrics are ideal for:
- **Request-driven services**: REST APIs, GraphQL APIs, microservices
- **Web applications**: Backend services handling HTTP requests
- **Service-to-service communication**: Internal API calls
- **User-facing endpoints**: Any service where users make requests

## Example Prometheus Queries

The examples use Prometheus as the metric source with queries like:

**Rate:**
```promql
rate(http_requests_total[5m])
```

**Errors:**
```promql
sum(rate(http_requests_total{status=~"2..|3.."}[5m])) / sum(rate(http_requests_total[5m]))
```

**Duration:**
```promql
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))
```

## Complementary Frameworks

- **USE Metrics**: For infrastructure/resource monitoring (see `../use-metrics/`)
- **Four Golden Signals**: Google's framework (similar to RED + Saturation)
- **DORA Metrics**: For measuring DevOps/deployment performance

## References

- [The RED Method](https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/) - Tom Wilkie
- [Monitoring Distributed Systems](https://sre.google/sre-book/monitoring-distributed-systems/) - Google SRE Book
- OpenSLO Specification: https://github.com/OpenSLO/OpenSLO

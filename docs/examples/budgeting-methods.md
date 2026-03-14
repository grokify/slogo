# Budgeting Methods

Compares different SLO budgeting methods: Occurrences, Timeslices, and RatioTimeslices.

## What are Budgeting Methods?

OpenSLO supports different methods for calculating error budgets and SLO compliance:

| Method | Description | Use Case |
|--------|-------------|----------|
| **Occurrences** | Count each event equally | High-volume, uniform requests |
| **Timeslices** | Divide time into windows | Time-based availability |
| **RatioTimeslices** | Ratio within time windows | Combined approach |

## SLOs in this Example (3 total)

| SLO | Method | Description |
|-----|--------|-------------|
| `ExampleAvailabilitySLO` | Timeslices | Internet provider 99% uptime |
| `ExampleOccurrencesSLO` | Occurrences | Search latency (all searches equal) |
| `ExampleRatioTimeslicesSLO` | RatioTimeslices | Main page availability by response codes |

## Usage

```go
import budgeting "github.com/grokify/slogo/examples/budgeting-method"

// Get individual SLOs
availSLO := budgeting.ExampleAvailabilitySLO()
searchSLO := budgeting.ExampleOccurrencesSLO()

// Get all SLOs
slos := budgeting.SLOs()
```

## Occurrences Method

Treats every event (request) equally. Good for:

- High-volume APIs where every request matters equally
- Services where traffic volume is consistent
- Situations where you want to protect all users equally

```yaml
budgetingMethod: Occurrences
objectives:
  - displayName: "99.9% of requests under 200ms"
    target: 0.999
    operator: lt
    value: 0.2
```

**Example**: Search service where every search should be fast

```promql
sum(rate(search_latency_seconds_bucket{le="0.2"}[5m])) /
sum(rate(search_latency_seconds_count[5m]))
```

## Timeslices Method

Divides time into fixed windows and evaluates each window. Good for:

- Availability SLOs
- Time-based SLAs (e.g., "99% uptime")
- Services with variable traffic

```yaml
budgetingMethod: Timeslices
objectives:
  - displayName: "99% uptime"
    target: 0.99
    timeSliceTarget: 1.0
    timeSliceWindow: "1m"
```

**Example**: Internet provider availability

## RatioTimeslices Method

Combines ratio calculation within time windows. Good for:

- Complex availability calculations
- Response code-based availability
- Situations needing both time and ratio perspective

```yaml
budgetingMethod: RatioTimeslices
objectives:
  - displayName: "Main page availability"
    target: 0.98
    timeSliceTarget: 0.99
    timeSliceWindow: "5m"
```

**Example**: Web page availability based on response codes

## Comparison

### Scenario: 99% Target, 100K daily requests

| Method | Failure Mode | Budget Interpretation |
|--------|--------------|----------------------|
| **Occurrences** | 1000 failed requests | Any 1000 failures OK |
| **Timeslices** | 14.4 minutes downtime | Time-based failures |
| **RatioTimeslices** | Varies by window | Per-window ratio |

### When to Use Each

| Situation | Recommended Method |
|-----------|-------------------|
| High-volume, uniform traffic | Occurrences |
| SLA with uptime percentage | Timeslices |
| Variable traffic patterns | Timeslices |
| Response code analysis | RatioTimeslices |
| Need per-window compliance | RatioTimeslices |

## Low Traffic Considerations

!!! warning "Low Traffic Trap"
    With Occurrences, a single failure in low-traffic periods can consume significant budget. Consider Timeslices for services with variable traffic.

See [Treat Low Traffic as Equally Important](https://github.com/OpenSLO/OpenSLO/tree/main/examples/treat-low-traffic-as-equally-important) for strategies.

## Ontology Labels

Budgeting method examples use service-layer labels:

```go
ontology.LabelFramework:  ontology.FrameworkCustom,
ontology.LabelLayer:      ontology.LayerService,
ontology.LabelScope:      ontology.ScopeCustomerFacing,
ontology.LabelAudience:   ontology.AudienceSRE,
```

## References

- 📜 [OpenSLO Specification - Budgeting Methods](https://github.com/OpenSLO/OpenSLO#budgetingmethod)
- 📖 [Google SRE Book - Error Budgets](https://sre.google/sre-book/embracing-risk/)
- 📊 [SLO Math Deep Dive](https://www.nobl9.com/resources/slo-math-deep-dive)

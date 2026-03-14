# USE Metrics

USE (Utilization, Saturation, Errors) is Brendan Gregg's methodology for infrastructure and resource monitoring.

## What are USE Metrics?

USE provides a systematic approach to analyzing system performance by examining every resource:

- 📊 **Utilization**: Percentage of resource being used (0-100%)
- 🌊 **Saturation**: Degree of queuing or backlog
- ⚠️ **Errors**: Error events on the resource

## SLOs in this Example (11 total)

### Utilization SLOs

| SLO | Resource | Description |
|-----|----------|-------------|
| `ExampleCPUUtilizationSLO` | CPU | Average CPU usage |
| `ExampleMemoryUtilizationSLO` | Memory | Memory usage percentage |
| `ExampleDiskUtilizationSLO` | Disk | Disk space usage |

### Saturation SLOs

| SLO | Resource | Description |
|-----|----------|-------------|
| `ExampleCPULoadAverageSLO` | CPU | Load average / CPU count |
| `ExampleSwapUsageSLO` | Memory | Swap space usage |
| `ExampleDiskIOSaturationSLO` | Disk | I/O wait percentage |
| `ExampleNetworkBandwidthSLO` | Network | Bandwidth utilization |

### Error SLOs

| SLO | Resource | Description |
|-----|----------|-------------|
| `ExampleDiskIOErrorsSLO` | Disk | I/O error rate |
| `ExampleNetworkErrorsSLO` | Network | Packet error rate |
| `ExampleMemoryECCErrorsSLO` | Memory | ECC memory errors |
| `ExampleCPUThrottlingSLO` | CPU | Throttling events |

## Usage

```go
import usemetrics "github.com/grokify/slogo/examples/use-metrics"

// Get individual SLOs
cpuSLO := usemetrics.ExampleCPUUtilizationSLO()
memorySLO := usemetrics.ExampleMemoryUtilizationSLO()

// Get all SLOs
slos := usemetrics.SLOs()
```

## When to Use USE Metrics

USE metrics are ideal for:

- **Physical servers**: CPU, memory, disk monitoring
- **Virtual machines**: VM resource tracking
- **Containers**: Pod/container resource limits
- **Network infrastructure**: Bandwidth, packet handling
- **Storage systems**: IOPS, throughput, capacity

## Prometheus Queries

### CPU Utilization

```promql
100 - (avg(rate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
```

### Memory Utilization

```promql
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100
```

### Disk I/O Saturation

```promql
avg(rate(node_disk_io_time_weighted_seconds_total[5m]))
```

### Network Bandwidth

```promql
rate(node_network_receive_bytes_total[5m]) + rate(node_network_transmit_bytes_total[5m])
```

## Ontology Labels

All USE metric SLOs use these labels:

```go
ontology.LabelFramework:    ontology.FrameworkUSE,
ontology.LabelLayer:        ontology.LayerInfrastructure,
ontology.LabelScope:        ontology.ScopeInternal,
ontology.LabelAudience:     ontology.AudienceSRE,
ontology.LabelResourceType: ontology.ResourceTypeCPU, // or Memory, Disk, Network
```

## Resource Matrix

For each resource, ask: **Utilization? Saturation? Errors?**

| Resource | Utilization | Saturation | Errors |
|----------|-------------|------------|--------|
| CPU | Usage % | Load average, run queue | Throttling |
| Memory | Used % | Swap usage, OOM events | ECC errors |
| Disk | Space %, IOPS | I/O wait | I/O errors |
| Network | Bandwidth % | Queue depth | Packet errors |

## Complementary Frameworks

| Framework | Focus | Use Case |
|-----------|-------|----------|
| [RED Metrics](red-metrics.md) | Services | Request handling |
| Four Golden Signals | Hybrid | Services + Saturation |

## References

- ⚡ [The USE Method](https://www.brendangregg.com/usemethod.html) - Brendan Gregg
- 📖 [Systems Performance](https://www.brendangregg.com/systems-performance-2nd-edition-book.html) - Brendan Gregg
- 📜 [OpenSLO Specification](https://github.com/OpenSLO/OpenSLO)

# USE Metrics Examples

This directory contains OpenSLO examples for USE (Utilization, Saturation, Errors) metrics, a resource-focused monitoring methodology for infrastructure and hardware.

## What are USE Metrics?

USE is a monitoring methodology created by Brendan Gregg for analyzing system performance. It focuses on every resource in your system and checks three key metrics:

- **Utilization**: The average time that the resource was busy servicing work (% busy)
- **Saturation**: The degree to which the resource has extra work which it can't service (queue length)
- **Errors**: The count of error events for that resource

## Files in this Directory

### utilization.go
Tracks resource utilization (how busy resources are):
- `ExampleCPUUtilizationSLO` - Monitor CPU usage percentage
- `ExampleMemoryUtilizationSLO` - Monitor memory usage percentage
- `ExampleDiskUtilizationSLO` - Monitor disk space usage percentage

### saturation.go
Measures resource saturation (queueing and contention):
- `ExampleCPUSaturationSLO` - Track CPU load average per core
- `ExampleMemorySaturationSLO` - Monitor swap usage as memory pressure indicator
- `ExampleDiskSaturationSLO` - Track disk I/O utilization and queuing
- `ExampleNetworkSaturationSLO` - Monitor network bandwidth saturation

### errors.go
Monitors resource-level errors:
- `ExampleDiskErrorsSLO` - Track disk I/O errors
- `ExampleNetworkErrorsSLO` - Monitor network packet errors and drops
- `ExampleMemoryErrorsSLO` - Track ECC memory errors
- `ExampleCPUThrottlingErrorsSLO` - Monitor CPU throttling in containers

### use_test.go
Validation tests for all USE metric SLOs

## When to Use USE Metrics

USE metrics are ideal for:
- **Infrastructure monitoring**: Servers, VMs, containers, bare metal
- **Resource capacity planning**: Understanding resource constraints
- **Performance debugging**: Identifying bottlenecks at the hardware level
- **System health**: Monitoring physical and virtual resource health

## How to Apply USE Methodology

For each resource in your system, ask:

1. **What is the utilization?** (How busy is it?)
2. **What is the saturation?** (Is work queueing?)
3. **Are there any errors?** (Hardware/resource failures?)

Common resources to monitor:
- CPU (cores)
- Memory (RAM)
- Disk (storage devices)
- Network (interfaces)
- Storage I/O
- GPU (if applicable)

## Example Prometheus Queries

The examples use Prometheus with node_exporter metrics:

**CPU Utilization:**
```promql
avg(rate(node_cpu_seconds_total{mode!="idle"}[5m])) * 100
```

**CPU Saturation:**
```promql
node_load5 / count(node_cpu_seconds_total{mode="idle"})
```

**Disk Errors:**
```promql
rate(node_disk_io_errors_total[5m])
```

## Complementary Frameworks

- **RED Metrics**: For request-driven service monitoring (see `../red-metrics/`)
- **Four Golden Signals**: Google's framework (combines aspects of both)
- **TSI Method**: Thread, Saturation, Instrumentation (for application-level)

## References

- [The USE Method](https://www.brendangregg.com/usemethod.html) - Brendan Gregg
- [Systems Performance](https://www.brendangregg.com/systems-performance-2nd-edition-book.html) - Book by Brendan Gregg
- [Linux Performance Analysis in 60 seconds](https://netflixtechblog.com/linux-performance-analysis-in-60-000-milliseconds-accc10403c55) - Netflix
- OpenSLO Specification: https://github.com/OpenSLO/OpenSLO

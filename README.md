# slogo

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

## Examples

The [`examples`](examples/) directory contains real-world examples of SLO definitions with detailed explanations:

### [Budgeting Methods](examples/budgeting-method/)

Compares different budgeting methods (Occurrences, Timeslices, and RatioTimeslices) with practical examples:

- **Availability SLO** (`timeslices-slo.go`) - Reflects an SLA for internet provider availability (99% uptime)
- **Occurrences SLO** (`occurences-slo.go`) - Measures search latency treating all searches equally
- **Ratio Timeslices SLO** (`ratio-timeslices.go`) - Tracks main page availability based on response codes

### [Treat Low Traffic as Equally Important](examples/treat-low-traffic-as-equally-important/)

Demonstrates how to maintain high reliability regardless of traffic volume, ensuring that services with low usage receive the same attention as high-traffic services.

Each example includes:
- Complete, working Go code
- Detailed descriptions of what is being measured
- Explanations of why specific budgeting methods are chosen
- Automated tests to validate the SLO definitions

## Dependencies

- [github.com/OpenSLO/go-sdk](https://github.com/OpenSLO/go-sdk) - OpenSLO Go SDK
- [github.com/grokify/mogo](https://github.com/grokify/mogo) - Go utilities

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Resources

- [OpenSLO Specification](https://github.com/OpenSLO/OpenSLO)
- [OpenSLO Go SDK](https://github.com/OpenSLO/go-sdk)
- [Site Reliability Engineering (SRE) Book](https://sre.google/books/)

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

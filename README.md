```sh
                           ____        _     _ _      ___ ____  
                          |  _ \ _   _| |__ | (_) ___|_ _|  _ \ 
                          | |_) | | | | '_ \| | |/ __|| || | | |
                          |  __/| |_| | |_) | | | (__ | || |_| |
                          |_|    \__,_|_.__/|_|_|\___|___|____/ 
```
<!-- [![Sourcegraph](https://sourcegraph.com/github.com/agentstation/publicid/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/agentstation/publicid?badge) -->
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/agentstation/publicid)
[![Go Report Card](https://goreportcard.com/badge/github.com/agentstation/publicid?style=flat-square)](https://goreportcard.com/report/github.com/agentstation/publicid)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/agentstation/publicid/ci.yaml?style=flat-square)](https://github.com/agentstation/publicid/actions)
[![codecov](https://codecov.io/gh/agentstation/publicid/branch/master/graph/badge.svg?token=35UM5QX1Q3)](https://codecov.io/gh/agentstation/publicid)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/agentstation/publicid/master/LICENSE)
<!-- [![Forum](https://img.shields.io/badge/community-forum-00afd1.svg?style=flat-square)](https://github.com/agentstation/publicid/discussions) -->
<!-- [![Twitter](https://img.shields.io/badge/twitter-@agentstationHQ-55acee.svg?style=flat-square)](https://twitter.com/agentstationHQ) -->

The `publicid` package generates and validates NanoID strings designed to be publicly exposed.

## Installation

```sh
go get github.com/agentstation/publicid
```

## Usage

To use the `publicid` package in your Go code, follow these steps:

1. Import the package:

```go
import "github.com/agentstation/publicid"
```

2. Generate a short public ID (8 characters):

```go
id, err := publicid.New()
if err != nil {
    log.Fatalf("Failed to generate public ID: %v", err)
}
fmt.Println("Generated short public ID:", id)
// Output: Generated short public ID: Ab3xY9pQ
```

3. Generate a long public ID (12 characters):

```go
longID, err := publicid.NewLong()
if err != nil {
    log.Fatalf("Failed to generate long public ID: %v", err)
}
fmt.Println("Generated long public ID:", longID)
// Output: Generated long public ID: 7Zt3xY9pQr5W
```

4. Use the `Attempts` option to specify the number of generation attempts:

```go
id, err := publicid.New(publicid.Attempts(5))
if err != nil {
    log.Fatalf("Failed to generate public ID: %v", err)
}
fmt.Println("Generated public ID with 5 attempts:", id)
// Output: Generated public ID with 5 attempts: Kj2mN8qL
```

5. Validate a short public ID:

```go
shortID := "Ab3xY9pQ"
err := publicid.Validate(shortID)
if err != nil {
    fmt.Println("Invalid short ID:", err)
} else {
    fmt.Println("Valid short ID:", shortID)
}
// Output: Valid short ID: Ab3xY9pQ
```

6. Validate a long public ID:

```go
longID := "7Zt3xY9pQr5W"
err := publicid.ValidateLong("exampleField", longID)
if err != nil {
    fmt.Println("Invalid long ID:", err)
} else {
    fmt.Println("Valid long ID:", longID)
}
// Output: Valid long ID: 7Zt3xY9pQr5W
```

## Makefile

```sh
make help

Usage:
  make <target>

General
  help                  Display the list of targets and their descriptions

Tooling
  install-devbox        Install Devbox
  devbox-update         Update Devbox
  devbox                Run Devbox shell

Installation
  install               Download go modules

Development
  fmt                   Run go fmt
  generate              Generate and embed go documentation into README.md
  vet                   Run go vet
  lint                  Run golangci-lint

Testing & Benchmarking
  test                  Run Go tests
  bench                 Run Go benchmarks
  ```

## Benchmarks

> **Note:** These benchmarks were run on an Apple M2 Max CPU with 12 cores (8 performance and 4 efficiency) and 32 GB of memory, running macOS 14.6.1.

*Your mileage may vary.*

```sh
go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/agentstation/publicid
BenchmarkNew-12                      2012978           574.8 ns/op
BenchmarkNewWithAttempts-12          2091734           577.3 ns/op
BenchmarkLong-12                     1966120           616.9 ns/op
BenchmarkLongWithAttempts-12         1952052           610.4 ns/op
BenchmarkValidate-12                100000000            10.73 ns/op
BenchmarkValidateLong-12            99347000            13.31 ns/op
PASS
ok      github.com/agentstation/publicid    9.790s
```

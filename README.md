# go-otelx

[![tag](https://img.shields.io/github/tag/cyg-pd/go-otelx.svg)](https://github.com/cyg-pd/go-otelx/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.24-%23007d9c)
[![GoDoc](https://godoc.org/github.com/cyg-pd/go-otelx?status.svg)](https://pkg.go.dev/github.com/cyg-pd/go-otelx)
![Build Status](https://github.com/cyg-pd/go-otelx/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/cyg-pd/go-otelx)](https://goreportcard.com/report/github.com/cyg-pd/go-otelx)
[![Coverage](https://img.shields.io/codecov/c/github/cyg-pd/go-otelx)](https://codecov.io/gh/cyg-pd/go-otelx)
[![Contributors](https://img.shields.io/github/contributors/cyg-pd/go-otelx)](https://github.com/cyg-pd/go-otelx/graphs/contributors)
[![License](https://img.shields.io/github/license/cyg-pd/go-otelx)](./LICENSE)

## 🚀 Install

```sh
go get github.com/cyg-pd/go-otelx@v1
```

This library is v1 and follows SemVer strictly.

No breaking changes will be made to exported APIs before v2.0.0.

## 💡 Usage

You can import `otelx` using:

```go
package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/cyg-pd/go-otelx"
	_ "github.com/cyg-pd/go-otelx/autoconf" // auto setup opentelemetry sdk with environment variable
)

var tracer = otelx.Tracer()

// $ export OTEL_METRICS_EXPORTER=console
// $ export OTEL_TRACES_EXPORTER=console
// $ export OTEL_LOGS_EXPORTER=console

func main() {
	doWork(context.Background())

	// wait console output
	<-time.After(time.Second * 10)
	// {"Name":"main","SpanContext":{"TraceID":"8298c81701e88ce4195f36d221a7561a","SpanID":"7933655ddc3d221d","TraceFlags":"01","TraceState":"","Remote...
}

func doWork(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "main")
	defer span.End()

	// do your work with ctx
	slog.InfoContext(ctx, span.SpanContext().TraceID().String()) // 2025/05/23 00:00:00 INFO 63b6c64c62d8fcb943e1e34f8f5ac757
}
```

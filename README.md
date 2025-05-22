# go-otelx

[![tag](https://img.shields.io/github/tag/cyg-pd/go-otelx.svg)](https://github.com/cyg-pd/go-otelx/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
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

This library has no dependencies outside the Go standard library.

## 💡 Usage

You can import `go-otelx` using:

```go
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/cyg-pd/go-otelx"
	slogotel "github.com/remychantenay/slog-otel"
)

var tracer = otelx.Tracer()

func init() {
	slog.SetDefault(slog.New(slogotel.OtelHandler{
		Next: slog.NewTextHandler(os.Stdout, nil),
	}))
}

func main() {
	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	slog.InfoContext(ctx, "hi") // time=2025-05-22T00:00:00.000Z level=INFO msg=hi trace_id=47d484a406418847427ac3a0a1f3f145 span_id=44097bd488b383a5
}
```

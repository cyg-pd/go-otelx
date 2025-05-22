package otelx

import (
	"github.com/cyg-pd/go-reflectx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func Tracer() trace.Tracer {
	return otel.Tracer(reflectx.PackageName(2))
}

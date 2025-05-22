package otelx

import (
	"github.com/cyg-pd/go-reflectx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

func Meter() metric.Meter {
	return otel.Meter(reflectx.PackageName(2))
}

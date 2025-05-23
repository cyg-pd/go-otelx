package autoconf

import (
	"context"

	"go.opentelemetry.io/contrib/exporters/autoexport"
	"go.opentelemetry.io/contrib/propagators/autoprop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

func init() {
	ctx := context.Background()

	otel.SetTextMapPropagator(autoprop.NewTextMapPropagator())

	if r, err := autoexport.NewMetricReader(ctx); err == nil {
		otel.SetMeterProvider(metric.NewMeterProvider(metric.WithReader(r)))
	}

	if e, err := autoexport.NewSpanExporter(ctx); err == nil {
		otel.SetTracerProvider(trace.NewTracerProvider(trace.WithBatcher(e)))
	}

	if e, err := autoexport.NewLogExporter(ctx); err == nil {
		global.SetLoggerProvider(log.NewLoggerProvider(log.WithProcessor(log.NewBatchProcessor(e))))
	}
}

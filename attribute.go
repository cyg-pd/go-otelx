package otelx

import (
	"go.opentelemetry.io/otel/attribute"
)

func Uint(key string, value uint) attribute.KeyValue {
	return attribute.Int(key, int(value))
}

func Uint64(key string, value uint64) attribute.KeyValue {
	return attribute.Int64(key, int64(value))
}

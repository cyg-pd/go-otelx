package otelx

import "go.opentelemetry.io/otel/sdk/resource"

var svcName string

// ServiceName read service.name from otel resource
func ServiceName() string {
	if svcName == "" {
		svcName = serviceName()
	}
	return svcName
}

func serviceName() string {
	itr := resource.Default().Iter()

	for itr.Next() {
		kv := itr.Attribute()
		if kv.Key == "service.name" {
			return kv.Value.AsString()
		}
	}

	return ""
}

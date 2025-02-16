package metrics

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type Metrics interface {
	GetExporter() http.Handler
	GetMetricsProvider() metric.MeterProvider
	RegisterCounter(name, description string) error
	AddCount(ctx context.Context, name string, value int64, labels map[string]attribute.Value) error
	RegisterUpDownSumObserver(name, description string, callbackFunc metric.Int64Callback) error
	RegisterValueObserver(name, description string, callbackFunc metric.Int64Callback) error
}

var M Metrics

func GetExporter() http.Handler {
	if M == nil {
		return nil
	}
	return M.GetExporter()
}

func GetMetricsProvider() metric.MeterProvider {
	if M == nil {
		return nil
	}
	return M.GetMetricsProvider()
}

func RegisterCounter(name, description string) error {
	if M == nil {
		return nil
	}
	return M.RegisterCounter(name, description)
}

func AddCount(ctx context.Context, name string, value int64, labels map[string]attribute.Value) error {
	if M == nil {
		return nil
	}
	return M.AddCount(ctx, name, value, labels)
}

func RegisterUpDownSumObserver(name, description string, callbackFunc metric.Int64Callback) error {
	if M == nil {
		return nil
	}
	return M.RegisterUpDownSumObserver(name, description, callbackFunc)
}

func RegisterValueObserver(name, description string, callbackFunc metric.Int64Callback) error {
	if M == nil {
		return nil
	}
	return M.RegisterValueObserver(name, description, callbackFunc)
}

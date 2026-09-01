package tracing

import (
	"context"

	"github.com/tsingsun/woocoo/contrib/telemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var (
	noopSpan = noop.Span{}
)

// conditionalTracer wraps an OTel tracer and skips span creation when tracing
// is disabled, avoiding allocations on the hot path.
type conditionalTracer struct {
	noop.Tracer
	name string
}

// NewTracer returns a tracer that is conditional on the global tracingEnabled flag.
func NewTracer(name string) trace.Tracer {
	return &conditionalTracer{name: name}
}

func (t *conditionalTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if telemetry.GlobalConfig() == nil {
		return ctx, noopSpan
	}
	return otel.Tracer(t.name).Start(ctx, spanName, opts...)
}

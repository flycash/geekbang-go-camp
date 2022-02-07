package rpc

import (
	"context"
	"fmt"
	"geekbang/geekbang-go-camp/ninth/dto"
	"github.com/opentracing/opentracing-go"
)

type ServerSideTracingFilterBuilder struct {
}

func (t *ServerSideTracingFilterBuilder) Build(next Filter) Filter {
	return func(ctx context.Context, req *dto.Request) (*dto.Response, error) {
		fmt.Printf("this is server tracing filter\n")
		tracer := opentracing.GlobalTracer()
		spanCtx, err := tracer.Extract(opentracing.TextMap,
			opentracing.TextMapCarrier(req.Meta))
		if err == nil {
			operationName := "rpc.server." + req.ServiceName + "#" + req.Method
			span := tracer.StartSpan(operationName, opentracing.ChildOf(spanCtx))
			defer span.Finish()
			return next(opentracing.ContextWithSpan(ctx, span), req)
		}
		return next(ctx, req)
	}
}

type ClientSideTracingFilterBuilder struct {
}

func (t *ClientSideTracingFilterBuilder) CtxReader(ctx context.Context) map[string]string {
	span := opentracing.SpanFromContext(ctx)
	res := make(map[string]string, 4)
	carrier := opentracing.TextMapCarrier(res)
	err := opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, carrier)
	if err != nil {
		fmt.Printf("could not read tracing ctx: %v", err)
	}
	return res
}

func (t *ClientSideTracingFilterBuilder) Build(next Filter) Filter {
	return func(ctx context.Context, req *dto.Request) (*dto.Response, error) {
		operationName := "rpc.client." + req.ServiceName + "#" + req.Method
		span, spanCtx := opentracing.StartSpanFromContext(ctx, operationName)
		fmt.Printf("this is client tracing filter\n")
		defer span.Finish()
		return next(spanCtx, req)
	}
}

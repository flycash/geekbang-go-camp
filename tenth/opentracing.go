package tenth

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/openzipkin/zipkin-go"
	"github.com/uber/jaeger-client-go"
	"net/http"
	"time"

	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"

	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

func OpenTracing() {
	// 从根开始
	span := opentracing.StartSpan("my-operation")
	defer span.Finish()

	// 执行一些你的业务操作，然后再记录一点东西
	span.LogFields(log.String("error", "my-error"))
	span.LogFields(log.Int("status", 404))

	childSpan := opentracing.StartSpan("child-operation", opentracing.ChildOf(span.Context()))
	defer childSpan.Finish()

	// 执行一些你的业务操作，然后再记录一点东西
	childSpan.LogFields(log.String("user_id", "uid_123"))
	childSpan.SetTag("user_id", "uid_123")
	childSpan.SetBaggageItem("baggage-item", "ba-123")

	// 传递一个 context
	SomeBusiness(opentracing.ContextWithSpan(context.Background(), childSpan))
}

func SomeBusiness(ctx context.Context) {
	// 尝试利用 ctx 重建 tracing 的上下文。
	// StartSpanFromContext 如果发现 ctx 里面已经有一个 span 了，那么这个 span 就会作为 父 span
	// 如果没有就是创建一个根 span
	span, ctx := opentracing.StartSpanFromContext(ctx, "some-business")
	defer span.Finish()
}

// --- span start ---
// inject tracing ——可以不必自己再开一个 span
// --- span end
func InjectTracing(ctx context.Context) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		httpClient := &http.Client{}
		httpReq, _ := http.NewRequest("GET", "http://your request/", nil)

		// Transmit the span's TraceContext as HTTP headers on our
		// outbound request.
		err := opentracing.GlobalTracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(httpReq.Header))
		if err != nil {
			panic(err)
		}
		resp, err := httpClient.Do(httpReq)
		if err != nil {
			panic(err)
		}
		println(resp.Proto)
	}
}

func ExtractTracing(req *http.Request) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var serverSpan opentracing.Span

		wireContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(req.Header))
		if err != nil {
			// tracing 链路构建不起来，
		}

		// 开启自己本地的 span
		serverSpan = opentracing.StartSpan(
			"my operation",
			ext.RPCServerOption(wireContext))

		defer serverSpan.Finish()
		ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)
		SomeBusiness(ctx)
	})
}

// 在你的 main 函数里面，程序启动之前，调用这个方法
func initZipkin() {
	reporter := zipkinhttp.NewReporter("http://localhost:9411/api/v2/spans")
	endpoint, err := zipkin.NewEndpoint("geekbangZipkinTracingService",
		"myservice.mydomain.com:80")
	if err != nil {
		panic(err)
	}
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		panic(err)
	}
	tracer := zipkinot.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(tracer)
}

func initJaeger() {
	cfg := jaegerConfig.Configuration{
		ServiceName: "geekbangJaegerTracingService",
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeRemote,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LocalAgentHostPort:  "127.0.0.1:6831",
			LogSpans:            true,
			BufferFlushInterval: 5 * time.Second,
		},
	}
	nativeTracer, _, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(nativeTracer)
}
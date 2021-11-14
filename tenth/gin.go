package tenth

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
)

func TracingMiddleware(ctx *gin.Context) {
	span := opentracing.StartSpan("http_request")
	span.SetTag("http.method", ctx.Request.Method)
	// ...
}

type MetricsMiddlewareBuilder struct {
	s prometheus.SummaryVec
	h prometheus.HistogramVec
}

func (m *MetricsMiddlewareBuilder) Handle(ctx *gin.Context) {

}

func main() {
	g := gin.New()
	m := &MetricsMiddlewareBuilder{
		// 在这里初始化你的 vector
	}
	g.Use(TracingMiddleware, m.Handle)
	g.Run()
}

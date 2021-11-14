package tenth

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func Counter() {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "my-namespace",
		Subsystem: "my-subsystem",
		Name: "test-counter",
	})
	prometheus.MustRegister(counter)
	// +1
	counter.Inc()
	// 必须是正数
	counter.Add(10.2)
}

func Gauge() {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "my-namespace",
		Subsystem: "my-subsystem",
		Name: "test-gauge",
	})
	prometheus.MustRegister(gauge)
	gauge.Set(12)
	gauge.Add(10.2)
	gauge.Add(-3)
	gauge.Sub(3)

}

func Histogram() {
	hist := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "my-namespace",
		Subsystem: "my-subsystem",
		Name: "test-histogram",
		// 按照这个来分桶
		Buckets: []float64{10, 50, 100, 200, 500, 1000, 10000},
	})
	prometheus.MustRegister(hist)
	hist.Observe(12.4)
}

func Summary()  {
	s := prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "my-namespace",
		Subsystem: "my-subsystem",
		Name: "test-summary",
		// key 是百分比，value 是误差，比如说 0.5 - 0.01，
		// 那么它实际上计算的可能是 0.49 到0.51之间的
		Objectives: map[float64]float64{
			0.5:   0.01,
			0.75:  0.01,
			0.90:  0.005,
			0.98:  0.002,
			0.99:  0.001,
			0.999: 0.0001,
		},
	})
	prometheus.MustRegister(s)
	s.Observe(12.3)
}

func StartPrometheusListener() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Vector() {
	summaryVec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:      "geekbang",
		Subsystem: "http_request",
		ConstLabels: map[string]string{
			"server":  "localhost:9091",
			"env":     "test",
			"appname": "test_app",
		},
		Help: "The statics info for http request",
	}, []string{"pattern", "method", "status"})

	summaryVec.WithLabelValues("/user/:id", "POST", "200").Observe(128)
}

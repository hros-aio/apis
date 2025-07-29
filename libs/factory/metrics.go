package factory

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tinh-tinh/prompt"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

const (
	HTTP_REQUESTS_TOTAL           = "http_requests_total"
	HTTP_ACTIVE_REQUESTS          = "http_active_requests"
	HTTP_REQUEST_DURATION_SECONDS = "http_request_duration_seconds"
	POST_REQUEST_DURATION_SECONDS = "post_request_duration_seconds"
)

func Metric() core.Modules {
	return prompt.Register(&prompt.Config{
		Metrics: []prompt.Metric{
			{
				Name: HTTP_REQUESTS_TOTAL,
				Collector: prometheus.NewCounterVec(prometheus.CounterOpts{
					Name: HTTP_REQUESTS_TOTAL,
					Help: "Total number of HTTP requests received",
				}, []string{"path", "method"}),
			},
			{
				Name: HTTP_ACTIVE_REQUESTS,
				Collector: prometheus.NewGauge(
					prometheus.GaugeOpts{
						Name: HTTP_ACTIVE_REQUESTS,
						Help: "Number of active connections to the service",
					},
				),
			},
			{
				Name: HTTP_REQUEST_DURATION_SECONDS,
				Collector: prometheus.NewHistogramVec(prometheus.HistogramOpts{
					Name:    HTTP_REQUEST_DURATION_SECONDS,
					Help:    "Duration of HTTP requests",
					Buckets: prometheus.DefBuckets,
				}, []string{"path", "method"}),
			},
			{
				Name: POST_REQUEST_DURATION_SECONDS,
				Collector: prometheus.NewSummary(prometheus.SummaryOpts{
					Name: POST_REQUEST_DURATION_SECONDS,
					Help: "Duration of requests to service",
					Objectives: map[float64]float64{
						0.5:  0.05,  // Median (50th percentile) with a 5% tolerance
						0.9:  0.01,  // 90th percentile with a 1% tolerance
						0.99: 0.001, // 99th percentile with a 0.1% tolerance
					},
				}),
			},
		},
	})
}

func MetricMiddleware(ctx core.Ctx) error {
	method := ctx.Req().Method
	path := ctx.Req().URL.Path

	counter := prompt.InjectCounterVec(ctx, HTTP_REQUESTS_TOTAL)
	if counter != nil {

		counter.WithLabelValues(path, method).Inc()
	}

	histogram := prompt.InjectHistogramVec(ctx, HTTP_REQUEST_DURATION_SECONDS)
	if histogram != nil {
		now := time.Now()

		delay := time.Duration(rand.Intn(900)) * time.Millisecond
		time.Sleep(delay)

		histogram.With(prometheus.Labels{
			"method": method, "path": path,
		}).Observe(time.Since(now).Seconds())
	}

	summary := prompt.InjectSummary(ctx, POST_REQUEST_DURATION_SECONDS)
	if summary != nil {
		now := time.Now()

		summary.Observe(time.Since(now).Seconds())
	}
	return ctx.Next()
}

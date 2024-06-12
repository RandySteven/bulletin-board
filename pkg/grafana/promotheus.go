package grafana

import "github.com/prometheus/client_golang/prometheus"

type HttpRequestMetrics struct {
	CounterVec *prometheus.CounterVec
	Histogram  *prometheus.HistogramVec
}

var requestMetrics *HttpRequestMetrics

func NewHttpRequestMetrics() *HttpRequestMetrics {
	return &HttpRequestMetrics{
		CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "http_requests_total_with_path",
			Help: "Number of HTTP requests by path.",
		}, []string{"path", "method"}),
		Histogram: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Response time of HTTP request.",
		}, []string{"path", "method"}),
	}
}

func InitGrafana() *HttpRequestMetrics {
	if requestMetrics == nil {
		requestMetrics = NewHttpRequestMetrics()
		prometheus.MustRegister(requestMetrics.CounterVec, requestMetrics.Histogram)
	}
	return requestMetrics
}

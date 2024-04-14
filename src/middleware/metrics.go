package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var vangoghRegistry = prometheus.NewRegistry()

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_milliseconds",
			Help:    "Duration of HTTP requests in milliseconds",
			Buckets: []float64{50, 80, 100, 250, 500, 1000, 2000, 3000},
		},
		[]string{"route"})

	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"route"},
	)
)

func init() {
	vangoghRegistry.MustRegister(requestDuration)
	vangoghRegistry.MustRegister(requestsTotal)
}

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		if c.Request.RequestURI != "/metrics" {
			requestsTotal.WithLabelValues(c.Request.URL.Path).Inc()
		}

		if c.Request.Method == "POST" {
			duration := time.Since(start).Milliseconds()
			requestDuration.WithLabelValues(c.Request.URL.Path).Observe(float64(duration))
		}
	}
}

func GetCustomRegistry() *prometheus.Registry {
	return vangoghRegistry
}

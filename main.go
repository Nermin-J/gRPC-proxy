package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests received",
		},
		[]string{"endpoint"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time for HTTP requests",
			Buckets: prometheus.LinearBuckets(0.1, 0.1, 10), // Buckets: 0.1s to 1s
		},
		[]string{"endpoint"},
	)
)

func init() {
	// Register metrics with Prometheus
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Simulating random response time
	sleepTime := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepTime)

	// Update metrics
	requestCount.WithLabelValues(r.URL.Path).Inc()
	requestDuration.WithLabelValues(r.URL.Path).Observe(time.Since(start).Seconds())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Go!"))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Expose metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Example endpoints generating random request durations
	http.HandleFunc("/api/fast", handler)
	http.HandleFunc("/api/slow", handler)

	// Start HTTP server
	serverAddr := ":8080"
	println("Server running at http://localhost" + serverAddr)
	http.ListenAndServe(serverAddr, nil)
}
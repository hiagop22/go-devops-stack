package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type KorpProjectResponse struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"},
	)
)

func KorpProjectHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	response := KorpProjectResponse{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

	duration := time.Since(start).Seconds()
	httpRequestDuration.WithLabelValues("/projeto-korp", "GET", "200").Observe(duration)
	httpRequestsTotal.WithLabelValues("/projeto-korp", "GET", "200").Inc()
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})

	duration := time.Since(start).Seconds()
	httpRequestDuration.WithLabelValues("/health", "GET", "200").Observe(duration)

	httpRequestsTotal.WithLabelValues("/health", "GET", "200").Inc()
}

func main() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)

	mux := http.NewServeMux()
	mux.Handle("GET /health", http.HandlerFunc(HealthHandler))
	mux.Handle("GET /projeto-korp", http.HandlerFunc(KorpProjectHandler))
	mux.Handle("GET /metrics", promhttp.Handler())

	log.Println("Server listenning on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

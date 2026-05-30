package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type KorpProjectResponse struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func KorpProjectHandler(w http.ResponseWriter, r *http.Request) {
	response := KorpProjectResponse{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", HealthHandler)
	mux.HandleFunc("GET /projeto-korp", KorpProjectHandler)
	mux.Handle("GET /metrics", promhttp.Handler())

	log.Println("Server listenning on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"infrasnap-agent/system" // Update if using a module
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics, err := system.GetMetrics()
	if err != nil {
		http.Error(w, "Failed to gather metrics", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)
	log.Println("InfraSnap agent running at http://localhost:3000/metrics")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

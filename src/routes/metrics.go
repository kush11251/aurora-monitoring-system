package routes

import (
    "encoding/json"
    "log"
    "net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
    metrics := []Metric{{Name: "CPU", Value: 90}, {Name: "Memory", Value: 80}}
    json.NewEncoder(w).Encode(metrics)
}

type Metric struct {
    Name  string `json:"name"`
    Value int    `json:"value"`
}

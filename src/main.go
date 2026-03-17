package main

import (
    "aurora-monitoring-system/src/routes"
    "log"
    "net/http"
)

func main() {
    log.Println("Server is starting...")
    http.HandleFunc("/metrics", routes.MetricsHandler)
    http.ListenAndServe(":8080", nil)
}

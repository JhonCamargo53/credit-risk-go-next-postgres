package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/middlewares"
)

type MetricsResponse struct {
	Timestamp     time.Time                    `json:"timestamp"`
	TotalRequests uint64                       `json:"total_requests"`
	Endpoints     []middlewares.EndpointMetric `json:"endpoints"`
}

func Metrics(w http.ResponseWriter, r *http.Request) {
	resp := MetricsResponse{
		Timestamp:     time.Now(),
		TotalRequests: middlewares.GetTotalRequests(),
		Endpoints:     middlewares.GetEndpointMetrics(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

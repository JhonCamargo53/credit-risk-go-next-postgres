package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	StatusCode   int
	BytesWritten int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.StatusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func (rr *responseRecorder) Write(b []byte) (int, error) {
	if rr.StatusCode == 0 {
		rr.StatusCode = http.StatusOK
	}
	n, err := rr.ResponseWriter.Write(b)
	rr.BytesWritten += n
	return n, err
}

var totalRequests uint64

var (
	mu                 sync.RWMutex
	requestsByPath     = make(map[string]uint64)
	totalLatencyByPath = make(map[string]int64)
)

func GetTotalRequests() uint64 {
	return atomic.LoadUint64(&totalRequests)
}

type EndpointMetric struct {
	Path         string  `json:"path"`
	Requests     uint64  `json:"requests"`
	AvgLatencyMs float64 `json:"avg_latency_ms"`
}

func GetEndpointMetrics() []EndpointMetric {
	mu.RLock()
	defer mu.RUnlock()

	metrics := make([]EndpointMetric, 0, len(requestsByPath))
	for path, count := range requestsByPath {
		totalLat := totalLatencyByPath[path]
		avg := 0.0
		if count > 0 {
			avg = float64(totalLat) / float64(count)
		}
		metrics = append(metrics, EndpointMetric{
			Path:         path,
			Requests:     count,
			AvgLatencyMs: avg,
		})
	}
	return metrics
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		rr := &responseRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		fmt.Printf("[REQ] %s %s\n", r.Method, r.RequestURI)

		next.ServeHTTP(rr, r)

		latency := time.Since(start)
		latencyMs := latency.Milliseconds()

		fmt.Printf("[DONE] %s %s -> %d in %v\n",
			r.Method, r.RequestURI, rr.StatusCode, latency)

		atomic.AddUint64(&totalRequests, 1)

		mu.Lock()
		requestsByPath[r.URL.Path]++
		totalLatencyByPath[r.URL.Path] += latencyMs
		mu.Unlock()
		entry := map[string]interface{}{
			"timestamp":  time.Now().Format(time.RFC3339),
			"level":      "info",
			"event":      "http_request",
			"method":     r.Method,
			"path":       r.URL.Path,
			"query":      r.URL.RawQuery,
			"user_agent": r.UserAgent(),
			"remote_ip":  r.RemoteAddr,
			"status":     rr.StatusCode,
			"bytes_sent": rr.BytesWritten,
			"latency_ms": latencyMs,
		}

		if data, err := json.Marshal(entry); err == nil {
			log.Println(string(data))
		} else {
			log.Printf("error serializando log: %v\n", err)
		}
	})
}

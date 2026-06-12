package httptransport

import "net/http"

func MetricsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"status":  "placeholder",
			"message": "Metrics exporter is not enabled yet. This endpoint reserves the local metrics path for future Prometheus integration.",
		})
	}
}

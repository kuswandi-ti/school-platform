package httptransport

import (
	"net/http"

	"school-platform/services/api-gateway/internal/response"
)

func MetricsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{
			"status":  "placeholder",
			"message": "Metrics exporter is not enabled yet. Local Prometheus support can scrape this path in a future sprint.",
		}, nil)
	}
}

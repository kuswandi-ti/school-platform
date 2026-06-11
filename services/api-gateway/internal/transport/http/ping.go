package httptransport

import (
	"net/http"

	"school-platform/services/api-gateway/internal/response"
)

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{"message": "pong"}, nil)
	}
}

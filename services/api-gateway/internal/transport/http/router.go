package httptransport

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"school-platform/services/api-gateway/internal/client"
	"school-platform/services/api-gateway/internal/config"
	"school-platform/services/api-gateway/internal/middleware"
	"school-platform/services/api-gateway/internal/response"
)

func NewRouter(cfg config.Config, logger *slog.Logger, identityClient client.Identity, authValidator *middleware.JWTValidator) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestContext)
	router.Use(middleware.Recover(logger))
	router.Use(middleware.CORS(cfg.CORSAllowedOrigins))
	router.Use(middleware.AccessLog(logger))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		response.Error(w, http.StatusNotFound, response.CodeNotFound, "Resource not found.", nil)
	})

	router.Get("/healthz", HealthHandler())
	router.Get("/readyz", ReadinessHandler())
	router.Get("/metrics", MetricsHandler())

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", PingHandler())
		r.Post("/auth/login", LoginHandler(identityClient))
		r.Post("/auth/refresh", RefreshHandler(identityClient))
		r.With(middleware.RequireAuth(authValidator)).Post("/auth/logout", LogoutHandler(identityClient))
	})

	return router
}

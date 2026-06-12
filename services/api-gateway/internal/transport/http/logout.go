package httptransport

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	identityv1 "school-platform/packages/proto/gen/go/identity/v1"
	"school-platform/services/api-gateway/internal/client"
	"school-platform/services/api-gateway/internal/middleware"
	"school-platform/services/api-gateway/internal/response"
)

const maxLogoutBodyBytes = 16 * 1024

type logoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,max=1024"`
}

type logoutResponse struct {
	LoggedOut bool `json:"logged_out"`
}

func LogoutHandler(identityClient client.Identity) http.HandlerFunc {
	validate := validator.New()
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := middleware.ActorContextFromContext(r.Context()); !ok {
			response.Error(w, http.StatusForbidden, response.CodeForbidden, "Actor context tidak tersedia.", nil)
			return
		}
		accessToken, ok := middleware.AccessTokenFromContext(r.Context())
		if !ok {
			response.Error(w, http.StatusForbidden, response.CodeForbidden, "Access token tidak tersedia.", nil)
			return
		}

		var request logoutRequest
		decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxLogoutBodyBytes))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Data logout tidak valid.", nil)
			return
		}
		if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Data logout tidak valid.", nil)
			return
		}
		if err := validate.Struct(request); err != nil {
			response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Refresh token wajib diisi.", nil)
			return
		}

		_, err := identityClient.Logout(r.Context(), &identityv1.LogoutRequest{
			AccessToken:   accessToken,
			RefreshToken:  request.RefreshToken,
			RequestId:     middleware.RequestIDFromContext(r.Context()),
			CorrelationId: middleware.CorrelationIDFromContext(r.Context()),
		})
		if err != nil {
			writeLogoutError(w, err)
			return
		}
		response.JSON(w, http.StatusOK, logoutResponse{LoggedOut: true}, nil)
	}
}

func writeLogoutError(w http.ResponseWriter, err error) {
	switch status.Code(err) {
	case codes.InvalidArgument:
		response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Data logout tidak valid.", nil)
	case codes.Unauthenticated:
		response.Error(w, http.StatusUnauthorized, response.CodeUnauthorized, "Sesi tidak valid.", nil)
	case codes.PermissionDenied:
		response.Error(w, http.StatusForbidden, response.CodeForbidden, "Sesi bukan milik pengguna.", nil)
	case codes.Unavailable, codes.DeadlineExceeded:
		response.Error(w, http.StatusServiceUnavailable, response.CodeServiceUnavailable, "Layanan autentikasi tidak tersedia.", nil)
	default:
		response.Error(w, http.StatusInternalServerError, response.CodeInternalError, "Terjadi kesalahan internal.", nil)
	}
}

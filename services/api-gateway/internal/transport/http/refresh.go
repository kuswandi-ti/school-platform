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

const maxRefreshBodyBytes = 16 * 1024

type refreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,max=1024"`
}

type refreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

func RefreshHandler(identityClient client.Identity) http.HandlerFunc {
	validate := validator.New()
	return func(w http.ResponseWriter, r *http.Request) {
		var request refreshRequest
		decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxRefreshBodyBytes))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Data refresh token tidak valid.", nil)
			return
		}
		if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Data refresh token tidak valid.", nil)
			return
		}
		if err := validate.Struct(request); err != nil {
			response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Refresh token wajib diisi.", nil)
			return
		}

		result, err := identityClient.Refresh(r.Context(), &identityv1.RefreshRequest{
			RefreshToken:  request.RefreshToken,
			RequestId:     middleware.RequestIDFromContext(r.Context()),
			CorrelationId: middleware.CorrelationIDFromContext(r.Context()),
			IpAddress:     requestIPAddress(r),
			UserAgent:     r.UserAgent(),
		})
		if err != nil {
			writeRefreshError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, refreshResponse{
			AccessToken:  result.GetAccessToken(),
			RefreshToken: result.GetRefreshToken(),
			TokenType:    result.GetTokenType(),
			ExpiresIn:    result.GetExpiresIn(),
		}, nil)
	}
}

func writeRefreshError(w http.ResponseWriter, err error) {
	switch status.Code(err) {
	case codes.InvalidArgument:
		response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Data refresh token tidak valid.", nil)
	case codes.Unauthenticated:
		response.Error(w, http.StatusUnauthorized, response.CodeUnauthorized, "Refresh token tidak valid, kedaluwarsa, atau telah digunakan.", nil)
	case codes.PermissionDenied:
		response.Error(w, http.StatusForbidden, response.CodeForbidden, "Akun tidak aktif atau terkunci.", nil)
	case codes.Unavailable, codes.DeadlineExceeded:
		response.Error(w, http.StatusServiceUnavailable, response.CodeServiceUnavailable, "Layanan autentikasi tidak tersedia.", nil)
	default:
		response.Error(w, http.StatusInternalServerError, response.CodeInternalError, "Terjadi kesalahan internal.", nil)
	}
}

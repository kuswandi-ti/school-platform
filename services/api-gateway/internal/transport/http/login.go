package httptransport

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	identityv1 "school-platform/packages/proto/gen/go/identity/v1"
	"school-platform/services/api-gateway/internal/client"
	"school-platform/services/api-gateway/internal/middleware"
	"school-platform/services/api-gateway/internal/response"
)

const maxLoginBodyBytes = 16 * 1024

type loginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,max=1024"`
}

type loginResponse struct {
	UserID       string `json:"user_id"`
	DisplayName  string `json:"display_name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

func LoginHandler(identityClient client.Identity) http.HandlerFunc {
	validate := validator.New()
	return func(w http.ResponseWriter, r *http.Request) {
		var request loginRequest
		decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxLoginBodyBytes))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Data login tidak valid.", nil)
			return
		}
		if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Data login tidak valid.", nil)
			return
		}
		if err := validate.Struct(request); err != nil {
			response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Email dan password wajib diisi dengan format yang valid.", nil)
			return
		}

		result, err := identityClient.Login(r.Context(), &identityv1.LoginRequest{
			Email:         strings.TrimSpace(request.Email),
			Password:      request.Password,
			RequestId:     middleware.RequestIDFromContext(r.Context()),
			CorrelationId: middleware.CorrelationIDFromContext(r.Context()),
			IpAddress:     requestIPAddress(r),
			UserAgent:     r.UserAgent(),
		})
		if err != nil {
			writeLoginError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, loginResponse{
			UserID:       result.GetUserId(),
			DisplayName:  result.GetDisplayName(),
			AccessToken:  result.GetAccessToken(),
			RefreshToken: result.GetRefreshToken(),
			TokenType:    result.GetTokenType(),
			ExpiresIn:    result.GetExpiresIn(),
		}, nil)
	}
}

func writeLoginError(w http.ResponseWriter, err error) {
	switch status.Code(err) {
	case codes.InvalidArgument:
		response.Error(w, http.StatusBadRequest, response.CodeValidationError, "Data login tidak valid.", nil)
	case codes.Unauthenticated:
		response.Error(w, http.StatusUnauthorized, response.CodeUnauthorized, "Email atau password tidak valid.", nil)
	case codes.PermissionDenied:
		response.Error(w, http.StatusForbidden, response.CodeForbidden, "Akun tidak aktif atau terkunci.", nil)
	case codes.Unavailable, codes.DeadlineExceeded:
		response.Error(w, http.StatusServiceUnavailable, response.CodeServiceUnavailable, "Layanan autentikasi tidak tersedia.", nil)
	default:
		response.Error(w, http.StatusInternalServerError, response.CodeInternalError, "Terjadi kesalahan internal.", nil)
	}
}

func requestIPAddress(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}
	return r.RemoteAddr
}

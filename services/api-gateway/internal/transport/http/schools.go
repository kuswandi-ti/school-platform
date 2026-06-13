package httptransport

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	schoolcorev1 "school-platform/packages/proto/gen/go/schoolcore/v1"
	"school-platform/services/api-gateway/internal/client"
	"school-platform/services/api-gateway/internal/middleware"
	"school-platform/services/api-gateway/internal/response"
	"strings"
)

const maxSchoolBodyBytes = 32 * 1024

type createSchoolRequest struct {
	SchoolCode  string  `json:"school_code" validate:"required,max=50"`
	Name        string  `json:"name" validate:"required,max=150"`
	SchoolLevel string  `json:"school_level" validate:"required,oneof=kindergarten elementary junior_high senior_high"`
	NPSN        *string `json:"npsn" validate:"omitempty,max=50"`
	Address     *string `json:"address" validate:"omitempty,max=4000"`
	Phone       *string `json:"phone" validate:"omitempty,max=50"`
	Email       *string `json:"email" validate:"omitempty,email,max=255"`
	LogoFileID  *string `json:"logo_file_id" validate:"omitempty,uuid"`
	Status      string  `json:"status" validate:"required,oneof=active inactive"`
}
type updateSchoolRequest struct {
	SchoolCode  *string `json:"school_code" validate:"omitempty,max=50"`
	Name        *string `json:"name" validate:"omitempty,max=150"`
	SchoolLevel *string `json:"school_level" validate:"omitempty,oneof=kindergarten elementary junior_high senior_high"`
	NPSN        *string `json:"npsn" validate:"omitempty,max=50"`
	Address     *string `json:"address" validate:"omitempty,max=4000"`
	Phone       *string `json:"phone" validate:"omitempty,max=50"`
	Email       *string `json:"email" validate:"omitempty,email,max=255"`
	LogoFileID  *string `json:"logo_file_id" validate:"omitempty,uuid"`
	Status      *string `json:"status" validate:"omitempty,oneof=active inactive"`
}

func GetCurrentFoundationHandler(c client.SchoolCore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := c.GetCurrentFoundation(r.Context(), &schoolcorev1.GetCurrentFoundationRequest{Actor: schoolActor(r)})
		if err != nil {
			writeSchoolError(w, err)
			return
		}
		response.JSON(w, http.StatusOK, result.Foundation, nil)
	}
}
func ListSchoolsHandler(c client.SchoolCore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := c.ListSchools(r.Context(), &schoolcorev1.ListSchoolsRequest{Actor: schoolActor(r)})
		if err != nil {
			writeSchoolError(w, err)
			return
		}
		response.JSON(w, http.StatusOK, result.GetSchools(), map[string]int{"total": len(result.GetSchools())})
	}
}
func CreateSchoolHandler(c client.SchoolCore) http.HandlerFunc {
	validate := validator.New()
	return func(w http.ResponseWriter, r *http.Request) {
		var body createSchoolRequest
		if !decodeSchoolBody(w, r, &body) || validate.Struct(body) != nil {
			response.Error(w, http.StatusUnprocessableEntity, response.CodeValidationError, "Data sekolah tidak valid.", nil)
			return
		}
		result, err := c.CreateSchool(r.Context(), &schoolcorev1.CreateSchoolRequest{Actor: schoolActor(r), SchoolCode: strings.TrimSpace(body.SchoolCode), Name: strings.TrimSpace(body.Name), SchoolLevel: body.SchoolLevel, Npsn: body.NPSN, Address: body.Address, Phone: body.Phone, Email: body.Email, LogoFileId: body.LogoFileID, Status: body.Status})
		if err != nil {
			writeSchoolError(w, err)
			return
		}
		response.JSON(w, http.StatusCreated, result.GetSchool(), nil)
	}
}
func UpdateSchoolHandler(c client.SchoolCore) http.HandlerFunc {
	validate := validator.New()
	return func(w http.ResponseWriter, r *http.Request) {
		var body updateSchoolRequest
		if !decodeSchoolBody(w, r, &body) || validate.Struct(body) != nil || emptyUpdate(body) {
			response.Error(w, http.StatusUnprocessableEntity, response.CodeValidationError, "Data sekolah tidak valid.", nil)
			return
		}
		result, err := c.UpdateSchool(r.Context(), &schoolcorev1.UpdateSchoolRequest{Actor: schoolActor(r), SchoolId: chi.URLParam(r, "school_id"), SchoolCode: body.SchoolCode, Name: body.Name, SchoolLevel: body.SchoolLevel, Npsn: body.NPSN, Address: body.Address, Phone: body.Phone, Email: body.Email, LogoFileId: body.LogoFileID, Status: body.Status})
		if err != nil {
			writeSchoolError(w, err)
			return
		}
		response.JSON(w, http.StatusOK, result.GetSchool(), nil)
	}
}
func decodeSchoolBody(w http.ResponseWriter, r *http.Request, dst any) bool {
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxSchoolBodyBytes))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		return false
	}
	return errors.Is(decoder.Decode(&struct{}{}), io.EOF)
}
func emptyUpdate(v updateSchoolRequest) bool {
	return v.SchoolCode == nil && v.Name == nil && v.SchoolLevel == nil && v.NPSN == nil && v.Address == nil && v.Phone == nil && v.Email == nil && v.LogoFileID == nil && v.Status == nil
}
func schoolActor(r *http.Request) *schoolcorev1.ActorContext {
	a, _ := middleware.ActorContextFromContext(r.Context())
	return &schoolcorev1.ActorContext{UserId: a.UserID, FoundationId: a.FoundationID, SchoolId: a.SchoolID, Roles: a.Roles, Permissions: a.Permissions, RequestId: a.RequestID, CorrelationId: a.CorrelationID, IpAddress: requestIPAddress(r), UserAgent: r.UserAgent()}
}
func writeSchoolError(w http.ResponseWriter, err error) {
	switch status.Code(err) {
	case codes.InvalidArgument:
		response.Error(w, http.StatusUnprocessableEntity, response.CodeValidationError, "Data sekolah tidak valid.", nil)
	case codes.Unauthenticated:
		response.Error(w, http.StatusUnauthorized, response.CodeUnauthorized, "Autentikasi tidak valid.", nil)
	case codes.PermissionDenied:
		response.Error(w, http.StatusForbidden, response.CodeForbidden, "Anda tidak memiliki akses.", nil)
	case codes.NotFound:
		response.Error(w, http.StatusNotFound, response.CodeNotFound, "Sekolah tidak ditemukan.", nil)
	case codes.AlreadyExists:
		response.Error(w, http.StatusConflict, response.CodeConflict, "Kode sekolah atau NPSN sudah digunakan.", nil)
	case codes.Unavailable, codes.DeadlineExceeded:
		response.Error(w, http.StatusServiceUnavailable, response.CodeServiceUnavailable, "Layanan sekolah tidak tersedia.", nil)
	default:
		response.Error(w, http.StatusInternalServerError, response.CodeInternalError, "Terjadi kesalahan internal.", nil)
	}
}

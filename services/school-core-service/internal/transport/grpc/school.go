package grpctransport

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	schoolcorev1 "school-platform/packages/proto/gen/go/schoolcore/v1"
	"school-platform/services/school-core-service/internal/domain"
	"school-platform/services/school-core-service/internal/usecase"
	"time"
)

type SchoolCoreServer struct {
	schoolcorev1.UnimplementedSchoolCoreServiceServer
	service *usecase.Service
}

func NewSchoolCoreServer(s *usecase.Service) *SchoolCoreServer { return &SchoolCoreServer{service: s} }
func (s *SchoolCoreServer) GetCurrentFoundation(ctx context.Context, r *schoolcorev1.GetCurrentFoundationRequest) (*schoolcorev1.GetCurrentFoundationResponse, error) {
	a, err := actor(r.GetActor())
	if err != nil {
		return nil, err
	}
	v, err := s.service.GetCurrentFoundation(ctx, a)
	if err != nil {
		return nil, mapError(err)
	}
	return &schoolcorev1.GetCurrentFoundationResponse{Foundation: foundation(v)}, nil
}
func (s *SchoolCoreServer) ListSchools(ctx context.Context, r *schoolcorev1.ListSchoolsRequest) (*schoolcorev1.ListSchoolsResponse, error) {
	a, err := actor(r.GetActor())
	if err != nil {
		return nil, err
	}
	items, err := s.service.ListSchools(ctx, a)
	if err != nil {
		return nil, mapError(err)
	}
	out := make([]*schoolcorev1.School, 0, len(items))
	for _, v := range items {
		out = append(out, school(v))
	}
	return &schoolcorev1.ListSchoolsResponse{Schools: out}, nil
}
func (s *SchoolCoreServer) CreateSchool(ctx context.Context, r *schoolcorev1.CreateSchoolRequest) (*schoolcorev1.CreateSchoolResponse, error) {
	a, err := actor(r.GetActor())
	if err != nil {
		return nil, err
	}
	logo, err := optionalUUID(r.LogoFileId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid logo_file_id")
	}
	v, err := s.service.CreateSchool(ctx, a, usecase.CreateSchoolInput{SchoolCode: r.SchoolCode, Name: r.Name, SchoolLevel: r.SchoolLevel, NPSN: r.Npsn, Address: r.Address, Phone: r.Phone, Email: r.Email, LogoFileID: logo, Status: r.Status})
	if err != nil {
		return nil, mapError(err)
	}
	return &schoolcorev1.CreateSchoolResponse{School: school(v)}, nil
}
func (s *SchoolCoreServer) UpdateSchool(ctx context.Context, r *schoolcorev1.UpdateSchoolRequest) (*schoolcorev1.UpdateSchoolResponse, error) {
	a, err := actor(r.GetActor())
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(r.GetSchoolId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid school_id")
	}
	logo, err := optionalUUID(r.LogoFileId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid logo_file_id")
	}
	v, err := s.service.UpdateSchool(ctx, a, usecase.UpdateSchoolInput{SchoolID: id, SchoolCode: r.SchoolCode, Name: r.Name, SchoolLevel: r.SchoolLevel, NPSN: r.Npsn, Address: r.Address, Phone: r.Phone, Email: r.Email, LogoFileID: logo, Status: r.Status, SetNPSN: r.Npsn != nil, SetAddress: r.Address != nil, SetPhone: r.Phone != nil, SetEmail: r.Email != nil, SetLogoFileID: r.LogoFileId != nil})
	if err != nil {
		return nil, mapError(err)
	}
	return &schoolcorev1.UpdateSchoolResponse{School: school(v)}, nil
}
func actor(v *schoolcorev1.ActorContext) (domain.Actor, error) {
	if v == nil {
		return domain.Actor{}, status.Error(codes.Unauthenticated, "actor is required")
	}
	userID, e1 := uuid.Parse(v.GetUserId())
	foundationID, e2 := uuid.Parse(v.GetFoundationId())
	if e1 != nil || e2 != nil {
		return domain.Actor{}, status.Error(codes.Unauthenticated, "invalid actor")
	}
	var schoolID *uuid.UUID
	if v.GetSchoolId() != "" {
		id, err := uuid.Parse(v.GetSchoolId())
		if err != nil {
			return domain.Actor{}, status.Error(codes.Unauthenticated, "invalid actor school")
		}
		schoolID = &id
	}
	return domain.Actor{UserID: userID, FoundationID: foundationID, SchoolID: schoolID, Roles: v.GetRoles(), Permissions: v.GetPermissions(), RequestID: v.GetRequestId(), CorrelationID: v.GetCorrelationId(), IPAddress: text(v.GetIpAddress()), UserAgent: text(v.GetUserAgent())}, nil
}
func mapError(err error) error {
	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, usecase.ErrForbidden):
		return status.Error(codes.PermissionDenied, "forbidden")
	case errors.Is(err, usecase.ErrInvalidInput):
		return status.Error(codes.InvalidArgument, "invalid school data")
	case errors.Is(err, pgx.ErrNoRows):
		return status.Error(codes.NotFound, "resource not found")
	case errors.As(err, &pgErr) && pgErr.Code == "23505":
		return status.Error(codes.AlreadyExists, "school already exists")
	default:
		return status.Error(codes.Internal, "school operation failed")
	}
}
func optionalUUID(v *string) (*uuid.UUID, error) {
	if v == nil || *v == "" {
		return nil, nil
	}
	id, err := uuid.Parse(*v)
	return &id, err
}
func text(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
func str(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
func idstr(v *uuid.UUID) string {
	if v == nil {
		return ""
	}
	return v.String()
}
func foundation(v domain.Foundation) *schoolcorev1.Foundation {
	return &schoolcorev1.Foundation{Id: v.ID.String(), FoundationCode: v.FoundationCode, Name: v.Name, LegalName: str(v.LegalName), Address: str(v.Address), Phone: str(v.Phone), Email: str(v.Email), LogoFileId: idstr(v.LogoFileID), Timezone: v.Timezone, Status: v.Status, CreatedAt: v.CreatedAt.Format(time.RFC3339Nano), UpdatedAt: v.UpdatedAt.Format(time.RFC3339Nano)}
}
func school(v domain.School) *schoolcorev1.School {
	return &schoolcorev1.School{Id: v.ID.String(), FoundationId: v.FoundationID.String(), SchoolCode: v.SchoolCode, Name: v.Name, SchoolLevel: v.SchoolLevel, Npsn: str(v.NPSN), Address: str(v.Address), Phone: str(v.Phone), Email: str(v.Email), LogoFileId: idstr(v.LogoFileID), Status: v.Status, CreatedAt: v.CreatedAt.Format(time.RFC3339Nano), UpdatedAt: v.UpdatedAt.Format(time.RFC3339Nano)}
}

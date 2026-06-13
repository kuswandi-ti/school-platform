package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"school-platform/services/school-core-service/internal/domain"
	"school-platform/services/school-core-service/internal/repository"
)

var (
	ErrForbidden    = errors.New("forbidden")
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
)

type Store interface {
	GetFoundation(context.Context, uuid.UUID) (domain.Foundation, error)
	ListSchools(context.Context, uuid.UUID) ([]domain.School, error)
	GetSchool(context.Context, uuid.UUID, uuid.UUID) (domain.School, error)
	CreateSchool(context.Context, repository.SchoolWrite, repository.AuditContext) (domain.School, error)
	UpdateSchool(context.Context, repository.SchoolWrite, repository.AuditContext) (domain.School, error)
}
type Service struct{ store Store }

func NewService(store Store) *Service { return &Service{store: store} }
func (s *Service) GetCurrentFoundation(ctx context.Context, a domain.Actor) (domain.Foundation, error) {
	if !allowed(a, "school.foundation.view", true) {
		return domain.Foundation{}, ErrForbidden
	}
	return s.store.GetFoundation(ctx, a.FoundationID)
}
func (s *Service) ListSchools(ctx context.Context, a domain.Actor) ([]domain.School, error) {
	if !allowed(a, "school.school.view", true) {
		return nil, ErrForbidden
	}
	return s.store.ListSchools(ctx, a.FoundationID)
}

type CreateSchoolInput struct {
	SchoolCode, Name, SchoolLevel, Status string
	NPSN, Address, Phone, Email           *string
	LogoFileID                            *uuid.UUID
}

func (s *Service) CreateSchool(ctx context.Context, a domain.Actor, in CreateSchoolInput) (domain.School, error) {
	if !allowed(a, "school.school.manage", true) {
		return domain.School{}, ErrForbidden
	}
	if !validSchool(in.SchoolCode, in.Name, in.SchoolLevel, in.Status) {
		return domain.School{}, ErrInvalidInput
	}
	p := repository.SchoolWrite{ID: uuid.New(), FoundationID: a.FoundationID, SchoolCode: strings.TrimSpace(in.SchoolCode), Name: strings.TrimSpace(in.Name), SchoolLevel: in.SchoolLevel, NPSN: clean(in.NPSN), Address: clean(in.Address), Phone: clean(in.Phone), Email: clean(in.Email), LogoFileID: in.LogoFileID, Status: in.Status}
	return s.store.CreateSchool(ctx, p, repository.AuditContext{Actor: a, NewValues: schoolAuditValues(p), Action: "school.school.created"})
}

type UpdateSchoolInput struct {
	SchoolID                                               uuid.UUID
	SchoolCode, Name, SchoolLevel, Status                  *string
	NPSN, Address, Phone, Email                            *string
	LogoFileID                                             *uuid.UUID
	SetNPSN, SetAddress, SetPhone, SetEmail, SetLogoFileID bool
}

func (s *Service) UpdateSchool(ctx context.Context, a domain.Actor, in UpdateSchoolInput) (domain.School, error) {
	if !allowed(a, "school.school.manage", true) {
		return domain.School{}, ErrForbidden
	}
	current, err := s.store.GetSchool(ctx, in.SchoolID, a.FoundationID)
	if err != nil {
		return domain.School{}, err
	}
	next := repository.SchoolWrite{ID: current.ID, FoundationID: current.FoundationID, SchoolCode: current.SchoolCode, Name: current.Name, SchoolLevel: current.SchoolLevel, NPSN: current.NPSN, Address: current.Address, Phone: current.Phone, Email: current.Email, LogoFileID: current.LogoFileID, Status: current.Status}
	if in.SchoolCode != nil {
		next.SchoolCode = strings.TrimSpace(*in.SchoolCode)
	}
	if in.Name != nil {
		next.Name = strings.TrimSpace(*in.Name)
	}
	if in.SchoolLevel != nil {
		next.SchoolLevel = *in.SchoolLevel
	}
	if in.Status != nil {
		next.Status = *in.Status
	}
	if in.SetNPSN {
		next.NPSN = clean(in.NPSN)
	}
	if in.SetAddress {
		next.Address = clean(in.Address)
	}
	if in.SetPhone {
		next.Phone = clean(in.Phone)
	}
	if in.SetEmail {
		next.Email = clean(in.Email)
	}
	if in.SetLogoFileID {
		next.LogoFileID = in.LogoFileID
	}
	if !validSchool(next.SchoolCode, next.Name, next.SchoolLevel, next.Status) {
		return domain.School{}, ErrInvalidInput
	}
	return s.store.UpdateSchool(ctx, next, repository.AuditContext{Actor: a, OldValues: schoolDomainAuditValues(current), NewValues: schoolAuditValues(next), Action: "school.school.updated"})
}
func allowed(a domain.Actor, permission string, admin bool) bool {
	if a.UserID == uuid.Nil || a.FoundationID == uuid.Nil || a.RequestID == "" || a.CorrelationID == "" {
		return false
	}
	if admin && !contains(a.Roles, "admin_yayasan") {
		return false
	}
	if contains(a.Permissions, permission) {
		return true
	}
	return strings.HasSuffix(permission, ".view") && contains(a.Permissions, strings.TrimSuffix(permission, ".view")+".manage")
}
func contains(values []string, want string) bool {
	for _, v := range values {
		if v == want {
			return true
		}
	}
	return false
}
func validSchool(code, name, level, status string) bool {
	return strings.TrimSpace(code) != "" && strings.TrimSpace(name) != "" && contains([]string{"kindergarten", "elementary", "junior_high", "senior_high"}, level) && contains([]string{"active", "inactive"}, status)
}
func clean(v *string) *string {
	if v == nil {
		return nil
	}
	s := strings.TrimSpace(*v)
	if s == "" {
		return nil
	}
	return &s
}

func schoolAuditValues(v repository.SchoolWrite) map[string]any {
	return map[string]any{"school_code": v.SchoolCode, "name": v.Name, "school_level": v.SchoolLevel, "npsn": v.NPSN, "address": v.Address, "phone": v.Phone, "email": v.Email, "logo_file_id": v.LogoFileID, "status": v.Status}
}

func schoolDomainAuditValues(v domain.School) map[string]any {
	return map[string]any{"school_code": v.SchoolCode, "name": v.Name, "school_level": v.SchoolLevel, "npsn": v.NPSN, "address": v.Address, "phone": v.Phone, "email": v.Email, "logo_file_id": v.LogoFileID, "status": v.Status}
}

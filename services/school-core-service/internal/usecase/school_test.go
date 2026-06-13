package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"school-platform/services/school-core-service/internal/domain"
	"school-platform/services/school-core-service/internal/repository"
	"testing"
)

type storeStub struct {
	foundation domain.Foundation
	schools    []domain.School
	school     domain.School
	created    repository.SchoolWrite
	updated    repository.SchoolWrite
}

func (s *storeStub) GetFoundation(context.Context, uuid.UUID) (domain.Foundation, error) {
	return s.foundation, nil
}
func (s *storeStub) ListSchools(context.Context, uuid.UUID) ([]domain.School, error) {
	return s.schools, nil
}
func (s *storeStub) GetSchool(context.Context, uuid.UUID, uuid.UUID) (domain.School, error) {
	if s.school.ID == uuid.Nil {
		return domain.School{}, errors.New("missing")
	}
	return s.school, nil
}
func (s *storeStub) CreateSchool(_ context.Context, p repository.SchoolWrite, _ repository.AuditContext) (domain.School, error) {
	s.created = p
	return domain.School{ID: p.ID, FoundationID: p.FoundationID, SchoolCode: p.SchoolCode}, nil
}
func (s *storeStub) UpdateSchool(_ context.Context, p repository.SchoolWrite, _ repository.AuditContext) (domain.School, error) {
	s.updated = p
	return domain.School{ID: p.ID, FoundationID: p.FoundationID, Name: p.Name}, nil
}
func actorFor(role, permission string) domain.Actor {
	return domain.Actor{UserID: uuid.New(), FoundationID: uuid.New(), Roles: []string{role}, Permissions: []string{permission}, RequestID: "req", CorrelationID: "corr"}
}
func TestCreateSchoolRequiresAdminAndPermission(t *testing.T) {
	tests := []struct{ name, role, permission string }{{"non admin", "tu_staff", "school.school.manage"}, {"missing permission", "admin_yayasan", "school.school.view"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewService(&storeStub{}).CreateSchool(context.Background(), actorFor(tt.role, tt.permission), CreateSchoolInput{SchoolCode: "SD", Name: "SD Test", SchoolLevel: "elementary", Status: "active"})
			require.ErrorIs(t, err, ErrForbidden)
		})
	}
}
func TestCreateSchoolUsesActorFoundationScope(t *testing.T) {
	store := &storeStub{}
	a := actorFor("admin_yayasan", "school.school.manage")
	result, err := NewService(store).CreateSchool(context.Background(), a, CreateSchoolInput{SchoolCode: " SD01 ", Name: " SD Test ", SchoolLevel: "elementary", Status: "active"})
	require.NoError(t, err)
	require.Equal(t, a.FoundationID, store.created.FoundationID)
	require.Equal(t, "SD01", result.SchoolCode)
}
func TestListSchoolsRejectsMissingViewPermission(t *testing.T) {
	_, err := NewService(&storeStub{}).ListSchools(context.Background(), actorFor("admin_yayasan", "school.foundation.view"))
	require.ErrorIs(t, err, ErrForbidden)
}

func TestListSchoolsRejectsNonAdminWithViewPermission(t *testing.T) {
	_, err := NewService(&storeStub{}).ListSchools(context.Background(), actorFor("kepala_sekolah", "school.school.view"))
	require.ErrorIs(t, err, ErrForbidden)
}
func TestUpdateSchoolKeepsFoundationScope(t *testing.T) {
	foundationID := uuid.New()
	schoolID := uuid.New()
	store := &storeStub{school: domain.School{ID: schoolID, FoundationID: foundationID, SchoolCode: "SD", Name: "Old", SchoolLevel: "elementary", Status: "active"}}
	a := actorFor("admin_yayasan", "school.school.manage")
	a.FoundationID = foundationID
	name := "New"
	_, err := NewService(store).UpdateSchool(context.Background(), a, UpdateSchoolInput{SchoolID: schoolID, Name: &name})
	require.NoError(t, err)
	require.Equal(t, foundationID, store.updated.FoundationID)
	require.Equal(t, "New", store.updated.Name)
}

package repository_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"school-platform/services/identity-service/internal/repository"
)

func TestAuthorizationRepositoryAssignsSchoolScopedRoleAndAudits(t *testing.T) {
	users, pool, ctx := newTestRepository(t)
	actorID := createAuthorizationUser(t, users, ctx, "actor@example.com")
	userID := createAuthorizationUser(t, users, ctx, "principal@example.com")
	foundationID := uuid.New()
	schoolID := uuid.New()
	repo := repository.NewAuthorizationRepository(pool)

	assignment, err := repo.AssignRole(ctx, repository.AssignRoleParams{
		ID: uuid.New(), UserID: userID, RoleCode: "kepala_sekolah",
		FoundationID: foundationID, SchoolID: &schoolID,
		ActorUserID: actorID, ActorRole: "admin_yayasan",
		RequestID: "request-assign-school", CorrelationID: "correlation-assign-school",
	})
	require.NoError(t, err)
	require.Equal(t, "kepala_sekolah", assignment.Role)
	require.Equal(t, foundationID, assignment.FoundationID)
	require.Equal(t, schoolID, *assignment.SchoolID)

	var storedSchoolID uuid.UUID
	require.NoError(t, pool.QueryRow(ctx, `
		SELECT school_id FROM user_role_assignments WHERE id = $1`, assignment.ID).Scan(&storedSchoolID))
	require.Equal(t, schoolID, storedSchoolID)

	var action, requestID, correlationID string
	require.NoError(t, pool.QueryRow(ctx, `
		SELECT action, request_id, correlation_id
		FROM identity_audit_logs
		WHERE entity_type = 'user_role_assignment' AND entity_id = $1`, assignment.ID).
		Scan(&action, &requestID, &correlationID))
	require.Equal(t, "identity.role.assigned", action)
	require.Equal(t, "request-assign-school", requestID)
	require.Equal(t, "correlation-assign-school", correlationID)
}

func TestAuthorizationRepositoryRejectsInvalidSchoolScope(t *testing.T) {
	users, pool, ctx := newTestRepository(t)
	actorID := createAuthorizationUser(t, users, ctx, "invalid-actor@example.com")
	userID := createAuthorizationUser(t, users, ctx, "invalid-principal@example.com")
	repo := repository.NewAuthorizationRepository(pool)

	_, err := repo.AssignRole(ctx, repository.AssignRoleParams{
		ID: uuid.New(), UserID: userID, RoleCode: "kepala_sekolah", FoundationID: uuid.New(),
		ActorUserID: actorID, ActorRole: "admin_yayasan",
		RequestID: "request-invalid-scope", CorrelationID: "correlation-invalid-scope",
	})
	require.ErrorIs(t, err, repository.ErrInvalidRoleScope)

	var assignments, audits int
	require.NoError(t, pool.QueryRow(ctx, "SELECT COUNT(*) FROM user_role_assignments").Scan(&assignments))
	require.NoError(t, pool.QueryRow(ctx, "SELECT COUNT(*) FROM identity_audit_logs").Scan(&audits))
	require.Zero(t, assignments)
	require.Zero(t, audits)
}

func TestAuthorizationRepositoryGetsRolesPermissionsAndScopes(t *testing.T) {
	users, pool, ctx := newTestRepository(t)
	actorID := createAuthorizationUser(t, users, ctx, "context-actor@example.com")
	userID := createAuthorizationUser(t, users, ctx, "teacher@example.com")
	foundationID := uuid.New()
	schoolID := uuid.New()
	classID := uuid.New()
	subjectID := uuid.New()
	repo := repository.NewAuthorizationRepository(pool)

	_, err := repo.AssignRole(ctx, repository.AssignRoleParams{
		ID: uuid.New(), UserID: userID, RoleCode: "guru",
		FoundationID: foundationID, SchoolID: &schoolID, ClassID: &classID, SubjectID: &subjectID,
		Scope:       json.RawMessage(`{"assignment":"subject_teacher"}`),
		ActorUserID: actorID, ActorRole: "admin_yayasan",
		RequestID: "request-teacher", CorrelationID: "correlation-teacher",
	})
	require.NoError(t, err)

	contextResult, err := repo.GetUserContext(ctx, userID, time.Now())
	require.NoError(t, err)
	require.Equal(t, userID, contextResult.UserID)
	require.Equal(t, []string{"guru"}, contextResult.Roles)
	require.Contains(t, contextResult.Permissions, "academic.attendance.manage")
	require.Contains(t, contextResult.Permissions, "academic.grade.manage")
	require.NotContains(t, contextResult.Permissions, "identity.role.assign")
	require.Len(t, contextResult.Assignments, 1)
	require.Equal(t, schoolID, *contextResult.Assignments[0].SchoolID)
	require.Equal(t, classID, *contextResult.Assignments[0].ClassID)
	require.Equal(t, subjectID, *contextResult.Assignments[0].SubjectID)
	require.JSONEq(t, `{"assignment":"subject_teacher"}`, string(contextResult.Assignments[0].Scope))
}

func createAuthorizationUser(t *testing.T, users *repository.UserRepository, ctx context.Context, email string) uuid.UUID {
	t.Helper()
	userID := uuid.New()
	_, err := users.CreateUser(ctx, repository.CreateUserParams{
		ID: userID, Email: email, PasswordHash: "test-password-hash",
		DisplayName: email, Status: "active",
	})
	require.NoError(t, err)
	return userID
}

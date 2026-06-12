package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"school-platform/services/identity-service/internal/db/sqlc"
	"school-platform/services/identity-service/internal/domain"
)

var ErrInvalidRoleScope = errors.New("invalid role assignment scope")

type AuthorizationRepository struct {
	db      dbsqlc.DBTX
	queries *dbsqlc.Queries
	now     func() time.Time
}

type AssignRoleParams struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	RoleCode      string
	FoundationID  uuid.UUID
	SchoolID      *uuid.UUID
	ClassID       *uuid.UUID
	StudentID     *uuid.UUID
	EmployeeID    *uuid.UUID
	SubjectID     *uuid.UUID
	Scope         json.RawMessage
	StartsAt      *time.Time
	EndsAt        *time.Time
	ActorUserID   uuid.UUID
	ActorRole     string
	RequestID     string
	CorrelationID string
	IPAddress     *string
	UserAgent     *string
}

func NewAuthorizationRepository(db dbsqlc.DBTX) *AuthorizationRepository {
	return &AuthorizationRepository{db: db, queries: dbsqlc.New(db), now: time.Now}
}

func (r *AuthorizationRepository) AssignRole(ctx context.Context, params AssignRoleParams) (domain.RoleAssignment, error) {
	if params.ID == uuid.Nil || params.UserID == uuid.Nil || params.FoundationID == uuid.Nil || params.ActorUserID == uuid.Nil || params.ActorRole == "" || params.RequestID == "" || params.CorrelationID == "" {
		return domain.RoleAssignment{}, ErrInvalidRoleScope
	}
	if params.EndsAt != nil && params.StartsAt != nil && !params.EndsAt.After(*params.StartsAt) {
		return domain.RoleAssignment{}, ErrInvalidRoleScope
	}

	beginner, ok := r.db.(interface {
		Begin(context.Context) (pgx.Tx, error)
	})
	if !ok {
		return domain.RoleAssignment{}, fmt.Errorf("assign role: database does not support transactions")
	}
	tx, err := beginner.Begin(ctx)
	if err != nil {
		return domain.RoleAssignment{}, fmt.Errorf("begin role assignment: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	queries := r.queries.WithTx(tx)
	role, err := queries.GetRoleByCode(ctx, params.RoleCode)
	if err != nil {
		return domain.RoleAssignment{}, fmt.Errorf("find assignment role: %w", err)
	}
	if err := validateRoleScope(role.Code, params); err != nil {
		return domain.RoleAssignment{}, err
	}

	assignment, err := queries.CreateUserRoleAssignment(ctx, dbsqlc.CreateUserRoleAssignmentParams{
		ID:           params.ID,
		UserID:       params.UserID,
		RoleID:       role.ID,
		FoundationID: params.FoundationID,
		SchoolID:     params.SchoolID,
		ClassID:      params.ClassID,
		StudentID:    params.StudentID,
		EmployeeID:   params.EmployeeID,
		SubjectID:    params.SubjectID,
		ScopeJson:    params.Scope,
		StartsAt:     nullableTimestamp(params.StartsAt),
		EndsAt:       nullableTimestamp(params.EndsAt),
	})
	if err != nil {
		return domain.RoleAssignment{}, fmt.Errorf("create role assignment: %w", err)
	}

	newValues, err := json.Marshal(map[string]any{
		"user_id":       params.UserID,
		"role":          role.Code,
		"foundation_id": params.FoundationID,
		"school_id":     params.SchoolID,
		"class_id":      params.ClassID,
		"student_id":    params.StudentID,
		"employee_id":   params.EmployeeID,
		"subject_id":    params.SubjectID,
	})
	if err != nil {
		return domain.RoleAssignment{}, fmt.Errorf("marshal role assignment audit: %w", err)
	}
	if err := queries.CreateIdentityAuditLog(ctx, dbsqlc.CreateIdentityAuditLogParams{
		ID:            uuid.New(),
		FoundationID:  params.FoundationID,
		SchoolID:      params.SchoolID,
		ActorUserID:   &params.ActorUserID,
		ActorRole:     optionalText(params.ActorRole),
		Action:        "identity.role.assigned",
		Module:        "identity",
		EntityType:    "user_role_assignment",
		EntityID:      assignment.ID,
		NewValuesJson: newValues,
		IpAddress:     params.IPAddress,
		UserAgent:     textValue(params.UserAgent),
		RequestID:     params.RequestID,
		CorrelationID: params.CorrelationID,
		OccurredAt:    pgtype.Timestamptz{Time: r.now().UTC(), Valid: true},
	}); err != nil {
		return domain.RoleAssignment{}, fmt.Errorf("audit role assignment: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return domain.RoleAssignment{}, fmt.Errorf("commit role assignment: %w", err)
	}

	return mapRoleAssignment(assignment, role.Code, nil), nil
}

func (r *AuthorizationRepository) GetUserContext(ctx context.Context, userID uuid.UUID, at time.Time) (domain.UserContext, error) {
	rows, err := r.queries.GetUserContextAssignments(ctx, dbsqlc.GetUserContextAssignmentsParams{
		UserID:   userID,
		StartsAt: pgtype.Timestamptz{Time: at.UTC(), Valid: true},
	})
	if err != nil {
		return domain.UserContext{}, fmt.Errorf("get user context assignments: %w", err)
	}

	contextResult := domain.UserContext{UserID: userID, Assignments: make([]domain.RoleAssignment, 0, len(rows))}
	roles := make(map[string]struct{})
	permissions := make(map[string]struct{})
	for _, row := range rows {
		assignment := domain.RoleAssignment{
			ID:           row.AssignmentID,
			Role:         row.RoleCode,
			FoundationID: row.FoundationID,
			SchoolID:     row.SchoolID,
			ClassID:      row.ClassID,
			StudentID:    row.StudentID,
			EmployeeID:   row.EmployeeID,
			SubjectID:    row.SubjectID,
			Scope:        row.ScopeJson,
			StartsAt:     nullableTime(row.StartsAt),
			EndsAt:       nullableTime(row.EndsAt),
			Permissions:  row.PermissionCodes,
		}
		contextResult.Assignments = append(contextResult.Assignments, assignment)
		roles[row.RoleCode] = struct{}{}
		for _, permission := range row.PermissionCodes {
			permissions[permission] = struct{}{}
		}
	}
	contextResult.Roles = sortedKeys(roles)
	contextResult.Permissions = sortedKeys(permissions)
	return contextResult, nil
}

func validateRoleScope(roleCode string, params AssignRoleParams) error {
	hasNestedScope := params.ClassID != nil || params.StudentID != nil || params.EmployeeID != nil || params.SubjectID != nil
	if hasNestedScope && params.SchoolID == nil {
		return ErrInvalidRoleScope
	}
	switch roleCode {
	case "admin_yayasan":
		if params.SchoolID != nil || hasNestedScope {
			return ErrInvalidRoleScope
		}
	case "kepala_sekolah", "tu_staff", "bendahara_sekolah":
		if params.SchoolID == nil || hasNestedScope {
			return ErrInvalidRoleScope
		}
	case "guru":
		if params.SchoolID == nil || (params.ClassID == nil && params.SubjectID == nil) || params.StudentID != nil {
			return ErrInvalidRoleScope
		}
	case "orang_tua", "siswa":
		if params.SchoolID == nil || params.StudentID == nil || params.ClassID != nil || params.SubjectID != nil || params.EmployeeID != nil {
			return ErrInvalidRoleScope
		}
	default:
		return ErrInvalidRoleScope
	}
	return nil
}

func mapRoleAssignment(assignment dbsqlc.UserRoleAssignment, roleCode string, permissions []string) domain.RoleAssignment {
	return domain.RoleAssignment{
		ID: assignment.ID, Role: roleCode, FoundationID: assignment.FoundationID,
		SchoolID: assignment.SchoolID, ClassID: assignment.ClassID, StudentID: assignment.StudentID,
		EmployeeID: assignment.EmployeeID, SubjectID: assignment.SubjectID, Scope: assignment.ScopeJson,
		StartsAt: nullableTime(assignment.StartsAt), EndsAt: nullableTime(assignment.EndsAt), Permissions: permissions,
	}
}

func nullableTimestamp(value *time.Time) pgtype.Timestamptz {
	if value == nil {
		return pgtype.Timestamptz{}
	}
	return pgtype.Timestamptz{Time: value.UTC(), Valid: true}
}

func optionalText(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func sortedKeys(values map[string]struct{}) []string {
	result := make([]string, 0, len(values))
	for value := range values {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

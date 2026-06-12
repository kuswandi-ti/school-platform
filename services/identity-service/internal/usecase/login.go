package usecase

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"school-platform/services/identity-service/internal/domain"
	"school-platform/services/identity-service/internal/password"
	"school-platform/services/identity-service/internal/repository"
	"school-platform/services/identity-service/internal/token"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserInactive       = errors.New("user is not active")
)

type UserStore interface {
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateLastLogin(ctx context.Context, id uuid.UUID, lastLoginAt time.Time) (domain.User, error)
}

type SessionStore interface {
	CreateSession(ctx context.Context, params repository.CreateSessionParams) error
}

type LoginContextStore interface {
	GetUserContext(ctx context.Context, userID uuid.UUID, at time.Time) (domain.UserContext, error)
}

type TokenIssuer interface {
	Issue(userID uuid.UUID, actorClaims token.ActorClaims) (token.Tokens, error)
}

type Login struct {
	users             UserStore
	sessions          SessionStore
	contexts          LoginContextStore
	tokens            TokenIssuer
	dummyPasswordHash string
	now               func() time.Time
}

type LoginInput struct {
	Email     string
	Password  string
	IPAddress *string
	UserAgent *string
}

type LoginOutput struct {
	UserID       uuid.UUID
	DisplayName  string
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

func NewLogin(users UserStore, sessions SessionStore, contexts LoginContextStore, tokens TokenIssuer) (*Login, error) {
	dummyPasswordHash, err := password.Hash(uuid.NewString())
	if err != nil {
		return nil, fmt.Errorf("create dummy password hash: %w", err)
	}
	return &Login{
		users:             users,
		sessions:          sessions,
		contexts:          contexts,
		tokens:            tokens,
		dummyPasswordHash: dummyPasswordHash,
		now:               time.Now,
	}, nil
}

func (u *Login) Execute(ctx context.Context, input LoginInput) (LoginOutput, error) {
	user, err := u.users.FindByEmail(ctx, strings.TrimSpace(input.Email))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			_, _ = password.Verify(input.Password, u.dummyPasswordHash)
			return LoginOutput{}, ErrInvalidCredentials
		}
		return LoginOutput{}, fmt.Errorf("find login user: %w", err)
	}

	verified, err := password.Verify(input.Password, user.PasswordHash)
	if err != nil {
		return LoginOutput{}, fmt.Errorf("verify stored password hash: %w", err)
	}
	if !verified {
		return LoginOutput{}, ErrInvalidCredentials
	}
	if user.Status != "active" {
		return LoginOutput{}, ErrUserInactive
	}

	now := u.now().UTC()
	userContext, err := u.contexts.GetUserContext(ctx, user.ID, now)
	if err != nil {
		return LoginOutput{}, fmt.Errorf("load login user context: %w", err)
	}

	issued, err := u.tokens.Issue(user.ID, actorClaimsFromContext(userContext))
	if err != nil {
		return LoginOutput{}, fmt.Errorf("issue login tokens: %w", err)
	}
	if err := u.sessions.CreateSession(ctx, repository.CreateSessionParams{
		ID:               uuid.New(),
		UserID:           user.ID,
		RefreshTokenHash: issued.RefreshTokenHash,
		IPAddress:        input.IPAddress,
		UserAgent:        input.UserAgent,
		ExpiresAt:        issued.RefreshExpiresAt,
	}); err != nil {
		return LoginOutput{}, fmt.Errorf("persist login session: %w", err)
	}

	if _, err := u.users.UpdateLastLogin(ctx, user.ID, now); err != nil {
		return LoginOutput{}, fmt.Errorf("update login timestamp: %w", err)
	}

	return LoginOutput{
		UserID:       user.ID,
		DisplayName:  user.DisplayName,
		AccessToken:  issued.AccessToken,
		RefreshToken: issued.RefreshToken,
		ExpiresIn:    int64(issued.AccessExpiresAt.Sub(now).Seconds()),
	}, nil
}

func actorClaimsFromContext(userContext domain.UserContext) token.ActorClaims {
	foundationIDs := make(map[string]struct{})
	schoolIDs := make(map[string]struct{})
	classIDs := make(map[string]struct{})
	studentIDs := make(map[string]struct{})
	subjectIDs := make(map[string]struct{})
	employeeIDs := make(map[string]struct{})

	for _, assignment := range userContext.Assignments {
		foundationIDs[assignment.FoundationID.String()] = struct{}{}
		if assignment.SchoolID != nil {
			schoolIDs[assignment.SchoolID.String()] = struct{}{}
		}
		if assignment.ClassID != nil {
			classIDs[assignment.ClassID.String()] = struct{}{}
		}
		if assignment.StudentID != nil {
			studentIDs[assignment.StudentID.String()] = struct{}{}
		}
		if assignment.SubjectID != nil {
			subjectIDs[assignment.SubjectID.String()] = struct{}{}
		}
		if assignment.EmployeeID != nil {
			employeeIDs[assignment.EmployeeID.String()] = struct{}{}
		}
	}

	scope := map[string]any{
		"foundation_ids": sortedClaimKeys(foundationIDs),
		"school_ids":     sortedClaimKeys(schoolIDs),
		"class_ids":      sortedClaimKeys(classIDs),
		"student_ids":    sortedClaimKeys(studentIDs),
		"subject_ids":    sortedClaimKeys(subjectIDs),
		"employee_ids":   sortedClaimKeys(employeeIDs),
	}

	claims := token.ActorClaims{
		Roles:       append([]string(nil), userContext.Roles...),
		Permissions: append([]string(nil), userContext.Permissions...),
		Scope:       scope,
	}
	if foundationID, ok := singleClaimValue(foundationIDs); ok {
		claims.FoundationID = foundationID
	}
	if schoolID, ok := singleClaimValue(schoolIDs); ok {
		claims.SchoolID = schoolID
	}
	return claims
}

func sortedClaimKeys(values map[string]struct{}) []string {
	keys := make([]string, 0, len(values))
	for value := range values {
		keys = append(keys, value)
	}
	slices.Sort(keys)
	return keys
}

func singleClaimValue(values map[string]struct{}) (string, bool) {
	if len(values) != 1 {
		return "", false
	}
	for value := range values {
		return value, true
	}
	return "", false
}

package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"school-platform/services/school-core-service/internal/db/sqlc"
	"school-platform/services/school-core-service/internal/domain"
)

type SchoolRepository struct {
	pool    *pgxpool.Pool
	queries *dbsqlc.Queries
	now     func() time.Time
}

func NewSchoolRepository(pool *pgxpool.Pool) *SchoolRepository {
	return &SchoolRepository{pool: pool, queries: dbsqlc.New(pool), now: time.Now}
}

func (r *SchoolRepository) GetFoundation(ctx context.Context, id uuid.UUID) (domain.Foundation, error) {
	v, err := r.queries.GetFoundationByID(ctx, id)
	if err != nil {
		return domain.Foundation{}, err
	}
	return mapFoundation(v), nil
}
func (r *SchoolRepository) ListSchools(ctx context.Context, foundationID uuid.UUID) ([]domain.School, error) {
	rows, err := r.queries.ListSchoolsByFoundation(ctx, foundationID)
	if err != nil {
		return nil, err
	}
	out := make([]domain.School, 0, len(rows))
	for _, v := range rows {
		out = append(out, mapSchool(v))
	}
	return out, nil
}
func (r *SchoolRepository) GetSchool(ctx context.Context, id, foundationID uuid.UUID) (domain.School, error) {
	v, err := r.queries.GetSchoolByScope(ctx, dbsqlc.GetSchoolByScopeParams{ID: id, FoundationID: foundationID})
	if err != nil {
		return domain.School{}, err
	}
	return mapSchool(v), nil
}

type SchoolWrite struct {
	ID           uuid.UUID
	FoundationID uuid.UUID
	SchoolCode   string
	Name         string
	SchoolLevel  string
	NPSN         *string
	Address      *string
	Phone        *string
	Email        *string
	LogoFileID   *uuid.UUID
	Status       string
}
type AuditContext struct {
	Actor     domain.Actor
	OldValues any
	NewValues any
	Action    string
}

func (r *SchoolRepository) CreateSchool(ctx context.Context, p SchoolWrite, a AuditContext) (domain.School, error) {
	return r.write(ctx, p, a, true)
}
func (r *SchoolRepository) UpdateSchool(ctx context.Context, p SchoolWrite, a AuditContext) (domain.School, error) {
	return r.write(ctx, p, a, false)
}

func (r *SchoolRepository) write(ctx context.Context, p SchoolWrite, a AuditContext, create bool) (domain.School, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.School{}, fmt.Errorf("begin school write: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	q := r.queries.WithTx(tx)
	params := dbsqlc.CreateSchoolParams{ID: p.ID, FoundationID: p.FoundationID, SchoolCode: p.SchoolCode, Name: p.Name, SchoolLevel: p.SchoolLevel, Npsn: p.NPSN, Address: p.Address, Phone: p.Phone, Email: p.Email, LogoFileID: p.LogoFileID, Status: p.Status}
	var row dbsqlc.School
	if create {
		row, err = q.CreateSchool(ctx, params)
	} else {
		row, err = q.UpdateSchool(ctx, dbsqlc.UpdateSchoolParams(params))
	}
	if err != nil {
		return domain.School{}, err
	}
	oldJSON, err := marshalOptional(a.OldValues)
	if err != nil {
		return domain.School{}, err
	}
	newJSON, err := json.Marshal(a.NewValues)
	if err != nil {
		return domain.School{}, err
	}
	role := firstRole(a.Actor.Roles)
	schoolID := p.ID
	if err = q.CreateAuditLog(ctx, dbsqlc.CreateAuditLogParams{ID: uuid.New(), FoundationID: p.FoundationID, SchoolID: &schoolID, ActorUserID: &a.Actor.UserID, ActorRole: &role, Action: a.Action, Module: "school", EntityType: "school", EntityID: p.ID, OldValuesJson: oldJSON, NewValuesJson: newJSON, IpAddress: a.Actor.IPAddress, UserAgent: a.Actor.UserAgent, RequestID: a.Actor.RequestID, CorrelationID: a.Actor.CorrelationID, OccurredAt: pgtype.Timestamptz{Time: r.now().UTC(), Valid: true}}); err != nil {
		return domain.School{}, fmt.Errorf("create school audit: %w", err)
	}
	if create {
		envelope := map[string]any{"event_id": uuid.New(), "event_type": "school.school.created", "event_version": 1, "source_service": "school-core-service", "occurred_at": r.now().UTC(), "request_id": a.Actor.RequestID, "correlation_id": a.Actor.CorrelationID, "actor": map[string]any{"user_id": a.Actor.UserID, "role": role}, "tenant": map[string]any{"foundation_id": p.FoundationID, "school_id": p.ID}, "entity": map[string]any{"entity_type": "school", "entity_id": p.ID}, "payload": map[string]any{"school_id": p.ID, "school_code": p.SchoolCode, "name": p.Name, "school_level": p.SchoolLevel, "status": p.Status}, "metadata": map[string]any{"classification": "internal"}}
		payload, marshalErr := json.Marshal(envelope)
		if marshalErr != nil {
			return domain.School{}, marshalErr
		}
		eventID := envelope["event_id"].(uuid.UUID)
		if err = q.CreateOutboxEvent(ctx, dbsqlc.CreateOutboxEventParams{ID: uuid.New(), EventID: eventID, EventType: "school.school.created", EventVersion: 1, AggregateType: "school", AggregateID: p.ID, PayloadJson: payload}); err != nil {
			return domain.School{}, fmt.Errorf("create school outbox event: %w", err)
		}
	}
	if err = tx.Commit(ctx); err != nil {
		return domain.School{}, fmt.Errorf("commit school write: %w", err)
	}
	return mapSchool(row), nil
}
func marshalOptional(v any) (json.RawMessage, error) {
	if v == nil {
		return nil, nil
	}
	return json.Marshal(v)
}
func firstRole(values []string) string {
	if len(values) == 0 {
		return ""
	}
	return values[0]
}
func mapFoundation(v dbsqlc.Foundation) domain.Foundation {
	return domain.Foundation{ID: v.ID, FoundationCode: v.FoundationCode, Name: v.Name, LegalName: v.LegalName, Address: v.Address, Phone: v.Phone, Email: v.Email, LogoFileID: v.LogoFileID, Timezone: v.Timezone, Status: v.Status, CreatedAt: v.CreatedAt.Time, UpdatedAt: v.UpdatedAt.Time}
}
func mapSchool(v dbsqlc.School) domain.School {
	return domain.School{ID: v.ID, FoundationID: v.FoundationID, SchoolCode: v.SchoolCode, Name: v.Name, SchoolLevel: v.SchoolLevel, NPSN: v.Npsn, Address: v.Address, Phone: v.Phone, Email: v.Email, LogoFileID: v.LogoFileID, Status: v.Status, CreatedAt: v.CreatedAt.Time, UpdatedAt: v.UpdatedAt.Time}
}

var _ = pgx.ErrNoRows

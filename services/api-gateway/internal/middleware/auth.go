package middleware

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"school-platform/services/api-gateway/internal/response"
)

type actorContextKey string

const (
	actorContextContextKey actorContextKey = "actor_context"
	accessTokenContextKey  actorContextKey = "access_token"
)

type ActorContext struct {
	UserID        string
	FoundationID  string
	SchoolID      string
	Roles         []string
	Permissions   []string
	Scope         map[string]any
	RequestID     string
	CorrelationID string
}

type AccessTokenClaims struct {
	FoundationID string         `json:"foundation_id,omitempty"`
	SchoolID     string         `json:"school_id,omitempty"`
	Roles        []string       `json:"roles,omitempty"`
	Permissions  []string       `json:"permissions,omitempty"`
	Scope        map[string]any `json:"scope,omitempty"`
	jwt.RegisteredClaims
}

type JWTValidator struct {
	publicKey ed25519.PublicKey
	issuer    string
	audience  string
}

func NewJWTValidator(publicKey ed25519.PublicKey, issuer, audience string) (*JWTValidator, error) {
	if len(publicKey) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("invalid Ed25519 public key")
	}
	if issuer == "" || audience == "" {
		return nil, fmt.Errorf("invalid JWT validator configuration")
	}
	return &JWTValidator{publicKey: publicKey, issuer: issuer, audience: audience}, nil
}

func LoadPublicKey(path string) (ed25519.PublicKey, error) {
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read JWT public key: %w", err)
	}
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("decode JWT public key PEM")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse JWT public key: %w", err)
	}
	edPublicKey, ok := publicKey.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("JWT public key must be Ed25519")
	}
	return edPublicKey, nil
}

func (v *JWTValidator) ValidateAccessToken(accessToken string) (ActorContext, error) {
	claims := &AccessTokenClaims{}
	parsed, err := jwt.ParseWithClaims(accessToken, claims, func(parsed *jwt.Token) (any, error) {
		if parsed.Method != jwt.SigningMethodEdDSA {
			return nil, fmt.Errorf("unexpected access token signing method")
		}
		return v.publicKey, nil
	}, jwt.WithIssuer(v.issuer), jwt.WithAudience(v.audience), jwt.WithExpirationRequired())
	if err != nil {
		return ActorContext{}, fmt.Errorf("validate access token: %w", err)
	}
	if !parsed.Valid {
		return ActorContext{}, fmt.Errorf("validate access token: token is invalid")
	}
	if strings.TrimSpace(claims.Subject) == "" {
		return ActorContext{}, fmt.Errorf("access token subject is required")
	}

	return ActorContext{
		UserID:       claims.Subject,
		FoundationID: claims.FoundationID,
		SchoolID:     claims.SchoolID,
		Roles:        append([]string(nil), claims.Roles...),
		Permissions:  append([]string(nil), claims.Permissions...),
		Scope:        cloneScope(claims.Scope),
	}, nil
}

func RequireAuth(validator *JWTValidator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken, ok := bearerToken(r.Header.Get("Authorization"))
			if !ok {
				response.Error(w, http.StatusUnauthorized, response.CodeUnauthorized, "Access token wajib diisi.", nil)
				return
			}

			actorContext, err := validator.ValidateAccessToken(accessToken)
			if err != nil {
				response.Error(w, http.StatusUnauthorized, response.CodeUnauthorized, "Access token tidak valid.", nil)
				return
			}
			actorContext.RequestID = RequestIDFromContext(r.Context())
			actorContext.CorrelationID = CorrelationIDFromContext(r.Context())

			ctx := context.WithValue(r.Context(), actorContextContextKey, actorContext)
			ctx = context.WithValue(ctx, accessTokenContextKey, accessToken)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ActorContextFromContext(ctx context.Context) (ActorContext, bool) {
	value, ok := ctx.Value(actorContextContextKey).(ActorContext)
	return value, ok
}

// WithActorContext attaches a validated actor to an internal request context.
func WithActorContext(ctx context.Context, actor ActorContext) context.Context {
	return context.WithValue(ctx, actorContextContextKey, actor)
}

func AccessTokenFromContext(ctx context.Context) (string, bool) {
	value, ok := ctx.Value(accessTokenContextKey).(string)
	return value, ok
}

func bearerToken(header string) (string, bool) {
	parts := strings.Fields(header)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || parts[1] == "" {
		return "", false
	}
	return parts[1], true
}

func cloneScope(scope map[string]any) map[string]any {
	if len(scope) == 0 {
		return map[string]any{}
	}
	cloned := make(map[string]any, len(scope))
	for key, value := range scope {
		cloned[key] = value
	}
	return cloned
}

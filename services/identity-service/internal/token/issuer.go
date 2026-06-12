package token

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Tokens struct {
	AccessToken      string
	RefreshToken     string
	RefreshTokenHash string
	AccessExpiresAt  time.Time
	RefreshExpiresAt time.Time
}

type ActorClaims struct {
	FoundationID string         `json:"foundation_id,omitempty"`
	SchoolID     string         `json:"school_id,omitempty"`
	Roles        []string       `json:"roles,omitempty"`
	Permissions  []string       `json:"permissions,omitempty"`
	Scope        map[string]any `json:"scope,omitempty"`
}

type AccessTokenClaims struct {
	ActorClaims
	jwt.RegisteredClaims
}

type Issuer struct {
	privateKey ed25519.PrivateKey
	issuer     string
	audience   string
	accessTTL  time.Duration
	refreshTTL time.Duration
	now        func() time.Time
}

func NewIssuer(privateKey ed25519.PrivateKey, issuer, audience string, accessTTL, refreshTTL time.Duration) (*Issuer, error) {
	if len(privateKey) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("invalid Ed25519 private key")
	}
	if issuer == "" || audience == "" || accessTTL <= 0 || refreshTTL <= 0 {
		return nil, fmt.Errorf("invalid token issuer configuration")
	}
	return &Issuer{
		privateKey: privateKey,
		issuer:     issuer,
		audience:   audience,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		now:        time.Now,
	}, nil
}

func (i *Issuer) Issue(userID uuid.UUID, actorClaims ActorClaims) (Tokens, error) {
	now := i.now().UTC()
	accessExpiresAt := now.Add(i.accessTTL)
	claims := AccessTokenClaims{
		ActorClaims: actorClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    i.issuer,
			Subject:   userID.String(),
			Audience:  jwt.ClaimStrings{i.audience},
			ExpiresAt: jwt.NewNumericDate(accessExpiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ID:        uuid.NewString(),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims).SignedString(i.privateKey)
	if err != nil {
		return Tokens{}, fmt.Errorf("sign access token: %w", err)
	}

	refreshBytes := make([]byte, 32)
	if _, err := rand.Read(refreshBytes); err != nil {
		return Tokens{}, fmt.Errorf("generate refresh token: %w", err)
	}
	refreshToken := base64.RawURLEncoding.EncodeToString(refreshBytes)

	return Tokens{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshTokenHash: HashRefreshToken(refreshToken),
		AccessExpiresAt:  accessExpiresAt,
		RefreshExpiresAt: now.Add(i.refreshTTL),
	}, nil
}

func (i *Issuer) ValidateAccessToken(accessToken string) (uuid.UUID, error) {
	claims := &AccessTokenClaims{}
	parsed, err := jwt.ParseWithClaims(accessToken, claims, func(parsed *jwt.Token) (any, error) {
		if parsed.Method != jwt.SigningMethodEdDSA {
			return nil, fmt.Errorf("unexpected access token signing method")
		}
		return i.privateKey.Public(), nil
	}, jwt.WithIssuer(i.issuer), jwt.WithAudience(i.audience), jwt.WithExpirationRequired(), jwt.WithTimeFunc(i.now))
	if err != nil {
		return uuid.Nil, fmt.Errorf("validate access token: %w", err)
	}
	if !parsed.Valid {
		return uuid.Nil, fmt.Errorf("validate access token: token is invalid")
	}
	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, fmt.Errorf("parse access token subject: %w", err)
	}
	return userID, nil
}

func HashRefreshToken(refreshToken string) string {
	refreshHash := sha256.Sum256([]byte(refreshToken))
	return hex.EncodeToString(refreshHash[:])
}

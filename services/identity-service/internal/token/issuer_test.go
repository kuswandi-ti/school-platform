package token

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestIssuerCreatesVerifiableAccessAndOpaqueRefreshTokens(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	issuer, err := NewIssuer(privateKey, "identity-test", "school-platform-test", 15*time.Minute, 30*24*time.Hour)
	require.NoError(t, err)

	userID := uuid.New()
	tokens, err := issuer.Issue(userID)
	require.NoError(t, err)
	require.NotEmpty(t, tokens.AccessToken)
	require.NotEmpty(t, tokens.RefreshToken)
	require.NotEqual(t, tokens.RefreshToken, tokens.RefreshTokenHash)
	require.Len(t, tokens.RefreshTokenHash, 64)

	parsed, err := jwt.ParseWithClaims(tokens.AccessToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		require.Equal(t, jwt.SigningMethodEdDSA, token.Method)
		return publicKey, nil
	}, jwt.WithIssuer("identity-test"), jwt.WithAudience("school-platform-test"))
	require.NoError(t, err)
	require.True(t, parsed.Valid)
	claims, ok := parsed.Claims.(*jwt.RegisteredClaims)
	require.True(t, ok)
	require.Equal(t, userID.String(), claims.Subject)
	validatedUserID, err := issuer.ValidateAccessToken(tokens.AccessToken)
	require.NoError(t, err)
	require.Equal(t, userID, validatedUserID)
}

func TestIssuerRejectsInvalidAccessToken(t *testing.T) {
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	issuer, err := NewIssuer(privateKey, "identity-test", "school-platform-test", time.Minute, time.Hour)
	require.NoError(t, err)

	_, err = issuer.ValidateAccessToken("not-a-jwt")
	require.Error(t, err)
}

func TestIssuerUsesUniqueRefreshTokens(t *testing.T) {
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	issuer, err := NewIssuer(privateKey, "identity-test", "school-platform-test", time.Minute, time.Hour)
	require.NoError(t, err)

	first, err := issuer.Issue(uuid.New())
	require.NoError(t, err)
	second, err := issuer.Issue(uuid.New())
	require.NoError(t, err)
	require.NotEqual(t, first.RefreshToken, second.RefreshToken)
	require.NotEqual(t, first.RefreshTokenHash, second.RefreshTokenHash)
}

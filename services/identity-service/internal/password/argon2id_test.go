package password

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashAndVerify(t *testing.T) {
	hash, err := Hash("correct horse battery staple")
	require.NoError(t, err)
	require.NotContains(t, hash, "correct horse battery staple")
	require.True(t, strings.HasPrefix(hash, "$argon2id$v=19$"))

	verified, err := Verify("correct horse battery staple", hash)
	require.NoError(t, err)
	require.True(t, verified)
}

func TestVerifyRejectsWrongPassword(t *testing.T) {
	hash, err := Hash("correct password")
	require.NoError(t, err)

	verified, err := Verify("wrong password", hash)
	require.NoError(t, err)
	require.False(t, verified)
}

func TestHashUsesUniqueSalt(t *testing.T) {
	first, err := Hash("same password")
	require.NoError(t, err)
	second, err := Hash("same password")
	require.NoError(t, err)
	require.NotEqual(t, first, second)
}

func TestVerifyRejectsMalformedHash(t *testing.T) {
	tests := []string{
		"not-a-password-hash",
		"$argon2id$v=19junk$m=65536,t=3,p=2$c2FsdHNhbHRzYWx0c2FsdA$aGFzaGhhc2hoYXNoaGFzaGhhc2g",
		"$argon2id$v=19$m=65536,t=3,p=2junk$c2FsdHNhbHRzYWx0c2FsdA$aGFzaGhhc2hoYXNoaGFzaGhhc2g",
	}

	for _, encodedHash := range tests {
		verified, err := Verify("password", encodedHash)
		require.False(t, verified)
		require.ErrorIs(t, err, ErrInvalidHash)
	}
}

func TestPasswordLengthLimit(t *testing.T) {
	tooLong := strings.Repeat("a", maxPasswordBytes+1)
	_, err := Hash(tooLong)
	require.True(t, errors.Is(err, ErrPasswordTooLong))
}

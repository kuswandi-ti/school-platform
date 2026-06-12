package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	argon2Version    = 19
	defaultMemory    = 64 * 1024
	defaultTime      = 3
	defaultThreads   = 2
	defaultSaltBytes = 16
	defaultKeyBytes  = 32
	maxPasswordBytes = 1024
)

var (
	ErrInvalidHash     = errors.New("invalid password hash")
	ErrPasswordTooLong = errors.New("password exceeds maximum length")
)

type argon2Params struct {
	memory  uint32
	time    uint32
	threads uint8
}

func Hash(plainPassword string) (string, error) {
	if len(plainPassword) > maxPasswordBytes {
		return "", ErrPasswordTooLong
	}

	salt := make([]byte, defaultSaltBytes)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate password salt: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(plainPassword),
		salt,
		defaultTime,
		defaultMemory,
		defaultThreads,
		defaultKeyBytes,
	)
	encoding := base64.RawStdEncoding

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2Version,
		defaultMemory,
		defaultTime,
		defaultThreads,
		encoding.EncodeToString(salt),
		encoding.EncodeToString(hash),
	), nil
}

func Verify(plainPassword, encodedHash string) (bool, error) {
	if len(plainPassword) > maxPasswordBytes {
		return false, ErrPasswordTooLong
	}

	params, salt, expectedHash, err := parse(encodedHash)
	if err != nil {
		return false, err
	}

	actualHash := argon2.IDKey(
		[]byte(plainPassword),
		salt,
		params.time,
		params.memory,
		params.threads,
		uint32(len(expectedHash)),
	)

	return subtle.ConstantTimeCompare(actualHash, expectedHash) == 1, nil
}

func parse(encodedHash string) (argon2Params, []byte, []byte, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return argon2Params{}, nil, nil, ErrInvalidHash
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil ||
		version != argon2Version || parts[2] != fmt.Sprintf("v=%d", version) {
		return argon2Params{}, nil, nil, ErrInvalidHash
	}

	var params argon2Params
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.memory, &params.time, &params.threads); err != nil {
		return argon2Params{}, nil, nil, ErrInvalidHash
	}
	if parts[3] != fmt.Sprintf("m=%d,t=%d,p=%d", params.memory, params.time, params.threads) {
		return argon2Params{}, nil, nil, ErrInvalidHash
	}
	if params.memory < 8*1024 || params.memory > 256*1024 ||
		params.time < 1 || params.time > 10 ||
		params.threads < 1 || params.threads > 16 {
		return argon2Params{}, nil, nil, ErrInvalidHash
	}

	encoding := base64.RawStdEncoding
	salt, err := encoding.DecodeString(parts[4])
	if err != nil || len(salt) < 16 || len(salt) > 64 {
		return argon2Params{}, nil, nil, ErrInvalidHash
	}

	hash, err := encoding.DecodeString(parts[5])
	if err != nil || len(hash) < 16 || len(hash) > 64 {
		return argon2Params{}, nil, nil, ErrInvalidHash
	}

	return params, salt, hash, nil
}

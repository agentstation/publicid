package publicid

import (
	"fmt"

	nanoid "github.com/matoous/go-nanoid/v2"
)

const (
	alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	longLen  = 12
	shortLen = 8
)

// Option is a function type for configuring ID generation.
type Option func(*config)

// config holds the configuration for ID generation.
type config struct {
	attempts int
}

// Attempts returns an Option to set the number of attempts for ID generation.
func Attempts(n int) Option {
	return func(c *config) {
		c.attempts = n
	}
}

// New generates a unique nanoID with a length of 8 characters and the given options.
func New(opts ...Option) (string, error) {
	return generateID(shortLen, opts...)
}

// NewLong generates a unique nanoID with a length of 12 characters and the given options.
func NewLong(opts ...Option) (string, error) {
	return generateID(longLen, opts...)
}

// generateID is a helper function to generate IDs with the given length and options.
func generateID(length int, opts ...Option) (string, error) {
	cfg := &config{attempts: 1}
	for _, opt := range opts {
		opt(cfg)
	}

	var lastErr error
	for i := 0; i < cfg.attempts; i++ {
		id, err := nanoid.Generate(alphabet, length)
		if err == nil {
			return id, nil
		}
		lastErr = err
	}
	return "", fmt.Errorf("failed to generate ID after %d attempts: %w", cfg.attempts, lastErr)
}

// Validate checks if a given field name's public ID value is valid according to
// the constraints defined by package publicid.
func Validate(id string) error {
	return validate(id, shortLen)
}

// validateLong checks if a given field name's public ID value is valid according to
// the constraints defined by package publicid.
func ValidateLong(fieldName, id string) error {
	return validate(id, longLen)
}

// validate checks if a given public ID value is valid.
func validate(id string, expectedLen int) error {
	if id == "" {
		return fmt.Errorf("public ID is empty")
	}
	if len(id) != expectedLen {
		return fmt.Errorf("public ID has length %d, want %d", len(id), expectedLen)
	}
	for _, char := range id {
		if !isValidChar(char) {
			return fmt.Errorf("public ID contains invalid character: %c", char)
		}
	}
	return nil
}

// isValidChar checks if a given character is a valid public ID character.
func isValidChar(c rune) bool {
	return (c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}

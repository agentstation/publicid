// Package publicid generates and validates NanoID strings designed to be publicly exposed.
// It uses the nanoid library to generate IDs and provides options to configure the generation process.
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

// generator is the function used to generate nanoIDs.
var generator = nanoid.Generate

// Option is a function type for configuring ID generation.
type Option func(*config)

// config holds the configuration for ID generation.
type config struct {
	attempts int
	length   int
	alphabet string
}

// Attempts returns an Option to set the number of attempts for ID generation.
func Attempts(n int) Option {
	return func(c *config) {
		c.attempts = n
	}
}

// Len returns an Option to set the length of the ID to be generated.
func Len(n int) Option {
	return func(c *config) {
		c.length = n
	}
}

// Alphabet returns an Option to set the alphabet to be used for ID generation.
func Alphabet(a string) Option {
	return func(c *config) {
		c.alphabet = a
	}
}

// New generates a unique nanoID with a length of 8 characters and the given options.
func New(opts ...Option) (string, error) { return generateID(shortLen, opts...) }

// NewLong generates a unique nanoID with a length of 12 characters and the given options.
func NewLong(opts ...Option) (string, error) { return generateID(longLen, opts...) }

// generateID is a helper function to generate IDs with the given length and options.
func generateID(length int, opts ...Option) (string, error) {
	// set default configuration values
	cfg := &config{attempts: 1, length: length, alphabet: alphabet}
	for _, opt := range opts {
		opt(cfg)
	}
	// try to generate the ID
	var lastErr error
	for i := 0; i < cfg.attempts; i++ {
		id, err := generator(cfg.alphabet, cfg.length)
		if err == nil {
			return id, nil
		}
		lastErr = err
	}
	// if we get here, we failed to generate an ID
	return "", fmt.Errorf("failed to generate ID after %d attempts: %w", cfg.attempts, lastErr)
}

// Validate checks if a given field name's public ID value is valid according to
// the constraints defined by package publicid.
func Validate(id string) error { return validate(id, shortLen) }

// validateLong checks if a given field name's public ID value is valid according to
// the constraints defined by package publicid.
func ValidateLong(id string) error { return validate(id, longLen) }

// validate checks if a given public ID value is valid.
func validate(id string, expectedLen int) error {
	if id == "" { // if the ID is empty, it's not valid
		return fmt.Errorf("public ID is empty")
	}
	if len(id) != expectedLen { // if the ID is not the expected length, it's not valid
		return fmt.Errorf("public ID has length %d, want %d", len(id), expectedLen)
	}
	for _, char := range id {
		if !isValidChar(char) { // if the ID contains an invalid character, it's not valid
			return fmt.Errorf("public ID contains invalid character: %c", char)
		}
	}
	return nil // if we get here, the ID is valid
}

// isValidChar checks if a given character is a valid public ID character.
func isValidChar(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

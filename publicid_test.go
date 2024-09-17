//go:build !integration
// +build !integration

package publicid

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	id, err := New()
	if err != nil {
		t.Errorf("New() returned an error: %v", err)
	}
	if len(id) != 8 {
		t.Errorf("New() returned id with length %d, want 8", len(id))
	}
	if err := Validate(id); err != nil {
		t.Errorf("New() returned invalid id: %v", err)
	}
}

func TestNewWithAttempts(t *testing.T) {
	id, err := New(Attempts(5))
	if err != nil {
		t.Errorf("New(Attempts(5)) returned an error: %v", err)
	}
	if len(id) != 8 {
		t.Errorf("New(Attempts(5)) returned id with length %d, want 8", len(id))
	}
	if err := Validate(id); err != nil {
		t.Errorf("New(Attempts(5)) returned invalid id: %v", err)
	}
}

func TestLong(t *testing.T) {
	id, err := New(Long())
	if err != nil {
		t.Errorf("New(Long()) returned an error: %v", err)
	}
	if len(id) != LongIDLength {
		t.Errorf("New(Long()) returned id with length %d, want %d", len(id), LongIDLength)
	}
	if err := Validate(id, Long()); err != nil {
		t.Errorf("New(Long()) returned invalid id: %v", err)
	}
}

func TestLongWithAttempts(t *testing.T) {
	id, err := New(Long(), Attempts(5))
	if err != nil {
		t.Errorf("New(Long(), Attempts(5)) returned an error: %v", err)
	}
	if len(id) != LongIDLength {
		t.Errorf("New(Long(), Attempts(5)) returned id with length %d, want %d", len(id), LongIDLength)
	}
	if err := Validate(id, Long()); err != nil {
		t.Errorf("New(Long(), Attempts(5)) returned invalid id: %v", err)
	}
}

func TestNewWithCustomLength(t *testing.T) {
	customLength := 10
	id, err := New(Len(customLength))
	if err != nil {
		t.Errorf("New(Len(%d)) returned an error: %v", customLength, err)
	}
	if len(id) != customLength {
		t.Errorf("New(Len(%d)) returned id with length %d, want %d", customLength, len(id), customLength)
	}
	if err := Validate(id, Len(customLength)); err != nil {
		t.Errorf("New(Len(%d)) returned invalid id: %v", customLength, err)
	}
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		name      string
		id        string
		opts      []Option
		wantError bool
	}{
		{"Valid Default ID", "abCD1234", nil, false},
		{"Valid Long ID", "abCD1234EFGH", []Option{Long()}, false},
		{"Valid Custom Length ID", "abCD123456", []Option{Len(10)}, false},
		{"Empty ID", "", nil, true},
		{"Short ID", "abc123", nil, true},
		{"Long ID", "abcDEF123456", nil, true},
		{"Invalid char", "abCD12_4", nil, true},
		{"Invalid Long ID", "abCD1234EF", []Option{Long()}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Validate(tc.id, tc.opts...)
			if (err != nil) != tc.wantError {
				t.Errorf("Validate() error = %v, wantError %v", err, tc.wantError)
			}
		})
	}
}

func TestNewFailsAfterAttempts(t *testing.T) {

	// generator functions for testing
	nanoIDGenerator := generator
	mockGenerator := func(alphabet string, size int) (string, error) {
		return "", fmt.Errorf("mocked error")
	}

	// Replace the Generate function with our mock generator function,
	// and restore the original generator function in the deferred function.
	generator = mockGenerator
	defer func() { generator = nanoIDGenerator }()

	_, err := New(Attempts(3))
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else {
		expectedErrMsg := "failed to generate ID after 3 attempts: mocked error"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err.Error())
		}
	}
}

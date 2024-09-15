//go:build !integration
// +build !integration

package publicid

import (
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
	id, err := NewLong()
	if err != nil {
		t.Errorf("Long() returned an error: %v", err)
	}
	if len(id) != 12 {
		t.Errorf("Long() returned id with length %d, want 12", len(id))
	}
	if err := ValidateLong(id); err != nil {
		t.Errorf("Long() returned invalid id: %v", err)
	}
}

func TestLongWithAttempts(t *testing.T) {
	id, err := NewLong(Attempts(5))
	if err != nil {
		t.Errorf("Long(Attempts(5)) returned an error: %v", err)
	}
	if len(id) != 12 {
		t.Errorf("Long(Attempts(5)) returned id with length %d, want 12", len(id))
	}
	if err := ValidateLong(id); err != nil {
		t.Errorf("Long(Attempts(5)) returned invalid id: %v", err)
	}
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		name      string
		id        string
		wantError bool
	}{
		{"Valid ID", "abCD1234", false},
		{"Empty ID", "", true},
		{"Short ID", "abc123", true},
		{"Long ID", "abcDEF123456", true},
		{"Invalid char", "abCD12_4", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Validate(tc.id)
			if (err != nil) != tc.wantError {
				t.Errorf("Validate() error = %v, wantError %v", err, tc.wantError)
			}
		})
	}
}

func TestValidateLong(t *testing.T) {
	testCases := []struct {
		name      string
		id        string
		wantError bool
	}{
		{"Valid ID", "abCD1234EFGH", false},
		{"Empty ID", "", true},
		{"Short ID", "abcDEF123", true},
		{"Long ID", "abcDEF123456789", true},
		{"Invalid char", "abCD1234EF_H", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateLong(tc.id)
			if (err != nil) != tc.wantError {
				t.Errorf("ValidateLong() error = %v, wantError %v", err, tc.wantError)
			}
		})
	}
}

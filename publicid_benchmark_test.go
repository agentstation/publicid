package publicid_test

import (
	"testing"

	"github.com/agentstation/publicid"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := publicid.New()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewWithAttempts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := publicid.New(publicid.Attempts(5))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := publicid.NewLong()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLongWithAttempts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := publicid.NewLong(publicid.Attempts(5))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValidate(b *testing.B) {
	id, err := publicid.New()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := publicid.Validate(id)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValidateLong(b *testing.B) {
	id, err := publicid.NewLong()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := publicid.ValidateLong("BenchmarkValidateLong", id)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUniqueness(b *testing.B) {
	ids := make(map[string]bool)
	for i := 0; i < b.N; i++ {
		id, err := publicid.New()
		if err != nil {
			b.Fatalf("Failed to generate ID: %v", err)
		}
		if ids[id] {
			b.Fatalf("Duplicate ID generated: %s", id)
		}
		ids[id] = true
	}
}

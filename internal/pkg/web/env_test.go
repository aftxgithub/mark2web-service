package web

import (
	"os"
	"testing"
)

func setenv(key, value string) func() {
	os.Setenv(key, value)
	return func() {
		os.Unsetenv(key)
	}
}

func TestM2WPort(t *testing.T) {
	testPort := "3600"
	undo := setenv(M2W_PORT, testPort)
	defer undo()

	got := getPortFromEnv()
	if got != testPort {
		t.Fatalf("Expected port '%s', got '%s'", testPort, got)
	}
}

func TestFallbackPort(t *testing.T) {
	testPort := "3500"
	undo := setenv("PORT", testPort)
	defer undo()

	got := getPortFromEnv()
	if got != testPort {
		t.Fatalf("Expected port '%s', got '%s'", testPort, got)
	}
}

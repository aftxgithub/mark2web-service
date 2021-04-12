package web

import (
	"net"
	"testing"
)

func TestGetLastPath(t *testing.T) {
	testCases := []struct {
		URL      string
		expected string
	}{
		{"http://example.com/12345", "12345"},
		{"lorem.ipsum/loremipsum", "loremipsum"},
		{"example.test/test/finalpath", "finalpath"},
		{"url.nopath", ""},
		{"http://url.nopath", ""},
	}

	for _, tc := range testCases {
		got := getLastPath(tc.URL)
		if got != tc.expected {
			t.Errorf("expected '%s' from getLastPath, got '%s'", tc.expected, got)
		}
	}
}

func TestGetRunAddrWithPort(t *testing.T) {
	testPort := "3060"
	undo := setenv(M2W_PORT, testPort)
	defer undo()

	expected := net.JoinHostPort("", testPort)
	got := getRunAddr()
	if got != expected {
		t.Errorf("Expected %s from getRunAddr, got %s", expected, got)
	}
}

func TestGetRunAddrNoPort(t *testing.T) {
	expected := net.JoinHostPort("", "8080")
	got := getRunAddr()

	if got != expected {
		t.Errorf("Expected %s from getRunAddr, got %s", expected, got)
	}
}

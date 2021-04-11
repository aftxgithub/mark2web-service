package web

import (
	"net"
	"testing"
)

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

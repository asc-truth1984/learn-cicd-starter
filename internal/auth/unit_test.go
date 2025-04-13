package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testHeader := http.Header{}
	_, err := GetAPIKey(testHeader)
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("Expected error: %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}

package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey_Success(t *testing.T) {
    headers := http.Header{}
    headers.Set("Authorization", "ApiKey 12345")

    got, err := GetAPIKey(headers)
    want := "12345"

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if got != want {
        t.Errorf("GetAPIKey() = %v, want %v", got, want)
    }
}

func TestGetAPIKey_NoHeader(t *testing.T) {
    headers := http.Header{} // boş header

    _, err := GetAPIKey(headers)
    if err == nil {
        t.Fatalf("expected error but got nil")
    }
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
    headers := http.Header{}
    headers.Set("Authorization", "Bearer token") // yanlış format

    _, err := GetAPIKey(headers)
    if err == nil {
        t.Fatalf("expected error but got nil")
    }
}


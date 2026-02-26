package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matt-riley/go-api-template/internal/handler"
)

func TestHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
	rec := httptest.NewRecorder()

	handler.Hello(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var body map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body["message"] != "Hello, World!" {
		t.Fatalf("expected Hello, World!, got %q", body["message"])
	}
}

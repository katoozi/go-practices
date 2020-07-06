package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/check-status-ok", nil)
	if err != nil {
		t.Error("Error while creating request: ", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checkStatusOK)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler return %v", status)
	}

	expect := `Fine!`
	if rr.Body.String() != expect {
		t.Errorf("handler return: %v", rr.Body.String())
	}
}

func TestCheckStatusNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/check-status-not-found", nil)
	if err != nil {
		t.Error("Error while creating request: ", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statusNotFound)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusNotFound {
		t.Errorf("handler return %v", status)
	}
}

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetBudgetHandler(t *testing.T) {

	makeRequest := func(t *testing.T, body string, rr *httptest.ResponseRecorder) {

		t.Helper()

		req, err := http.NewRequest(http.MethodPost, "/calculate", strings.NewReader(body))

		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		handler := http.HandlerFunc(GetBudget)

		handler.ServeHTTP(rr, req)
	}

	t.Run("Send body and get a response", func(t *testing.T) {
		body := fmt.Sprintf(`
		{
			"Start" : "2019-01-01T00:00:00Z",
			"End" : "2019-01-02T00:00:00Z"
		}`)

		rr := httptest.NewRecorder()

		makeRequest(t, body, rr)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `"£0.03"`

		got := strings.TrimSuffix(rr.Body.String(), "\n")

		if got != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("Send body and get a response but needing 2 decimal places", func(t *testing.T) {
		body := fmt.Sprintf(`
		{
			"Start" : "2019-01-01T00:00:00Z",
			"End" : "2019-01-04T00:00:00Z"
		}`)

		rr := httptest.NewRecorder()

		makeRequest(t, body, rr)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `"£0.10"`

		got := strings.TrimSuffix(rr.Body.String(), "\n")

		if got != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("Send body but start date after end date and returns an error", func(t *testing.T) {
		body := fmt.Sprintf(`
		{
			"Start" : "2020-01-01T00:00:00Z",
			"End" : "2019-01-02T00:00:00Z"
		}`)

		rr := httptest.NewRecorder()

		makeRequest(t, body, rr)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("Send body but start year different to end year and returns an error", func(t *testing.T) {
		body := fmt.Sprintf(`
		{
			"Start" : "2020-01-01T00:00:00Z",
			"End" : "2019-01-02T00:00:00Z"
		}`)

		rr := httptest.NewRecorder()

		makeRequest(t, body, rr)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("No body, returns 400 status code", func(t *testing.T) {

		rr := httptest.NewRecorder()

		makeRequest(t, "", rr)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("Body provided but not correct, returns 400 status code", func(t *testing.T) {

		body := fmt.Sprintf(`
		{
			"Start" : "",
			"End" : ""
		}`)

		rr := httptest.NewRecorder()

		makeRequest(t, body, rr)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}

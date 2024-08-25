package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetForecast(t *testing.T) {
	var testSvrUrl string
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "points/39.7456,-97.0892") {
			w.Write([]byte(fmt.Sprintf(`{"properties":{"forecast":"%s/forecast"}}`, testSvrUrl)))
		} else if strings.Contains(r.URL.Path, "forecast") {
			w.Write([]byte(`{"properties":{"periods":[{"temperature": 102,"shortForecast":"Sunny"}]}}`))
		} else {
			t.Errorf("unexpected request path: %s", r.URL.Path)
		}
	}))
	testSvrUrl = svr.URL
	defer svr.Close()

	rh := RequestHandler{
		NoaaHost: testSvrUrl,
	}

	req := httptest.NewRequest(http.MethodGet, "/forecast/39.7456,-97.0892", nil)
	req.SetPathValue("coordinates", "39.7456,-97.0892")
	w := httptest.NewRecorder()
	rh.GetForecast(w, req)

	var result Response
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Error("Unmarshal error", err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Code)
	}

	if result.Characterization != "hot" {
		t.Errorf("Expected 'hot', got %s", result.Characterization)
	}

	if result.ShortForecast != "Sunny" {
		t.Errorf("Expected 'Sunny', got %s", result.ShortForecast)
	}
}

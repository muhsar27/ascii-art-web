package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHome(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "Valid GET /",
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Method POST /",
			method:         http.MethodPost,
			path:           "/",
			expectedStatus: http.StatusNotFound, // home calls http.NotFound
		},
		{
			name:           "Invalid path GET /invalid",
			method:         http.MethodGet,
			path:           "/invalid",
			expectedStatus: http.StatusNotFound, // home calls http.Error with StatusNotFound
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			home(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d; got %d", tt.expectedStatus, res.StatusCode)
			}
		})
	}
}

func TestAsciiArt(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		word           string
		banner         string
		expectedStatus int
		expectBody     bool
	}{
		{
			name:           "Valid POST /asciiArt standard",
			method:         http.MethodPost,
			path:           "/asciiArt",
			word:           "hello",
			banner:         "standard",
			expectedStatus: http.StatusOK,
			expectBody:     true,
		},
		{
			name:           "Valid POST /asciiArt shadow",
			method:         http.MethodPost,
			path:           "/asciiArt",
			word:           "test",
			banner:         "shadow",
			expectedStatus: http.StatusOK,
			expectBody:     true,
		},
		{
			name:           "Valid POST /asciiArt thinkertoy",
			method:         http.MethodPost,
			path:           "/asciiArt",
			word:           "go",
			banner:         "thinkertoy",
			expectedStatus: http.StatusOK,
			expectBody:     true,
		},
		{
			name:           "Empty word",
			method:         http.MethodPost,
			path:           "/asciiArt",
			word:           "",
			banner:         "standard",
			expectedStatus: http.StatusBadRequest,
			expectBody:     false,
		},
		{
			name:           "Invalid banner style (file not found)",
			method:         http.MethodPost,
			path:           "/asciiArt",
			word:           "hello",
			banner:         "nonexistent",
			expectedStatus: http.StatusOK, // main.go returns without status code setting (defaults to 200 OK)
			expectBody:     false,
		},
		{
			name:           "Invalid method GET",
			method:         http.MethodGet,
			path:           "/asciiArt",
			word:           "hello",
			banner:         "standard",
			expectedStatus: http.StatusNotFound,
			expectBody:     false,
		},
		{
			name:           "Invalid path /asciiArt/extra",
			method:         http.MethodPost,
			path:           "/asciiArt/extra",
			word:           "hello",
			banner:         "standard",
			expectedStatus: http.StatusNotFound,
			expectBody:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			var err error
			if tt.method == http.MethodGet {
				query := url.Values{}
				query.Add("word", tt.word)
				query.Add("banner", tt.banner)
				req, err = http.NewRequest(tt.method, tt.path+"?"+query.Encode(), nil)
			} else {
				form := url.Values{}
				form.Add("word", tt.word)
				form.Add("banner", tt.banner)
				req, err = http.NewRequest(tt.method, tt.path, strings.NewReader(form.Encode()))
				if err == nil {
					req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				}
			}
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rec := httptest.NewRecorder()
			asciiArt(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d; got %d", tt.expectedStatus, res.StatusCode)
			}

			if tt.expectBody {
				bodyBytes, err := io.ReadAll(res.Body)
				if err != nil {
					t.Fatalf("failed to read response body: %v", err)
				}
				bodyStr := string(bodyBytes)
				if !strings.Contains(bodyStr, "ASCII ART") {
					t.Errorf("expected body to contain 'ASCII ART', got: %s", bodyStr)
				}
			}
		})
	}
}

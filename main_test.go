package main

import (
	"ascii-art-web/functions"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHome(t *testing.T) {
	t.Run("Valid request GET /", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		home(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status %d; got %d", http.StatusOK, res.StatusCode)
		}
	})

	t.Run("Invalid path GET /invalid-path", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/invalid-path", nil)
		rec := httptest.NewRecorder()
		home(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Errorf("expected status %d; got %d", http.StatusNotFound, res.StatusCode)
		}
	})
}

func TestAsciiArt(t *testing.T) {
	t.Run("Valid POST /asciiArt", func(t *testing.T) {
		form := url.Values{}
		form.Add("word", "hello")
		form.Add("banner", "standard")
		req := httptest.NewRequest("POST", "/asciiArt", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		
		rec := httptest.NewRecorder()
		asciiArt(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status %d; got %d", http.StatusOK, res.StatusCode)
		}

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		if !strings.Contains(string(bodyBytes), "ASCII ART") {
			t.Errorf("expected body to contain 'ASCII ART'")
		}
	})

	t.Run("Empty word POST /asciiArt", func(t *testing.T) {
		form := url.Values{}
		form.Add("word", "")
		form.Add("banner", "standard")
		req := httptest.NewRequest("POST", "/asciiArt", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		
		rec := httptest.NewRecorder()
		asciiArt(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status %d; got %d", http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("Empty banner POST /asciiArt", func(t *testing.T) {
		form := url.Values{}
		form.Add("word", "hello")
		form.Add("banner", "")
		req := httptest.NewRequest("POST", "/asciiArt", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		
		rec := httptest.NewRecorder()
		asciiArt(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status %d; got %d", http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("Non-existent banner POST /asciiArt", func(t *testing.T) {
		form := url.Values{}
		form.Add("word", "hello")
		form.Add("banner", "nonexistent")
		req := httptest.NewRequest("POST", "/asciiArt", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		
		rec := httptest.NewRecorder()
		asciiArt(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Errorf("expected status %d; got %d", http.StatusNotFound, res.StatusCode)
		}
	})
}

func TestPrintAsciiArt(t *testing.T) {
	// We test with character space (ASCII 32). Max index accessed is (32-32)*9 + 9 = 9.
	bannerLines := make([]string, 10)
	for i := range bannerLines {
		bannerLines[i] = "[mock]"
	}

	result := functions.PrintAsciiArt(" ", bannerLines)
	// For 1 space, we print 9 lines, each having "[mock]" + "\n"
	expected := strings.Repeat("[mock]\n", 9)
	if result != expected {
		t.Errorf("expected %q; got %q", expected, result)
	}
}

package main

import (
	"ascii-art-web/functions"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/asciiArt", asciiArt)
	log.Println("Server running on port http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func asciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	if r.URL.Path != "/asciiArt" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	word := r.FormValue("word")
	bannerStyle := r.FormValue("banner")

	if word == "" {
		http.Error(w, "Word cannot be empty", http.StatusBadRequest)
		return
	}
	if bannerStyle == "" {
		http.Error(w, "Art style cannot be empty", http.StatusBadRequest)
		return
	}

	bannerPath := "banners/" + strings.ToLower(bannerStyle) + ".txt"
	bannerByte, err := os.ReadFile(bannerPath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Banner style not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	lines := strings.Split(string(bannerByte), "\n")
	art := functions.PrintAsciiArt(word, lines)

	tempErr := tmpl.ExecuteTemplate(w, "ascii.html", art)
	if tempErr != nil {
		log.Println("Template error:", tempErr)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

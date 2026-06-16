package main

import (
	"ascii-art-web/functions"
	"fmt"
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
	word := r.FormValue("word")
	banner := "banners/" + strings.ToLower(r.FormValue("banner")) + ".txt"

	if word == "" {
		http.Error(w, "Word cannot be empty", http.StatusBadRequest)
		return
	}
	if banner == "" {
		http.Error(w, "Art style cannot be empty", http.StatusBadRequest)
		return
	}

	bannerByte, err := os.ReadFile(banner)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	lines := strings.Split(string(bannerByte), "\n")
	art := functions.PrintAsciiArt(word, lines)
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	if r.URL.Path != "/asciiArt" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tempErr := tmpl.ExecuteTemplate(w, "ascii.html", art)
	if tempErr != nil {
		log.Fatal(tempErr)
	}
}

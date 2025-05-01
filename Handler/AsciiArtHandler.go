package handle

import (
	"fmt"
	"net/http"

	ascii "web/Ascii-Art"
)

type AsciiStruct struct {
	AsciiArt   string
	BannerName string
	Input      string
}

// IndexHandler handles GET requests to the root path ("/").
// It serves the main page of the ASCII art web application.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Error 405 | Mothod Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	renderTemplate(w, nil)
}

// AsciiArtHandler handles POST requests to the "/ascii-art" path.
// It processes the form data, generates ASCII art, and returns the result.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST.
	if r.Method != http.MethodPost {
		http.Error(w, "Error 405 | Mothod Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract form values
	inputUser := r.FormValue("inputUser")
	banner := r.FormValue("banner")

	// ensure text and banner are provided and valid
	if len(inputUser) > 1000 || inputUser == "" || !contains(banner) || !isValidInput(inputUser) {
		http.Error(w, "Error 400 | Bad Request", http.StatusBadRequest)
		return
	}

	// Generate ASCII art
	result, err := ascii.GenerateArt(inputUser, banner)
	if err != nil {
		http.Error(w, "Error 500 | Internal Server Error: Error generating ASCII art", http.StatusInternalServerError)
		return
	}

	data := AsciiStruct{
		AsciiArt:   result,
		BannerName: banner,
		Input:      inputUser,
	}

	if r.FormValue("submit") == "Download" {
		w.Header().Set("content-type", "text/plain")
		w.Header().Set("Content-Length", fmt.Sprintf("%v", len(result)))
		w.Header().Set("Content-Disposition", "attachment; filename=AsciiArt.txt")
		w.Write([]byte(result))
	} else {
		renderTemplate(w, data)
	}
	

}

// contains checks if the provided banner string is one of the valid options.
func contains(banner string) bool {
	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
		return true
	}
	return false
}

// is valid input, checks if the provided input string is printable
func isValidInput(inputUser string) bool {
	for _, char := range inputUser {
		if char == '\n' || char == '\r' {
			continue
		}
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}

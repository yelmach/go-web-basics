package handle

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate parses the HTML template and executes it with the provided data.
// It's used by both IndexHandler and AsciiArtHandler to render the page.
func renderTemplate(w http.ResponseWriter, data interface{}) {
	// Parse the HTML template file
	t, err := template.ParseFiles("Static/Templates/index.html")
	if err != nil {
		http.Error(w, "Error 500 | Internal Server Error: parsing template", http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Eroor 500 | Internal Server Error: executing template", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

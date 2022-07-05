package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"ascii-art-web/internal/ascii"
)

type Ascii struct {
	AsciiFont string
}

var files = []string{
	"./ui/html/home.html",
	"./ui/html/404NotFound.html",
	"./ui/html/405MethodNotAllowed.html",
	"./ui/html/500InternalServerError.html",
	"./ui/html/400BadRequest.html",
}

// GETHandler func receives only GET request and displays main page.
func GETHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "500InternalServerError.html", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		// http.Error(w, "ERROR-405\nMethod is not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(405)
		html.ExecuteTemplate(w, "405MethodNotAllowed.html", nil)
		return
	}
	if r.URL.Path != "/" {
		// http.Error(w, "ERROR-404\nPage not found", http.StatusNotFound)
		w.WriteHeader(404)
		html.ExecuteTemplate(w, "404NotFound.html", nil)
		return
	}
	if err = html.ExecuteTemplate(w, "home.html", nil); err != nil {
		fmt.Println(err)
	}
}

// Post handler responces only post request and processes data we receive through FromValue
// checks if text is correct and do not contain cyrilic alphabet, correct new lines,checks if font name is correct
// and if file that contains format font haven't been modified through HashSum & ConverFont func.
func POSTHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		html.ExecuteTemplate(w, "500InternalServerError.html", nil)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		html.ExecuteTemplate(w, "405MethodNotAllowed.html", nil)
		return
	}
	text := r.FormValue("formtext") // receives text to format
	font := r.FormValue("fonts")    // receives font's name
	text1 := ""
	// check if press enter for a new line
	for _, v := range text {
		if v == 10 || v == 13 { // checking if it's new line or '\r' (carriage ret)
			continue
		}
		if v < 32 || v > 126 { // if text is cyrrilic
			w.WriteHeader(400)
			html.ExecuteTemplate(w, "400BadRequest.html", nil)
			return
		}
	}
	// if text contains new line with carrige ret
	text1 = strings.ReplaceAll(text, "\r\n", "\n") // replace it with newline
	if text1 == "" {                               // if we receive empty text form
		w.WriteHeader(400)
		html.ExecuteTemplate(w, "400BadRequest.html", nil)
		return
	}
	// we need to modify font's name in order to open file with it
	switch font {
	case "":
		font = "standard.txt"
	case "Standard":
		font = "standard.txt"
	case "Shadow":
		font = "shadow.txt"
	case "Thinkertoy":
		font = "thinkertoy.txt"
	default:
		w.WriteHeader(404)
		html.ExecuteTemplate(w, "404NotFound.html", nil)
		return
	}
	if !ascii.ConvertFont(font) {
		w.WriteHeader(500)
		html.ExecuteTemplate(w, "500InternalServerError.html", nil)
		return
	}
	if ascii.HashSum(font) { // Checks font's file hashsum
		w.WriteHeader(400)
		html.ExecuteTemplate(w, "400BadRequest.html", nil)
		return
	}
	// Converting ascii output result  and saves it in string
	result := ascii.OutputAscii(text1, font)
	ascii := Ascii{
		AsciiFont: result,
	}
	if err = html.ExecuteTemplate(w, "home.html", ascii); err != nil {
		fmt.Println(err)
	}
}

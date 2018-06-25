package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("static/tracks.html"))

func printHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "tracks.html", tracks)
}

func byArtistHandler(w http.ResponseWriter, r *http.Request) {

}

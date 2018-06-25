package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("static/tracks.html"))

func printHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "tracks.html", tracks)
}

//func byArtistHandler(w http.ResponseWriter, r *http.Request) {
//	reversed, err := strconv.ParseBool(
//		r.URL.Query().Get("reversed"))
//	if err != nil{
//	http.Error(w, fmt.Errorf("can't parse bool: %v", err, http.StatusBadRequest)
//	}
//}

package main

import (
	"html/template"
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

var (
	templates = template.Must(template.ParseFiles("template/tracks.html"))

	tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
)

func main() {
	log.Fatal(http.ListenAndServe(":3001", setRouter()))
}

func setRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", printHandler)
	router.HandleFunc("/sort", sortHandler)

	return router
}

func printHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "tracks.html", tracks); err != nil {
		http.Error(w, "execute template error", http.StatusInternalServerError)
	}
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	var copiedTracks = make([]*Track, len(tracks))
	copy(copiedTracks, tracks)

	by := r.FormValue("by")
	reverse, err := strconv.ParseBool(r.FormValue("reverse"))
	if err != nil {
		http.Error(w, "parse error"+err.Error(), http.StatusInternalServerError)
		return
	}

	switch by {
	case "artist":
		byArtist(copiedTracks, reverse)
	case "track":
		byTrack(copiedTracks, reverse)
	default:
		http.Error(w, "wrong param", http.StatusBadRequest)
		return
	}

	if err := templates.ExecuteTemplate(w, "tracks.html", copiedTracks); err != nil {
		http.Error(w, "execute template error", http.StatusInternalServerError)
	}
}

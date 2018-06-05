package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var c contributors

func main() {
	startLoad()

	http.HandleFunc("/reload/contributors", reloadContribHandler)
	http.HandleFunc("/contributors", contribHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func startLoad() {
	err := c.load()
	if err != nil {
		log.Fatalf("failed to load contributors: %v", err)
	}

	log.Printf("successfully load %d contributors", len(c))
}

func reloadContribHandler(w http.ResponseWriter, r *http.Request) {
	err := c.load()
	if err != nil {
		log.Printf("failed to load contributors: %v", err)
		http.Error(w, fmt.Sprintf("failed to load contributors: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "successfully load %d contributors", len(c))
}

func contribHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	num, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, fmt.Sprintf("not valid page number: %v", err), http.StatusBadRequest)
		return
	}

	start := (num - 1) * itemsPERPAGE
	end := num * itemsPERPAGE
	if end > len(c) {
		end = len(c)
	}

	pageC := contribPage{CurrentNumber: num, Contributors: c[start:end], TotalNumber: len(c) / itemsPERPAGE}

	err = templates.ExecuteTemplate(w, "contributors.html", pageC)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

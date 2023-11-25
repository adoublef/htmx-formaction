package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/adoublef/prelude/template"
	"github.com/go-chi/chi/v5"
)

//go:embed all:*.html
var embedFS embed.FS
var fs = template.NewFS(embedFS)

const pageIndex = "index.html"

func main() {
	mux := chi.NewMux()
	mux.Get("/", handleIndex(fs))
	mux.Post("/1", handlePOST())
	mux.Post("/2", handlePOST())
	http.ListenAndServe(":8080", mux)
}

func handleIndex(fs *template.FS) http.HandlerFunc {
	t := fs.MustParse(pageIndex)
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}

func handlePOST() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("FormData(%s)", r.FormValue("a"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

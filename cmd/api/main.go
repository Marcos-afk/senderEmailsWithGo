package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	routes := chi.NewRouter()

	routes.Use(middleware.Logger)

	routes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")

		if search != "" {
      w.Write([]byte(search))
		}else {
			w.Write([]byte("Hello, World!"))
		}
	})

	routes.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Write([]byte(name))
	})

	http.ListenAndServe(":5000", routes)
}
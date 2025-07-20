package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()

    // Aquí montarás tus rutas reales
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

    log.Println("Starting Comments Play on :3000")
    http.ListenAndServe(":3000", r)
}
package reader

import (
    "net/http"

    "github.com/go-chi/chi/v5"
)

func RegisterReaderRoutes(r chi.Router) {
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Reader GET /"))
    })

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Reader GET /health"))
    })
}
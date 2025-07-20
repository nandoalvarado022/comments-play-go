package writer

import (
    "net/http"

    "github.com/go-chi/chi/v5"
)

func RegisterWriterRoutes(r chi.Router) {
    r.Post("/comment", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Writer POST /comment"))
    })
}

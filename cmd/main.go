package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nandoalvarado022/comments-play/internal/routes/reader"
	"github.com/nandoalvarado022/comments-play/internal/routes/writer"
)

func main() {
	r := chi.NewRouter()

	reader.RegisterReaderRoutes(r)
	writer.RegisterWriterRoutes(r)

	log.Println("Starting Comments Play on :3000")
	http.ListenAndServe(":3000", r)
}

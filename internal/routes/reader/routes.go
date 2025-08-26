package reader

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nandoalvarado022/comments-play/internal/config"
	"github.com/nandoalvarado022/comments-play/internal/handler/health"
	userHandlerPkg "github.com/nandoalvarado022/comments-play/internal/handler/user"
	userRepoPkg "github.com/nandoalvarado022/comments-play/internal/repository"
	servicePkg "github.com/nandoalvarado022/comments-play/internal/service"
)

func RegisterReaderRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Comments Play"))
	})

	r.Get("/health", health.Handler)

	// Users
	db, err := config.InitDB()
	if err != nil {
		panic("No se pudo conectar a la base de datos: " + err.Error())
	}
	userRepo := userRepoPkg.NewMySQLRepository(db)
	userService := servicePkg.NewUserService(userRepo)
	userHandler := userHandlerPkg.NewHandler(userService)

	r.Get("/users/{uid}", func(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "uid")
		r.URL.RawQuery = "uid=" + uid
		userHandler.GetUser(w, r)
	})
}

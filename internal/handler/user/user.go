package user

import (
	"encoding/json"
	"net/http"

	"github.com/nandoalvarado022/comments-play/internal/service"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("uid")

	user, err := h.userService.GetUserByUID(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

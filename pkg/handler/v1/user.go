package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.User.ListUsers(h.ctx)
	if err != nil {
		msg := "failed to get user"
		http.Error(w, msg, http.StatusInternalServerError)
		h.logger.Error(msg, zap.Error(err))
		return
	}

	b, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Error("failed to marshal user", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := h.repo.User.GetUser(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

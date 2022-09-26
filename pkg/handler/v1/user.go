package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) GetUsers(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := h.repo.User.ListUsers(ctx)
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
	})
}

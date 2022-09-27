package v1

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.repo.Review.ListReviews(h.ctx)
	if err != nil {
		msg := "failed to get review"
		http.Error(w, msg, http.StatusInternalServerError)
		h.logger.Error(msg, zap.Error(err))
		return
	}

	b, err := json.Marshal(reviews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Error("failed to marshal review", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

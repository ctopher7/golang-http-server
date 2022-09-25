package impl

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ctopher7/sirclo/v2/berat/model"
)

func (h *handler) InsertBerat(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req model.Berat
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err, errType := h.uc.InsertBerat(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), errType)
	}
}

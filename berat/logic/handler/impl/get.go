package impl

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *handler) GetBerat(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	id := 0
	query := r.URL.Query()
	if val, ok := query["id"]; ok && len(val) > 0 {
		temp, err := strconv.ParseInt(val[0], 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id = int(temp)
	}

	res, err := h.uc.GetBeratData(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

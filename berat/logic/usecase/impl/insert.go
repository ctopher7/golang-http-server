package impl

import (
	"context"
	"errors"
	"net/http"

	"github.com/ctopher7/sirclo/v2/berat/model"
)

func (u *usecase) InsertBerat(ctx context.Context, req model.Berat) (err error, errType int) {
	if req.Min > req.Max {
		return errors.New("min must be smaller than max"), http.StatusBadRequest
	}

	tx, err := u.repo.GetDB().BeginTx(ctx, nil)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	defer tx.Rollback()

	if err = u.repo.InsertBerat(ctx, tx, req); err != nil {
		errType = http.StatusInternalServerError
		return
	}
	if err = tx.Commit(); err != nil {
		errType = http.StatusInternalServerError
	}
	return
}

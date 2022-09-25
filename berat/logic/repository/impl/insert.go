package impl

import (
	"context"
	"database/sql"

	"github.com/ctopher7/sirclo/v2/berat/model"
)

func (r *repository) InsertBerat(ctx context.Context, tx *sql.Tx, req model.Berat) (err error) {
	_, err = tx.ExecContext(ctx, insertBeratQuery, req.Tanggal, req.Min, req.Max)

	return
}

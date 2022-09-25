package repository

import (
	"context"
	"database/sql"

	"github.com/ctopher7/sirclo/v2/berat/model"
)

type Repository interface {
	InsertBerat(ctx context.Context, tx *sql.Tx, req model.Berat) error
	GetDB() *sql.DB
	GetAllBerat(ctx context.Context) (res []model.Berat, err error)
	GetBeratByID(ctx context.Context, id int) (res model.Berat, err error)
}

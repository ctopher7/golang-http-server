package usecase

import (
	"context"

	"github.com/ctopher7/sirclo/v2/berat/model"
)

type Usecase interface {
	GetBeratData(ctx context.Context, id int) (res model.GetBeratRes, err error)
	InsertBerat(ctx context.Context, req model.Berat) (err error, errType int)
}

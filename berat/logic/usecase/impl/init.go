package impl

import (
	"github.com/ctopher7/sirclo/v2/berat/logic/repository"
	u "github.com/ctopher7/sirclo/v2/berat/logic/usecase"
)

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) u.Usecase {
	return &usecase{
		repo: repo,
	}
}

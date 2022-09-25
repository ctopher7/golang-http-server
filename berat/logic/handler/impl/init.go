package impl

import (
	h "github.com/ctopher7/sirclo/v2/berat/logic/handler"
	"github.com/ctopher7/sirclo/v2/berat/logic/usecase"
)

type handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) h.Handler {
	return &handler{
		uc: uc,
	}
}

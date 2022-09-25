package logic

import (
	"github.com/ctopher7/sirclo/v2/berat/logic/handler"
	hImpl "github.com/ctopher7/sirclo/v2/berat/logic/handler/impl"
	rImpl "github.com/ctopher7/sirclo/v2/berat/logic/repository/impl"
	uImpl "github.com/ctopher7/sirclo/v2/berat/logic/usecase/impl"
	"github.com/ctopher7/sirclo/v2/berat/resource"
)

func NewHandler(res *resource.Resource) handler.Handler {
	repo := rImpl.New(res)
	uc := uImpl.New(repo)
	return hImpl.New(uc)
}

package impl

import (
	"context"
	"testing"

	"github.com/ctopher7/sirclo/v2/berat/helper"
	"github.com/ctopher7/sirclo/v2/berat/logic/repository"
	"github.com/ctopher7/sirclo/v2/berat/model"
)

func Test_GetBeratData2(t *testing.T) {
	repoMock := new(repository.MockRepository)
	u := &usecase{
		repo: repoMock,
	}
	ctx := context.Background()
	type args struct {
		id int
	}

	helper.TestFramework(helper.TestFrameworkReq{
		Cases: []helper.Cases{
			{
				Name: "Success GetBeratByID",
				Want: []interface{}{
					model.GetBeratRes{
						RataRata: model.Berat{
							Max:       5,
							Min:       1,
							Perbedaan: 4,
						},
						Data: []model.Berat{
							{
								Max:       5,
								Min:       1,
								Perbedaan: 4,
								Id:        1,
							},
						},
					},
					nil,
				},
				MockFunc: []helper.MockFunc{
					{
						MockInterface: repoMock,
						MockFuncName:  "GetBeratByID",
						MockReq:       []interface{}{ctx, 1},
						MockRes: []interface{}{
							model.Berat{
								Max: 5,
								Min: 1,
								Id:  1,
							},
							nil,
						},
						TestForErr: helper.TestForErr{
							Test:        true,
							ErrPosition: 1,
						},
					},
				},
			},
			{
				Name: "Success GetAllBerat",
				Want: []interface{}{
					model.GetBeratRes{
						RataRata: model.Berat{
							Max:       7.5,
							Min:       1.5,
							Perbedaan: 6,
						},
						Data: []model.Berat{
							{
								Max:       10,
								Min:       2,
								Perbedaan: 8,
								Id:        2,
							},
							{
								Max:       5,
								Min:       1,
								Perbedaan: 4,
								Id:        1,
							},
						},
					},
					nil,
				},
				MockFunc: []helper.MockFunc{
					{
						MockInterface: repoMock,
						MockFuncName:  "GetAllBerat",
						MockReq:       []interface{}{ctx},
						MockRes: []interface{}{
							[]model.Berat{
								{
									Max:       10,
									Min:       2,
									Perbedaan: 8,
									Id:        2,
								},
								{
									Max:       5,
									Min:       1,
									Perbedaan: 4,
									Id:        1,
								},
							},
							nil,
						},
						TestForErr: helper.TestForErr{
							Test:        true,
							ErrPosition: 1,
						},
					},
				},
				PatchReq: func(input interface{}) interface{} {
					temp := input.(args)
					temp.id = 0
					return temp
				},
			},
		},
		T: t,
		TestFunc: func(input interface{}) []interface{} {
			inputArgs, ok := input.(args)
			if !ok {
				t.Errorf("Test_GetBeratData failed parsing argument")
			}
			got, err := u.GetBeratData(ctx, inputArgs.id)
			return []interface{}{got, err}
		},
		BaseReq: args{
			id: 1,
		},
	})
}

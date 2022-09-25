package impl

import (
	"context"

	"github.com/ctopher7/sirclo/v2/berat/model"
)

func (u *usecase) GetBeratData(ctx context.Context, id int) (res model.GetBeratRes, err error) {
	if id > 0 {
		temp, errGet := u.repo.GetBeratByID(ctx, id)
		if errGet != nil {
			err = errGet
			return
		}
		temp.Perbedaan = temp.Max - temp.Min
		res.Data = append(res.Data, temp)
		res.RataRata = temp
		res.RataRata.Id = 0
		return
	}

	data, err := u.repo.GetAllBerat(ctx)
	if err != nil {
		return
	}
	res.Data = data
	for i, val := range data {
		res.Data[i].Perbedaan = val.Max - val.Min
		res.RataRata.Max += val.Max
		res.RataRata.Min += val.Min
		res.RataRata.Perbedaan += res.Data[i].Perbedaan
	}

	res.RataRata.Max = res.RataRata.Max / float64(len(res.Data))
	res.RataRata.Min = res.RataRata.Min / float64(len(res.Data))
	res.RataRata.Perbedaan = res.RataRata.Perbedaan / float64(len(res.Data))

	return
}

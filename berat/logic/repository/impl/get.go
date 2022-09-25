package impl

import (
	"context"

	"github.com/ctopher7/sirclo/v2/berat/model"
)

func (r *repository) GetAllBerat(ctx context.Context) (res []model.Berat, err error) {
	rows, err := r.postgresDb.QueryContext(ctx, GetAllBeratQuery)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		temp := model.Berat{}
		err = rows.Scan(&temp.Id, &temp.Tanggal, &temp.Max, &temp.Min)
		if err != nil {
			return
		}
		res = append(res, temp)
	}

	return
}

func (r *repository) GetBeratByID(ctx context.Context, id int) (res model.Berat, err error) {
	row := r.postgresDb.QueryRowContext(ctx, GetBeratByIdQuery, id)

	err = row.Scan(&res.Id, &res.Tanggal, &res.Max, &res.Min)

	return
}

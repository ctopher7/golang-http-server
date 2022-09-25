package impl

import (
	"database/sql"

	r "github.com/ctopher7/sirclo/v2/berat/logic/repository"
	"github.com/ctopher7/sirclo/v2/berat/resource"
)

type repository struct {
	postgresDb *sql.DB
}

func New(res *resource.Resource) r.Repository {
	return &repository{
		postgresDb: res.PostgresDb,
	}
}

func (r *repository) GetDB() *sql.DB {
	return r.postgresDb
}

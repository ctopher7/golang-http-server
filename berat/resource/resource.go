package resource

import (
	"database/sql"
	"fmt"

	"github.com/ctopher7/sirclo/v2/berat/config"
	_ "github.com/lib/pq"
)

type Resource struct {
	PostgresDb *sql.DB
}

func Init(cfg *config.Config) *Resource {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.PostgresDb.Host, cfg.PostgresDb.Port, cfg.PostgresDb.User, cfg.PostgresDb.Password, cfg.PostgresDb.Dbname)
	fmt.Println(psqlconn)
	db, _ := sql.Open("postgres", psqlconn)

	return &Resource{
		PostgresDb: db,
	}
}

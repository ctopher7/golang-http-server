package berat

import (
	"fmt"
	"net/http"

	"github.com/ctopher7/sirclo/v2/berat/config"
	"github.com/ctopher7/sirclo/v2/berat/resource"
	"github.com/ctopher7/sirclo/v2/berat/route"
)

func Server(env string) error {
	//init config
	cfg, err := config.ReadConfig(env)
	if err != nil {
		return err
	}

	//init resources
	res := resource.Init(cfg)
	defer res.PostgresDb.Close()

	route.New(res)

	fmt.Printf("running berat server on %s \n", cfg.ServerAddress)
	return http.ListenAndServe(cfg.ServerAddress, nil)
}

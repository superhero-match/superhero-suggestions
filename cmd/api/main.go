package main

import (
	"github.com/superhero-suggestions/cmd/api/controller"
	"github.com/superhero-suggestions/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	ctrl, err := controller.NewController(cfg)
	if err != nil {
		panic(err)
	}

	r := ctrl.RegisterRoutes()

	err = r.RunTLS(
		":4000",
		"./cmd/api/certificate.pem",
		"./cmd/api/key.pem",
	)
	if err != nil {
		panic(err)
	}
}
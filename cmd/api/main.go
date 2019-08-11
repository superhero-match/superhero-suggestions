package main

import (
	"github.com/superhero-suggestions/cmd/api/controller"
)

func main() {
	ctrl := controller.NewController()

	r := ctrl.RegisterRoutes()

	err := r.RunTLS(
		":4000",
		"./cmd/api/certificate.pem",
		"./cmd/api/key.pem",
	)
	if err != nil {
		panic(err)
	}
}
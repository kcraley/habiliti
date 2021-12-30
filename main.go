package main

import (
	"github.com/kcraley/habiliti/internal/server"
	"github.com/kcraley/habiliti/pkg/terraform"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("starting application server")

	tfReg := terraform.NewRegistry(&terraform.RegistryOptions{})
	log.Infof("%+v", tfReg.Options())
	srv := server.New(&server.Options{
		Address:           "0.0.0.0",
		Port:              "8080",
		TerraformRegistry: tfReg,
	})
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("an error occurred: %v", err)
	}
}

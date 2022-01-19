package server

import (
	"github.com/kcraley/habiliti/pkg/terraform"
)

// Options is a set of attributes which configures the application server
type Options struct {
	Address           string              // the address to serve the application
	Endpoint          string              // the root endpoint which serves the Terraform registry
	Port              string              // the port to serve the application
	TerraformRegistry *terraform.Registry // Registry is an instance of a Terraform Registry
}

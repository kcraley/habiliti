package server

import (
	"github.com/kcraley/habiliti/pkg/terraform"
)

// Options is a set of attributes which configures the application server
type Options struct {
	// Address is the IP address which the server is listening on
	Address string
	// Port is the port which the server is listening on
	Port string
	// Registry is an instance of a Terraform Registry
	TerraformRegistry *terraform.Registry
}

// init is responsible for setting sane defaults for the application server
func (opt *Options) init() {
	if opt.Address == "" {
		opt.Address = "0.0.0.0"
	}
	if opt.Port == "" {
		opt.Port = "8080"
	}
}

package terraform

import (
	log "github.com/sirupsen/logrus"
)

type baseRegistry struct {
	opt *RegistryOptions
}

func newBaseRegistry(opt *RegistryOptions) *baseRegistry {
	return &baseRegistry{
		opt: opt,
	}
}

func (br *baseRegistry) Options() *RegistryOptions {
	return br.opt
}

// RegistryOptions is a set of options which configures the
// creation of a new Terraform registry.
type RegistryOptions struct {
	// Controls enabling the login endpoint
	EnableLogin bool
	// Controls enabling the modules endpoint
	EnableModules bool
	// Control enabling the providers endpoint
	EnableProviders bool
}

// Registry represents an instance of a Terraform registry.
type Registry struct {
	*baseRegistry
}

// NewRegistry creates and returns an instance of a Terraform registry.
func NewRegistry(opt *RegistryOptions) *Registry {
	log.Infof("Creating new reg: %+v", opt)

	return &Registry{
		baseRegistry: newBaseRegistry(opt),
	}
}

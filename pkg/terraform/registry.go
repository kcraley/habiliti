package terraform

import (
	log "github.com/sirupsen/logrus"
)

const (
	defaultLoginPath    = "/v1/login/"
	defaultModulePath   = "/v1/modules/"
	defaultProviderPath = "/v1/providers/"
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

// JSONMarshal

// RegistryOptions is a set of options which configures the
// creation of a new Terraform registry.
type RegistryOptions struct {
	// Relative path whichs serves the login discovery protocol
	// REF: https://www.terraform.io/docs/internals/login-protocol.html
	LoginPath string `json:"login.v1,omitempty"`

	// Relative path which serves the module discovery protocol
	// REF: https://www.terraform.io/docs/internals/module-registry-protocol.html
	ModulePath string `json:"modules.v1,omitempty"`

	// Relative path which serves the provider discovery protocol
	// REF: https://www.terraform.io/docs/internals/provider-registry-protocol.html
	ProviderPath string `json:"providers.v1,omitempty"`
}

func (ro *RegistryOptions) init() {
	if ro.LoginPath == "" {
		ro.LoginPath = defaultLoginPath
		log.Infof("no custom login path given in options, setting to: %s", ro.LoginPath)
	}
	if ro.ModulePath == "" {
		ro.ModulePath = defaultModulePath
		log.Infof("no custom module path given in options, setting to: %s", ro.ModulePath)
	}
	if ro.ProviderPath == "" {
		ro.ProviderPath = defaultProviderPath
		log.Infof("no custom provider path given in options, setting to: %s", ro.ProviderPath)
	}
}

// Registry represents an instance of a Terraform registry.
type Registry struct {
	*baseRegistry
}

// NewRegistry creates and returns an instance of a Terraform registry.
func NewRegistry(opt *RegistryOptions) *Registry {
	opt.init()
	log.Infof("Creating new reg: %+v", opt)

	return &Registry{
		baseRegistry: newBaseRegistry(opt),
	}
}

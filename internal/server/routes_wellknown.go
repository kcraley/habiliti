package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type WellKnownResponse struct {
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

func (s *Server) buildWellKnownResponse() *WellKnownResponse {
	wk := &WellKnownResponse{}
	if s.opt.TerraformRegistry.Options().EnableLogin {
		wk.LoginPath = fmt.Sprintf("%s/login/", s.opt.Endpoint)
	}
	if s.opt.TerraformRegistry.Options().EnableModules {
		wk.ModulePath = fmt.Sprintf("%s/modules/", s.opt.Endpoint)
	}
	if s.opt.TerraformRegistry.Options().EnableProviders {
		wk.ProviderPath = fmt.Sprintf("%s/providers/", s.opt.Endpoint)
	}
	return wk
}

func (s *Server) handleWellKnown() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		regJSON, err := json.Marshal(s.buildWellKnownResponse())
		if err != nil {
			log.Errorf("failed returning response for `.well-known/terraform.json`, %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(regJSON)
	}
}

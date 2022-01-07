package server

import (
	"encoding/json"
	"net/http"

	"github.com/kcraley/habiliti/pkg/terraform"
)

type modulesResponse struct {
	Modules []moduleVersions `json:"modules"`
}

type moduleVersions struct {
	Source   string          `json:"source"`
	Versions []moduleVersion `json:"versions"`
}

type moduleVersion struct {
	Version    string           `json:"version"`
	Root       moduleRoot       `json:"root"`
	Submodules moduleSubmodules `json:"submodules"`
}

type moduleRoot struct {
	Providers    []terraform.Provider   `json:"providers"`
	Dependencies []terraform.Dependency `json:"dependencies"`
}

type moduleSubmodules struct {
	Path         string                 `json:"path"`
	Providers    []terraform.Provider   `json:"providers"`
	Dependencies []terraform.Dependency `json:"dependencies"`
}

func (s *Server) handleModuleVersions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mResp := &modulesResponse{
			Modules: []moduleVersions{
				{
					Source: "hashicorp/consul/aws",
					Versions: []moduleVersion{
						{
							Version: "0.7.3",
							Root: moduleRoot{
								Providers: []terraform.Provider{
									{
										Name: "aws",
									},
								},
								Dependencies: []terraform.Dependency{},
							},
							Submodules: moduleSubmodules{
								Path:         "",
								Providers:    []terraform.Provider{},
								Dependencies: []terraform.Dependency{},
							},
						},
					},
				},
			},
		}
		// w.Write([]byte("Module Versions"))
		mRespBytes, err := json.Marshal(mResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed marshaling json response"))
		}
		w.Write(mRespBytes)
	}
}

func (s *Server) handleModuleDownload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Module Download"))
	}
}

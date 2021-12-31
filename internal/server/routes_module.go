package server

import (
	"net/http"
)

// type modulesResponse struct {
// 	modules []moduleVersions `json:"modules"`
// }

// type moduleVersions struct {
// 	source   string          `json:"source"`
// 	versions []moduleVersion `json:"versions"`
// }

// type moduleVersion struct {
// 	version    string
// 	root       moduleRoot       `json:"root"`
// 	submodules moduleSubmodules `json:"submodules"`
// }

// type moduleRoot struct {
// 	providers    []terraform.Provider   `json:"providers"`
// 	dependencies []terraform.Dependency `json:"dependencies"`
// }

// type moduleSubmodules struct {
// 	path         string                   `json:"path"`
// 	providers    []terraform.Provider     `json:"providers"`
// 	dependencies []terraform.Dependencies `json:"dependencies"`
// }

func (s *Server) handleModuleVersions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Module Versions"))
	}
}

func (s *Server) handleModuleDownload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Module Download"))
	}
}

package server

import "net/http"

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

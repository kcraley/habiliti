package server

import "net/http"

func (s *Server) handleModuleVersions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Module Versions"))
}

func (s *Server) handleModuleDownload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Module Download"))
}

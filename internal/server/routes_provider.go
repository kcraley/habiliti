package server

import "net/http"

func (s *Server) handleProviderVersions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Provider Versions"))
}

func (s *Server) handleProviderDownload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Provider Download"))
}

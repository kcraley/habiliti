package server

import "net/http"

func (s *Server) handleProviderVersions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Provider Versions"))
	}
}

func (s *Server) handleProviderDownload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Provider Download"))
	}
}

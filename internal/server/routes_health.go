package server

import "net/http"

// handleLiveAndReadiness is the single route which responds to liveness
// and readiness probes to determine application availability.
// TODO: This should really be broken out into two separate sets of logic.
func (s *Server) handleLiveAndReadiness(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "ok!"}`))
}

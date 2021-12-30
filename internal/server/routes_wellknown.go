package server

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (s *Server) handleWellKnown(w http.ResponseWriter, r *http.Request) {
	regJSON, err := json.Marshal(s.RegistryOptions())
	if err != nil {
		log.Errorf("failed returning response for `.well-known/terraform.json`, %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(regJSON)
}

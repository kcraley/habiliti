package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kcraley/habiliti/pkg/terraform"
	log "github.com/sirupsen/logrus"
)

type baseServer struct {
	opt *Options
}

func newBaseServer(opt *Options) *baseServer {
	return &baseServer{
		opt: opt,
	}
}

func (b *baseServer) RegistryOptions() *terraform.RegistryOptions {
	return b.opt.TerraformRegistry.Options()
}

// Server represents an application server
type Server struct {
	*baseServer
	ctx context.Context
	mux *mux.Router
}

// New creates and returns a new instance of an application server
func New(opt *Options) *Server {
	return &Server{
		baseServer: newBaseServer(opt),
		ctx:        context.Background(),
		mux:        mux.NewRouter(),
	}
}

// initializeMiddleware adds the necessary middleware to the router
func (s *Server) initializeMiddleware() {
	// Configure common middleware
	s.mux.Use(setHeaders)
	s.mux.Use(logRequests)
}

// initializeRoutes adds the necessary handler functions to the server
func (s *Server) initializeRoutes() {
	// Add health endpoints
	healthSubrouter := s.mux.PathPrefix("/healthz").Subrouter()
	healthSubrouter.HandleFunc("/liveness", s.handleLiveAndReadiness()).Methods(http.MethodGet)
	healthSubrouter.HandleFunc("/readiness", s.handleLiveAndReadiness()).Methods(http.MethodGet)

	// Main endpoints which handle Terraform's service discovery
	// REF: https://www.terraform.io/internals/remote-service-discovery
	s.mux.HandleFunc("/.well-known/terraform.json", s.handleWellKnown()).Methods(http.MethodGet)

	modSubrouter := s.mux.PathPrefix("/v1/modules/{namespace}/{name}/{system}").Subrouter()
	modSubrouter.HandleFunc("/versions", s.handleModuleVersions()).Methods(http.MethodGet)
	modSubrouter.HandleFunc("/{version}/download", s.handleModuleDownload()).Methods(http.MethodGet)

	provSubrouter := s.mux.PathPrefix("/v1/providers/{namespace}/{type}").Subrouter()
	provSubrouter.HandleFunc("/versions", s.handleProviderVersions()).Methods(http.MethodGet)
	provSubrouter.HandleFunc("/download/{os}/{arch}", s.handleProviderDownload()).Methods(http.MethodGet)
}

// ListenAndServe starts the application server
func (s *Server) ListenAndServe(address, port string) (err error) {
	// Add middleware and routes to mux router
	s.initializeMiddleware()
	s.initializeRoutes()

	// Create and run a new server
	httpsrv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", address, port),
		Handler:      s.mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
	}
	if err := httpsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("An error occurred: %s\n", err)
	} else {
		log.Info("Successfully shut down server...")
	}

	return
}

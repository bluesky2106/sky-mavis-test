package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bluesky2106/sky-mavis-test/part-3/backend/config"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/interfaces"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/services"
	"github.com/gin-gonic/gin"
)

// Server : struct
type Server struct {
	config     *config.Config
	g          *gin.Engine
	server     *http.Server
	visitorSvc interfaces.IVisitorService
}

// NewServer : userSvc, walletSvc, assetSvc, config
func NewServer(config *config.Config,
	g *gin.Engine,
	visitorSvc *services.VisitorService,
) *Server {
	return &Server{
		config: config,
		g:      g,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
			Handler: g,
		},
		visitorSvc: visitorSvc,
	}
}

// Run application
func (s *Server) Run() error {
	log.Printf("Start the server at : %s", s.server.Addr)

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Close app and all the resources
func (s *Server) Close(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.server.Shutdown(ctx)
}

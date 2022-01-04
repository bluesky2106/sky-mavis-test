package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultWelcome : ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Homepage")
}

// Welcome : ...
func (s *Server) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "REST API Homepage")
}

// this is for health check procedure
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (s *Server) liveCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

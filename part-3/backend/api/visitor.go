package api

import (
	"net/http"

	"github.com/bluesky2106/sky-mavis-test/part-3/backend/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (s *Server) GetCurrentVisitor(c *gin.Context) {
	timeoutCtx, cancel := createTimeoutContext(c)
	defer cancel()

	ip := c.ClientIP()
	// fmt.Println(ip)
	res, err := s.visitorSvc.GetCurrentVisitor(timeoutCtx, ip)
	if err != nil {
		log.GetLogger().Error("s.visitorSvc.GetCurrentVisitor", zap.Field{Key: "Internal Error", Type: zapcore.ErrorType, Interface: err})
		respondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, Response{Data: res})
}

func (s *Server) GetLast100Visitors(c *gin.Context) {
	timeoutCtx, cancel := createTimeoutContext(c)
	defer cancel()

	res, err := s.visitorSvc.GetLast100Visitors(timeoutCtx)
	if err != nil {
		log.GetLogger().Error("s.visitorSvc.GetLast100Visitors", zap.Field{Key: "Internal Error", Type: zapcore.ErrorType, Interface: err})
		respondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, Response{Data: res})
}

func (s *Server) GetTop100Visitors(c *gin.Context) {
	timeoutCtx, cancel := createTimeoutContext(c)
	defer cancel()

	res, err := s.visitorSvc.GetTop100Visitors(timeoutCtx)
	if err != nil {
		log.GetLogger().Error("s.visitorSvc.GetTop100Visitors", zap.Field{Key: "Internal Error", Type: zapcore.ErrorType, Interface: err})
		respondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, Response{Data: res})
}

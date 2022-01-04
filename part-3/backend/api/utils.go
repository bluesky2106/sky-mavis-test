package api

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

const (
	timeout = 30 * time.Second
)

func createTimeoutContext(ctx *gin.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx.Request.Context(), timeout)
}

func respondError(c *gin.Context, status int, err error) {
	c.JSON(status, Response{Error: err})
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bluesky2106/sky-mavis-test/part-3/backend/api"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/config"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/libs/storage"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Get global config
	conf := config.GetConfig()
	conf.Print()

	// Init logger
	log.InitLogger(conf.Env)

	storage := storage.NewGCStorage(conf.GCStorageCredentials, conf.GCStorageBucketName, conf.GCStorageBaseURL)
	fmt.Print(storage)

	if conf.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	svr := api.NewServer(conf, router)
	// authMw := api.AuthMiddleware(string(conf.TokenSecretKey), svr.Authenticate)
	// svr.Routes(authMw)
	go func() {
		if err := svr.Run(); err != nil {
			log.GetLogger().Error("svr.Run", zap.Error(err))
		}
	}()

	waitForInterruptSignal()

	// give 10 seconds to finish ongoing requests
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := svr.Close(ctx); err != nil {
		log.GetLogger().Error("svr.Close", zap.Error(err))
	}
}

func waitForInterruptSignal() {
	// wait for interrupt signal here
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

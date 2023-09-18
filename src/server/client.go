package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	sources "study/internal/api"
	"study/internal/config"
	"time"
)

func Run() {
	cfg := config.GetConfig()

	r := gin.Default()

	if cfg.IsDebug {
		err := r.SetTrustedProxies([]string{"0.0.0.0"})
		if err != nil {
			log.Fatal(err)
		}
	}

	r.POST("/authors", sources.CreateAuthor)
	r.GET("/authors", sources.GetAuthors)
	r.GET("/authors/:id", sources.GetAuthorByID)
	r.DELETE("/authors/:id", sources.DeleteAuthor)
	r.PUT("/authors/:id", sources.UpdateAuthor)

	// listen and serve on 0.0.0.0:8080
	go func() {
		if err := r.Run(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	_, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
}

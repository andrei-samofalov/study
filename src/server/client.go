package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	sources "study/internal/api"
	"study/internal/config"
	"syscall"
	"time"
)

func Run() {
	cfg := config.GetConfig()

	router := gin.Default()

	if cfg.IsDebug {
		err := router.SetTrustedProxies([]string{"0.0.0.0"})
		if err != nil {
			log.Fatal(err)
		}
	}
	api := router.Group("/api")
	{
		authors := api.Group("/authors")
		{
			authors.POST("/", sources.CreateAuthor)
			authors.GET("/", sources.GetAuthors)
			authors.GET("/:id", sources.GetAuthorByID)
			authors.DELETE("/:id", sources.DeleteAuthor)
			authors.PUT("/:id", sources.UpdateAuthor)
		}
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// listen and serve on 0.0.0.0:8080
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

package main

import (
	"context"
	"go-fcontrol-api/src/configs"
	"go-fcontrol-api/src/routers"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := configs.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	configs.InitLogger()
	configs.InitMongo()
	defer configs.DisconnectMongo()

	routes := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.AllowCredentials = true

	routes.Use(cors.New(config))

	routers.InitRouter(routes)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	go func() {
		slog.Info("Starting server on :8080")
		slog.Debug("Debug mode enabled")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %s", err)
	}

	slog.Info("Server exited gracefully")
}

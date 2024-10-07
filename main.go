package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/qinyul/messaging-api/configuration"
	"github.com/qinyul/messaging-api/controller"
	"github.com/qinyul/messaging-api/repository"
	"github.com/qinyul/messaging-api/router"
	"github.com/qinyul/messaging-api/service"
)

func main() {
	cfg, err := configuration.LoadConfig()
	if err != nil {
		slog.Error("Error loading configuration", "error", err)
	}

	db := configuration.NewDatabaseConfig(cfg)
	db.ConnectDatabase()
	db.Migrate()

	messageRepo := repository.NewMessageRepository(db.DB)
	messageService := service.NewMessageService(messageRepo)
	messageController := controller.NewMessageController(messageService)

	mux := router.NewRouter(*messageController)

	server := &http.Server{
		Addr:    ":" + cfg.PORT,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server: ", "error", err)
		}
	}()

	slog.Info(fmt.Sprintf("Server is running on port %s", cfg.PORT))

	quitCH := make(chan os.Signal, 1)
	signal.Notify(quitCH, os.Interrupt)

	<-quitCH
	slog.Info("Received termination signal, shutting down server...")
	shutdownCTX, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(shutdownCTX); err != nil {
		slog.Error("Failed to gracefully shut down server", "error", err)
	}
	slog.Info("Server shutdown sucessfull")
}

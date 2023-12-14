package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"time"
	"todolist/config"
	"todolist/internal/server"
	"todolist/internal/server/generated"
	"todolist/internal/storage"
	"todolist/internal/usecase"
	"todolist/pkg/logger"
	"todolist/pkg/pg"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%v", err)
	}
}

func run() error {
	time.Local = time.UTC
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	logZap, err := logger.New(cfg.Server.LogLevel)
	defer logZap.Fatal("application stopped")

	entDriver, err := pg.EntDriver(cfg.DB)
	if err != nil {
		logZap.Error("", zap.Error(err))
		return err
	}

	dbStorage, err := storage.New(ctx, logZap, cfg, entDriver)
	// Use cases
	useCases := usecase.New(cfg, logZap, dbStorage)

	// Server
	srv := server.New(logZap, cfg, useCases)
	options := srv.NewServerOptions()

	err = startServer(ctx, logZap, &http.Server{
		Handler:           generated.HandlerWithOptions(srv, options),
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
	})

	if err != nil {
		return fmt.Errorf("unable to start server: %w", err)
	}

	return nil
}

func startServer(ctx context.Context, logZap *zap.Logger, server *http.Server) error {
	logZap.Info("starting server", zap.String("addr", server.Addr))

	server.BaseContext = func(net.Listener) context.Context {
		return ctx
	}

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			logZap.Info("server has closed")

			return nil
		}

		return err
	case <-ctx.Done():
		logZap.Info("closing server due to context cancellation")

		if err := server.Close(); err != nil {
			return fmt.Errorf("failed to close http server: %w", err)
		}

		return nil
	}
}

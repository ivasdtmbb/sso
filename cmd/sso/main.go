// Entry point to our (app)
// go run cmd/sso/main.go --config=./config/local.yaml

package main

import (
	"log/slog"
	"os"
	"syscall"
	"os/signal"

	"sso/internal/config"
	"sso/internal/app"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//---------------initialize config object----------------//
	cfg := config.MustLoad()

	//---------------initialize logger-----------------------//
	log := setupLogger(cfg.Env)

	log.Info("starting application",
		slog.Any("cfg", cfg),
	)

	//---------------initialize an application (app)---------//
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	//---------------run gRPC-server (app)-------------------//
	go application.GRPCSrv.MustRun()

	//---------------Graceful shutdown-----------------------//
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	stopsignal := <-stop

	log.Info("stopping application", slog.String("signal", stopsignal.String()))
	
	application.GRPCSrv.Stop()
	
	log.Info("application stopped")
}

// Logger initialization
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	
	return log
}

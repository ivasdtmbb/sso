package app

import (
	"log/slog"
	"time"

	//"sso/internal/grpc/auth"
	"sso/internal/app/grpc"
	//"google.golang.org/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: initialize storage

	// TODO; init auth service (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}

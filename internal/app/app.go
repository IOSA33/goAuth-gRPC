package app

import (
	grpcapp "authService/internal/app/grpc"
	"authService/internal/services/auth"
	"authService/internal/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: init storage
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	// TODO: init auth service
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCrv: grpcApp,
	}
}

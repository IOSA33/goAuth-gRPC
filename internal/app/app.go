package app

import (
	grpcapp "authService/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: init storage

	// TODO: init auth service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCrv: grpcApp,
	}
}

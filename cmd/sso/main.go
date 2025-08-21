package main

import (
	"authService/internal/app"
	"authService/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// Init config: cleanEnv
	cfg := config.MustLoad()

	// Init logger: slog
	log := setupLogger(cfg.Env)
	log.Info("Starting application",
		slog.Any("config", cfg),
	)

	// Confining application from config file
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	// go run cmd/sso/main.go --config=./config/local.yaml
	// Starting server
	go application.GRPCrv.MustRun()

	// How it works: How we know channel is a lock function, so it waits until something
	// will be written in stop channel and only then it'll do unlock and continue reading code rows.
	// And in a different go routine "go application.GRPCrv.MustRun()"
	// application will get serve requests.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// If signal.Notify requests some signal, the signal will be written in stop channel,
	// and after that it will unlock the channel and the code will continue to the next rows.
	// So how it works, after starting application it waits for the stop signals and after that
	// it goes to the row "application.GRPCrv.Stop()" to do graceful shutdown.
	sign := <-stop

	log.Info("application stopping", slog.String("signal:", sign.String()))

	// Graceful shutdown
	application.GRPCrv.Stop()

	// Printing "application stopped"
	log.Info("application stopped")
}

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

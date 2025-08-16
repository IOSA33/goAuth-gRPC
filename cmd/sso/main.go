package main

import (
	"authService/internal/config"
	"fmt"
)

func main() {
	// TODO: init config: cleanEnv
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: app
}

package main

import (
	"log/slog"
	"music-playlist/internal/config"
	"music-playlist/internal/repository"
	"music-playlist/internal/service"
	"music-playlist/pkg/logger/handlers/slogpretty"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	setupLogger(cfg.Env)
	slog.Info("starting application", slog.Any("env", &cfg))

	repo := repository.NewPlaylist(cfg.StoragePath)
	cache := service.InitCache(repo)

	service.NewPlaylist(repo, cache)
}

func setupLogger(env string) {
	switch env {
	case envLocal:
		slog.SetDefault(setupPrettySlog())
	case envDev:
		slog.SetDefault(
			slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
		)
	}
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug}}
	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

package main

import (
	"log/slog"
	"music-playlist/internal/config"
	"music-playlist/internal/domain"
	"music-playlist/internal/repository"
	"music-playlist/pkg/logger/handlers/slogpretty"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting application", slog.Any("env", &cfg))

	repo := repository.NewPlaylist(cfg.StoragePath)

	err := repo.Upload([]domain.Song{})
	if err != nil {
		return
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug}}
	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

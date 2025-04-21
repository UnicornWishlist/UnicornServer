package main

import (
	"log/slog"

	"github.com/UnicornWishlist/UnicornServer/lib/config"
	"github.com/UnicornWishlist/UnicornServer/lib/slog/logger"
)

func main() {
	cfg := config.MustLoadConfig()
	log := logger.MustCreateLogger(cfg)
	log = log.With(slog.String("app", "unicorn-server"))

	log.Debug("config initialized", slog.Any("config", cfg))
	log.Debug("Logger created")

	log.Info("App started")
}

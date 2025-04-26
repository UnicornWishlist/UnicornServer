package main

import (
	"log/slog"
	"os"

	"github.com/UnicornWishlist/UnicornServer/lib/config"
	"github.com/UnicornWishlist/UnicornServer/lib/services"
	"github.com/UnicornWishlist/UnicornServer/lib/slog/logger"
	"github.com/UnicornWishlist/UnicornServer/lib/slog/logger/sl"
	"github.com/UnicornWishlist/UnicornServer/lib/storage"
)

func main() {
	cfg := config.MustLoadConfig()
	log := logger.MustCreateLogger(cfg)
	log = log.With(slog.String("app", "unicorn-server"))

	log.Debug("config initialized", slog.Any("config", cfg))
	log.Debug("Logger created")

	storage, err := storage.InitStorage(log, &cfg.Database)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		os.Exit(1)
	}
	log.Info("App finished successfully")

	userService := services.NewUserService(log, storage)

	_ = userService.CreateUser("Alonka")
}

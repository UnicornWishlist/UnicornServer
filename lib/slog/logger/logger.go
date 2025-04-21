package logger

import (
	"log/slog"
	"os"

	"github.com/UnicornWishlist/UnicornServer/lib/config"
	"github.com/UnicornWishlist/UnicornServer/lib/slog/prettyslog"
)

func MustCreateLogger(cfg *config.Config) *slog.Logger {
	var handler slog.Handler

	switch cfg.Env {
	case config.EnvProduction:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	case config.EnvDevelopment:
		handler = prettylog.New(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}, nil)
	default:
		panic("unable to intialize logger: unknown envirement")
	}

	log := slog.New(handler)

	return log
}

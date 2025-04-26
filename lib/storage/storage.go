package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"

	"github.com/UnicornWishlist/UnicornServer/lib/config"
	"github.com/UnicornWishlist/UnicornServer/lib/slog/logger/sl"
)

const (
	ErrPostgresDuplicateKey = "23505"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type Storage struct {
	Db  *sql.DB
	log *slog.Logger
}

func InitStorage(log *slog.Logger, cfg *config.DatabaseConfig) (*Storage, error) {
	op := "storage.InitStorage"
	log = log.With("op", op)

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("failed to open database connection", sl.Err(err))
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		log.Error("failed to ping database", sl.Err(err))
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Info(
		"connected to database",
		slog.String("host", cfg.Host),
		slog.Int("port", cfg.Port),
		slog.String("database", cfg.Database),
	)

	return &Storage{Db: db, log: log}, nil
}

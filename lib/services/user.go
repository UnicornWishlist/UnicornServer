package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/UnicornWishlist/UnicornServer/lib/slog/logger/sl"
	"github.com/UnicornWishlist/UnicornServer/lib/storage"
	"github.com/lib/pq"
)

type User struct {
	ID   int
	Name string
}

type UserService struct {
	log     *slog.Logger
	storage *storage.Storage
}

func NewUserService(log *slog.Logger, s *storage.Storage) *UserService {
	return &UserService{
		log:     log,
		storage: s,
	}
}

var ErrUserNotFound = errors.New("user not found")

func (s *UserService) GetUserByID(id int) (*User, error) {
	op := "userService.GetUserByID"
	log := s.log.With("op", op, "id", id)

	query := "SELECT id, name FROM users WHERE id = $1"
	log.Debug("Executing query", slog.String("query", query))
	row := s.storage.Db.QueryRow(query, id)
	user := &User{}

	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Debug("User not found")
			return nil, ErrUserNotFound
		}

		log.Debug("Failed to fetch user", sl.Err(err))

		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	log.Debug("User fetched successfully", slog.Any("user", user))
	return user, nil
}

func (s *UserService) CreateUser(name string) error {
	op := "userService.CreateUser"
	log := s.log.With("op", op, "name", name)

	query := "INSERT INTO users (name) VALUES ($1) RETURNING id"
	log.Debug("Executing query", slog.String("query", query))

	_, err := s.storage.Db.Exec(query, name)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == storage.ErrPostgresDuplicateKey {
			log.Debug("User already exists")
			return storage.ErrUserAlreadyExists
		}

		log.Debug("Failed to create user", sl.Err(err))

		return fmt.Errorf("failed to create user: %w", err)
	}

	log.Debug("User created successfully")
	return nil
}

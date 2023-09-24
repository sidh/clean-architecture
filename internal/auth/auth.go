package auth

import (
	"context"
	"errors"

	"github.com/sidh/clean-architecture/internal/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Auth interface {
	Permission(ctx context.Context, user string, action models.AuthAction) (bool, error)
}

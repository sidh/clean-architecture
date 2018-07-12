package auth

import (
	"context"

	"github.com/pkg/errors"

	"github.com/sidhman/clean-architecture/pkg/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Auth interface {
	Permission(ctx context.Context, user string, action models.AuthAction) (bool, error)
}

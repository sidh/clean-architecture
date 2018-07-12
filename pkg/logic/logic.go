package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/sidhman/clean-architecture/pkg/models"
)

var (
	// ErrActionDenied user does not have permissions for requested action
	ErrActionDenied = errors.New("action denied")
	// ErrAuthFailed failed to authenticate
	ErrAuthFailed = errors.New("failed to authenticate")
	// ErrActionFailed failed to perform action
	ErrActionFailed = errors.New("action failed")
)

// Core is the main interface for service logic
type Core interface {
	// Store checks user permissions and store key/value pair
	Store(ctx context.Context, user, key string, value models.Value) error
	// Load checks user permissions and loads value for a given key
	Load(ctx context.Context, user, key string) (models.Value, error)
}

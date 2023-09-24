package safe

import (
	"context"
	"errors"

	"github.com/sidh/clean-architecture/internal/auth"
	"github.com/sidh/clean-architecture/internal/logic"
	"github.com/sidh/clean-architecture/internal/models"
	"github.com/sidh/clean-architecture/internal/storage"
)

var _ logic.Core = &Core{}

// Core implements main interface for service logic
type Core struct {
	authorizer auth.Auth
	store      storage.Storage
}

// New constructs new logic
func New(a auth.Auth, s storage.Storage) *Core {
	return &Core{
		authorizer: a,
		store:      s,
	}
}

func (c *Core) permission(ctx context.Context, user string, action models.AuthAction) error {
	ok, err := c.authorizer.Permission(ctx, user, action)
	if err != nil {
		if err == auth.ErrUserNotFound {
			return logic.ErrActionDenied
		}

		return logic.ErrAuthFailed
	}

	if !ok {
		return logic.ErrActionDenied
	}

	return nil
}

// Store checks user permissions and store key/value pair
func (c *Core) Store(ctx context.Context, user, key string, value models.Value) error {
	if err := c.permission(ctx, user, models.AuthActionStore); err != nil {
		return err
	}

	if err := c.store.Store(ctx, key, value); err != nil {
		return logic.ErrActionFailed
	}

	return nil
}

// Load checks user permissions and loads value for a given key
func (c *Core) Load(ctx context.Context, user, key string) (models.Value, error) {
	if err := c.permission(ctx, user, models.AuthActionLoad); err != nil {
		return models.Value{}, err
	}

	value, err := c.store.Load(ctx, key)
	if err != nil {
		if errors.Is(err, storage.ErrKeyNotFound) {
			return models.Value{}, logic.ErrKeyNotFound
		}

		return models.Value{}, logic.ErrActionFailed
	}

	return value, nil
}

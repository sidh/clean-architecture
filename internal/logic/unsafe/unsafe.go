package unsafe

import (
	"context"
	"errors"
	"fmt"

	"github.com/sidh/clean-architecture/internal/logic"
	"github.com/sidh/clean-architecture/internal/models"
	"github.com/sidh/clean-architecture/internal/storage"
)

var _ logic.Core = &Core{}

// Core implements main interface for service logic
type Core struct {
	store storage.Storage
}

// New constructs new logic
func New(s storage.Storage) *Core {
	return &Core{
		store: s,
	}
}

// Store key/value pair
func (c *Core) Store(ctx context.Context, user, key string, value models.Value) error {
	fmt.Println("Unsafe module is active. Storing data.")

	if err := c.store.Store(ctx, key, value); err != nil {
		return logic.ErrActionFailed
	}

	return nil
}

// Load  value for a given key
func (c *Core) Load(ctx context.Context, user, key string) (models.Value, error) {
	fmt.Println("Unsafe module is active. Loading data.")

	value, err := c.store.Load(ctx, key)
	if err != nil {
		if errors.Is(err, storage.ErrKeyNotFound) {
			return models.Value{}, logic.ErrKeyNotFound
		}

		return models.Value{}, logic.ErrActionFailed
	}

	return value, nil
}

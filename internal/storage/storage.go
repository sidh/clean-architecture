package storage

import (
	"context"
	"errors"

	"github.com/sidh/clean-architecture/internal/models"
)

var (
	// ErrNotAvailable storage cannot perform requested operation
	ErrNotAvailable = errors.New("storage not available")
	// ErrKeyNotFound tried to load value for nonexistent key
	ErrKeyNotFound = errors.New("key not found")
	// ErrInvalidValue value is invalid, cannot either marshal or unmarshal
	ErrInvalidValue = errors.New("invalid value")
)

// Storage is an interface for key/value storage
type Storage interface {
	// Store stores key/value pair
	Store(ctx context.Context, key string, value models.Value) error
	// Load loads value for given key
	Load(ctx context.Context, key string) (models.Value, error)
}

package fs

import (
	"context"

	"github.com/sidh/clean-architecture/internal/models"
	"github.com/sidh/clean-architecture/internal/storage"
)

var _ storage.Storage = &Storage{}

// Storage implements Storage interface
type Storage struct {
	storage map[string]models.Value
}

// New constructs Storage
func New() *Storage {
	return &Storage{storage: make(map[string]models.Value)}
}

// Store stores key/value pair
func (s *Storage) Store(ctx context.Context, key string, value models.Value) error {
	s.storage[key] = value
	return nil
}

// Load loads value for given key
func (s *Storage) Load(ctx context.Context, key string) (models.Value, error) {
	value, ok := s.storage[key]
	if !ok {
		return models.Value{}, storage.ErrKeyNotFound
	}

	return value, nil
}

package fs

import (
	"context"
	"fmt"
	"os"

	"github.com/sidh/clean-architecture/internal/models"
	"github.com/sidh/clean-architecture/internal/storage"
)

var _ storage.Storage = &Storage{}

// Storage implements Storage interface
type Storage struct {
	path string
}

// New constructs Storage
func New(path string) *Storage {
	return &Storage{path: path}
}

// Store stores key/value pair
func (s *Storage) Store(ctx context.Context, key string, value models.Value) error {
	sv := fromValueModel(value)
	data, err := sv.Marshal()
	if err != nil {
		return storage.ErrInvalidValue
	}

	path := formatFilePath(s.path, key)
	if err := os.WriteFile(path, data, 0600); err != nil {
		if os.IsPermission(err) {
			fmt.Printf("Failed to access file at '%s': %s", path, err)
		} else {
			fmt.Printf("Failed to read file at '%s': %s", path, err)
		}

		return storage.ErrNotAvailable
	}

	return nil
}

// Load loads value for given key
func (s *Storage) Load(ctx context.Context, key string) (models.Value, error) {
	path := formatFilePath(s.path, key)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Printf("Failed to access file at '%s': %s", path, err)
		} else {
			fmt.Printf("Failed to read file at '%s': %s", path, err)
		}

		return models.Value{}, storage.ErrKeyNotFound
	}

	var sv storedValue
	if err = sv.Unmarshal(data); err != nil {
		return models.Value{}, storage.ErrInvalidValue
	}

	return toValueModel(sv), nil
}

func formatFilePath(path, key string) string {
	return path + "/" + key + ".storage"
}

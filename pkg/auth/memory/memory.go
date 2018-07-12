package memory

import (
	"context"

	"github.com/sidhman/clean-architecture/pkg/auth"
	"github.com/sidhman/clean-architecture/pkg/models"
)

type Permission struct {
	Store bool
	Load  bool
}

var _ auth.Auth = &Auth{}

type Auth struct {
	users map[string]Permission
}

func New(users map[string]Permission) *Auth {
	return &Auth{users: users}
}

func (a *Auth) Permission(ctx context.Context, user string, action models.AuthAction) (bool, error) {
	p, ok := a.users[user]
	if !ok {
		return false, auth.ErrUserNotFound
	}

	switch action {
	case models.AuthActionStore:
		return p.Store, nil
	case models.AuthActionLoad:
		return p.Load, nil
	default:
		return false, nil
	}
}

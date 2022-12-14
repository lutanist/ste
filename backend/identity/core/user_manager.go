package core

import (
	"errors"

	"github.com/lutanist/ste/backend/identity/ent"
)

type UserManager struct {
	client *ent.Client
}

func (m *UserManager) Create(user *ent.User, password string) (*IdentityResult, error) {
	return IdentityResultSuccess(), errors.New("not implemented")
}

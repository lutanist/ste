package core

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"github.com/lutanist/ste/backend/identity/ent"
	"github.com/lutanist/ste/backend/identity/ent/user"
)

type UserManager struct {
	options IdentityOptions
	client  *ent.Client

	passwordValidators []PasswordValidator
	passwordHasher     PasswordHasher
	userValidators     []func(*UserManager, *ent.User) *IdentityResult
}

func NewUserManager(client *ent.Client, validatePassword []PasswordValidator) *UserManager {
	um := &UserManager{
		options: IdentityOptions{},
		client:  client,

		passwordValidators: []PasswordValidator{},
		passwordHasher:     new(BcryptPasswordHahser),
		userValidators: []func(*UserManager, *ent.User) *IdentityResult{
			ValidateUser,
		},
	}

	if len(validatePassword) > 0 {
		um.passwordValidators = append(um.passwordValidators, validatePassword...)
	} else {
		um.passwordValidators = append(um.passwordValidators, EntoryPasswordValidator)
	}

	return um
}

func (m *UserManager) CreateWithPassword(ctx context.Context, user *ent.User, password string) (*IdentityResult, error) {
	// FIXME: 여기서 이렇게 설정하면 필드가 추가될때마다 여기를 바꿔야 한다. 이건 이상하다. 그렇다고 entgo를 쓰는 입장에서
	//   마땅히 repository 패턴을 쓸 방법이 없다.
	r, err := m.UpdatePasswordHash(user, password, true)
	if err != nil || !r.Succeeded {
		return r, err
	}

	return m.Create(ctx, user)
}

func (m *UserManager) Create(ctx context.Context, user *ent.User) (*IdentityResult, error) {
	if err := m.UpdateSecurityStampInternal(user); err != nil {
		return nil, err
	}

	if r := m.ValidateUser(user); !r.Succeeded {
		return r, nil
	}

	if m.options.Lockout.AllowedForNewUsers {
		user.LockoutEnabled = true
	}

	if _, err := m.client.User.Create().
		SetName(user.Name).
		SetUsername(user.Username).
		SetNillablePasswordHash(user.PasswordHash).
		SetNillableSecurityStamp(user.SecurityStamp).
		SetLockoutEnabled(user.LockoutEnabled).
		Save(ctx); err != nil {
		return nil, err
	}
	return IdentityResultSuccess(), errors.New("not implemented")
}

// validate security stamp, username, email
func (m *UserManager) ValidateUser(user *ent.User) *IdentityResult {
	errs := []IdentityError{}
	for _, v := range m.userValidators {
		r := v(m, user)
		if !r.Succeeded {
			errs = append(errs, r.Errors...)
		}
	}
	if len(errs) > 0 {
		return IdentityResultFailed(errs...)
	}
	return IdentityResultSuccess()
}

func (m *UserManager) UpdatePasswordHash(user *ent.User, newPassword string, validatePassword bool) (*IdentityResult, error) {
	if validatePassword {
		result := m.ValidatePassword(user, newPassword)
		if !result.Succeeded {
			return result, nil
		}
	}

	hash, err := m.passwordHasher.HashPassword(newPassword)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = &hash
	if err = m.UpdateSecurityStampInternal(user); err != nil {
		return nil, err
	}
	return IdentityResultSuccess(), nil
}

func (m *UserManager) ValidatePassword(user *ent.User, password string) *IdentityResult {
	errors := []IdentityError{}
	for _, v := range m.passwordValidators {
		r := v(password)
		if !r.Succeeded {
			errors = append(errors, r.Errors...)
		}
	}
	if len(errors) > 0 {
		return IdentityResultFailed(errors...)
	}
	return IdentityResultSuccess()
}

func (m *UserManager) UpdateSecurityStampInternal(user *ent.User) error {
	stamp, err := NewSecurityStamp()
	if err != nil {
		return err
	}
	user.SecurityStamp = &stamp
	return nil
}

func (m *UserManager) FindByUsername(ctx context.Context, username string) (*ent.User, error) {
	return m.client.User.Query().Where(user.Username(username)).First(ctx)
}

func (m *UserManager) CheckUsernameDuplicate(ctx context.Context, u *ent.User) (bool, error) {
	return m.client.User.Query().Where(user.And(user.UsernameEQ(u.Username), user.IDNEQ(u.ID))).Exist(ctx)
}

func (m *UserManager) CheckEmailDuplicate(ctx context.Context, u *ent.User) (bool, error) {
	return m.client.User.Query().Where(user.And(user.EmailEQ(u.Email), user.IDNEQ(u.ID))).Exist(ctx)
}

func NewSecurityStamp() (string, error) {
	stamp := make([]byte, 20)
	if _, err := rand.Read(stamp); err != nil {
		return "", err
	}
	return base32.StdEncoding.EncodeToString(stamp), nil
}

func ValidateUser(um *UserManager, user *ent.User) *IdentityResult {
	errs := []IdentityError{}

	nerrs := validateUsername(um, user)
	errs = append(errs, nerrs...)

	err := validateEmail(um, user)
	if err.Code != "" {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return IdentityResultFailed(errs...)
	}
	return IdentityResultSuccess()
}

// TODO: i18n
func validateUsername(um *UserManager, user *ent.User) []IdentityError {
	errors := []IdentityError{}
	username := strings.TrimSpace(user.Username)
	if username == "" || len(username) != len(user.Username) {
		errors = append(errors,
			IdentityError{Code: "InvalidUserName", Description: fmt.Sprintf("InvalidUserName %s", user.Username)})
	}

	if um.options.User.AllowedUserNameCharacters != "" {
		contained := false
		for _, c := range username {
			if !strings.ContainsRune(um.options.User.AllowedUserNameCharacters, c) {
				contained = true
				break
			}
		}

		if contained {
			errors = append(errors,
				IdentityError{Code: "InvalidUserName", Description: fmt.Sprintf("InvalidUserName %s", user.Username)})
		}
	}

	exists, err := um.CheckUsernameDuplicate(context.Background(), user)
	if err != nil {
		errors = append(errors, IdentityError{Code: "System", Description: err.Error()})
	}
	if exists {
		errors = append(errors, IdentityError{Code: "DuplicateUserName", Description: fmt.Sprintf("DuplicateUserName %s", user.Username)})
	}
	return errors
}

func validateEmail(um *UserManager, user *ent.User) IdentityError {
	if email := strings.TrimSpace(user.Email); email == "" || len(email) != len(user.Email) {
		return IdentityError{Code: "InvalidEmail", Description: fmt.Sprintf("InvalidEmail %s", user.Email)}
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return IdentityError{Code: "InvalidEmail", Description: fmt.Sprintf("InvalidEmail %s", user.Email)}
	}

	exists, err := um.CheckEmailDuplicate(context.Background(), user)
	if err != nil {
		return IdentityError{Code: "System", Description: err.Error()}
	}
	if exists {
		return IdentityError{Code: "DuplicateEmail", Description: fmt.Sprintf("DuplicateEmail %s", user.Email)}
	}
	return IdentityError{}
}

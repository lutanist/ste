package core

import "golang.org/x/crypto/bcrypt"

type PasswordVerificationResult int

const (
	Failed PasswordVerificationResult = iota
	Success
	// password hash 알고리즘이 변경되서 rehash가 필요한 경우. 추후 구현
	SuccessRehashNeeded
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	VerifyHashedPassword(hashedPassword, providedPassword string) (PasswordVerificationResult, error)
}

type BcryptPasswordHahser struct {
}

func (h *BcryptPasswordHahser) HashPassword(password string) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hash), err
	}
}

func (h *BcryptPasswordHahser) VerifyHashedPassword(hashedPassword, providedPassword string) (PasswordVerificationResult, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return Failed, nil
		}
		// todo logging?
		return Failed, err
	}
	return Success, nil
}

package core

import passwordvalidator "github.com/wagslane/go-password-validator"

type PasswordValidator func(password string) *IdentityResult

// adator
func EntoryPasswordValidator(password string) *IdentityResult {
	// TODO: min entory move to config
	//   errors message support i18n
	err := passwordvalidator.Validate(password, 60)
	return IdentityResultFailed(IdentityError{Code: "Password", Description: err.Error()})
}

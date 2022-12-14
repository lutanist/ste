package core

type IdentityError struct {
}

type IdentityResult struct {
	Succeeded bool
	Errors    []IdentityError
}

var success *IdentityResult = &IdentityResult{
	Succeeded: true,
	Errors:    []IdentityError{},
}

func IdentityResultSuccess() *IdentityResult {
	return success
}

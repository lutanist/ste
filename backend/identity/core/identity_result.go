package core

type IdentityError struct {
	Code        string
	Description string
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

func IdentityResultFailed(erros ...IdentityError) *IdentityResult {
	return &IdentityResult{
		Succeeded: false,
		// TODO: 이거 빈 슬라이스로 초기화되는지 확인 필요
		Errors: erros,
	}
}

package core

type SignInManager struct {
	schemas AuthenticationSchemeProvider
}

func NewSignInManager(schemaProvider AuthenticationSchemeProvider) *SignInManager {
	sm := new(SignInManager)
	sm.schemas = schemaProvider
	return sm
}

func (s *SignInManager) GetExternalAuthenticationSchemes() []*AuthenticationScheme {
	n := []*AuthenticationScheme{}
	for _, s := range s.schemas.GetAllSchemes() {
		if s.DisplayName != "" {
			n = append(n, s)
		}
	}
	return n
}

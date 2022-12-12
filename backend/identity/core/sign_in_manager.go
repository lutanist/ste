package core

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type AuthenticationScheme struct {
	Name        string
	DisplayName string
}

type AuthenticationSchemeProvider interface {
	AddSchema(schema *AuthenticationScheme) error
	GetAllSchemes() []*AuthenticationScheme
}

type DefaultAuthenticationSchemeProvider struct {
	schemas map[string]*AuthenticationScheme
}

func (p *DefaultAuthenticationSchemeProvider) AddSchema(schema *AuthenticationScheme) error {
	if _, ok := p.schemas[schema.Name]; ok {
		return fmt.Errorf("schema already exists: %s", schema.Name)
	}
	p.schemas[schema.Name] = schema
	return nil
}

func (p *DefaultAuthenticationSchemeProvider) GetAllSchemes() []*AuthenticationScheme {
	return maps.Values(p.schemas)
}

var (
	defaultAuthenticationSchemeProvider *DefaultAuthenticationSchemeProvider
)

func init() {
	defaultAuthenticationSchemeProvider = new(DefaultAuthenticationSchemeProvider)
	defaultAuthenticationSchemeProvider.schemas = map[string]*AuthenticationScheme{}
}

func GetDefaultAuthenticationSchemeProvider() *DefaultAuthenticationSchemeProvider {
	return defaultAuthenticationSchemeProvider
}

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

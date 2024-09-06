package irepository

import (
	"auto-monitoring/internal/domain"
)

type ITokenRepository interface {
	Generate(domain.TokenClaims) (string, error)
	Verify(string) (domain.TokenClaims, error)
	Delete(string) error
}

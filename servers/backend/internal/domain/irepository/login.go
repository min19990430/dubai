package irepository

import (
	"auto-monitoring/internal/domain"
)

type ILoginRepository interface {
	Varify(username, password string) (domain.User, error)
	SucceedHandle(domain.User) error
	ErrorHandle(domain.User, error) error
}

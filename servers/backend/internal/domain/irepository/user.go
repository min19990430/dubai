package irepository

import "auto-monitoring/internal/domain"

type IUserRepository interface {
	List(domain.User) ([]domain.User, error)
	ListWithRoot(domain.User, bool) ([]domain.User, error)
	Find(domain.User) (domain.User, error)

	UpdatePassword(domain.User, string) error
}

type IUserAuthRepository interface {
	FindByUser(domain.User) (domain.UserWithAuthority, error)
}

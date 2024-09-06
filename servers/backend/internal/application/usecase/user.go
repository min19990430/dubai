package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type UserUsecase struct {
	user irepository.IUserRepository
}

func NewUserUsecase(user irepository.IUserRepository) *UserUsecase {
	return &UserUsecase{user: user}
}

func (uu *UserUsecase) List(user domain.User) ([]domain.User, error) {
	return uu.user.List(user)
}

func (uu *UserUsecase) ListWithRoot(user domain.User, isRoot bool) ([]domain.User, error) {
	return uu.user.ListWithRoot(user, isRoot)
}

func (uu *UserUsecase) Find(user domain.User) (domain.User, error) {
	return uu.user.Find(user)
}

func (uu *UserUsecase) UpdatePassword(user domain.User, password string) error {
	return uu.user.UpdatePassword(user, password)
}

type UserAuthUsecase struct {
	userAuth irepository.IUserAuthRepository
}

func NewUserAuthUsecase(userAuth irepository.IUserAuthRepository) *UserAuthUsecase {
	return &UserAuthUsecase{userAuth: userAuth}
}

func (uau *UserAuthUsecase) FindByUser(user domain.User) (domain.UserWithAuthority, error) {
	return uau.userAuth.FindByUser(user)
}

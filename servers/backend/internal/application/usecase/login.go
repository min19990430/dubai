package usecase

import (
	"errors"

	"auto-monitoring/internal/domain/irepository"
)

type LoginUsecase struct {
	login irepository.ILoginRepository

	tokenUsecase TokenUsecase
}

func NewLoginUsecase(
	login irepository.ILoginRepository,
	tokenUsecase TokenUsecase,
) *LoginUsecase {
	return &LoginUsecase{
		login:        login,
		tokenUsecase: tokenUsecase,
	}
}

func (lu *LoginUsecase) Login(username, password string) (string, error) {
	user, varifyErr := lu.login.Varify(username, password)
	if varifyErr != nil {
		handleErr := lu.login.ErrorHandle(user, varifyErr)
		if handleErr != nil {
			return "", handleErr
		}
		return "", errors.New("username or password is incorrect")
	}

	handleErr := lu.login.SucceedHandle(user)
	if handleErr != nil {
		return "", handleErr
	}

	return lu.tokenUsecase.NewWithUser(user)
}

func (lu *LoginUsecase) Logout(userUUID string) error {
	return lu.tokenUsecase.Delete(userUUID)
}

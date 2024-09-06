package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type TokenUsecase struct {
	token irepository.ITokenRepository

	userAuth UserAuthUsecase
}

func NewTokenUsecase(
	token irepository.ITokenRepository,
	userAuth UserAuthUsecase,
) *TokenUsecase {
	return &TokenUsecase{
		token:    token,
		userAuth: userAuth,
	}
}

func (tu *TokenUsecase) NewWithUser(user domain.User) (string, error) {
	userAuth, findErr := tu.userAuth.FindByUser(domain.User{UUID: user.UUID})
	if findErr != nil {
		return "", findErr
	}

	tokenClaims := domain.TokenClaims{
		UserUUID:    user.UUID,
		Username:    user.Username,
		FullName:    user.FullName,
		NickName:    user.NickName,
		AuthorityID: userAuth.Authority.ID,
	}

	token, generateErr := tu.token.Generate(tokenClaims)
	if generateErr != nil {
		return "", generateErr
	}
	return token, nil
}

func (tu *TokenUsecase) Verify(token string) (domain.TokenClaims, error) {
	return tu.token.Verify(token)
}

func (tu *TokenUsecase) Delete(userUUID string) error {
	return tu.token.Delete(userUUID)
}

package model

import (
	"github.com/golang-jwt/jwt/v4"

	"auto-monitoring/internal/domain"
)

type TokenClaims struct {
	UserUUID    string `json:"user_uuid"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	NickName    string `json:"nick_name"`
	AuthorityID string `json:"authority_id"`

	jwt.StandardClaims `json:"standard_claims"`
}

func (tc TokenClaims) ToDomain() domain.TokenClaims {
	return domain.TokenClaims{
		UserUUID:    tc.UserUUID,
		Username:    tc.Username,
		FullName:    tc.FullName,
		NickName:    tc.NickName,
		AuthorityID: tc.AuthorityID,
	}
}

func (tc TokenClaims) FromDomain(tokenClaims domain.TokenClaims) TokenClaims {
	return TokenClaims{
		UserUUID:    tokenClaims.UserUUID,
		Username:    tokenClaims.Username,
		FullName:    tokenClaims.FullName,
		NickName:    tokenClaims.NickName,
		AuthorityID: tokenClaims.AuthorityID,
	}
}

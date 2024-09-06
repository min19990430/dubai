package repository

import (
	"context"
	"crypto/rand"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"

	"auto-monitoring/internal/adapter/redispool/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

const (
	jwtsaltPrefix = "jwtsalt:"
	saltSize      = 16
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type TokenRepository struct {
	redis *redis.Client

	expiration time.Duration
}

func NewTokenRepository(redis *redis.Client, expiration time.Duration) irepository.ITokenRepository {
	return &TokenRepository{
		redis:      redis,
		expiration: expiration,
	}
}

func (TokenRepository) groupName(userUUID string) string {
	return jwtsaltPrefix + userUUID
}

func (tr *TokenRepository) Generate(tokenClaims domain.TokenClaims) (string, error) {
	tokenClaimsPO := model.TokenClaims{}.FromDomain(tokenClaims)
	tokenClaimsPO.ExpiresAt = time.Now().Add(tr.expiration).Unix()

	salt := tr.generateRandomSalt(saltSize)

	_, setErr := tr.redis.Set(context.Background(), tr.groupName(tokenClaimsPO.UserUUID), string(salt), tr.expiration).Result()
	if setErr != nil {
		return "", setErr
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaimsPO).SignedString(append([]byte(tokenClaimsPO.UserUUID), salt...))
}

func (TokenRepository) generateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	return salt
}

func (tr *TokenRepository) Verify(token string) (domain.TokenClaims, error) {
	unverifiedToken, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return domain.TokenClaims{}, err
	}
	mapClaims, ok := unverifiedToken.Claims.(jwt.MapClaims)
	if !ok {
		return domain.TokenClaims{}, ErrInvalidToken
	}

	userUUID, ok := mapClaims["user_uuid"].(string)
	if !ok {
		return domain.TokenClaims{}, ErrInvalidToken
	}

	salt, getErr := tr.redis.Get(context.Background(), tr.groupName(userUUID)).Result()
	if getErr != nil {
		return domain.TokenClaims{}, getErr
	}

	key := []byte(userUUID + salt)

	tokenClaims, parseErr := jwt.ParseWithClaims(token, &model.TokenClaims{},
		func(_ *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if parseErr != nil {
		return domain.TokenClaims{}, parseErr
	}

	if tokenClaims == nil {
		return domain.TokenClaims{}, ErrInvalidToken
	}

	if !tokenClaims.Valid {
		return domain.TokenClaims{}, ErrInvalidToken
	}

	tokenClaimsResult, ok := tokenClaims.Claims.(*model.TokenClaims)
	if !ok {
		return domain.TokenClaims{}, ErrInvalidToken
	}

	return tokenClaimsResult.ToDomain(), nil
}

func (tr *TokenRepository) Delete(userUUID string) error {
	return tr.redis.Del(context.Background(), tr.groupName(userUUID)).Err()
}

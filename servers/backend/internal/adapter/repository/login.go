package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type LoginRepository struct {
	gorm *gorm.DB
}

func NewLoginRepository(gorm *gorm.DB) irepository.ILoginRepository {
	return &LoginRepository{
		gorm: gorm,
	}
}

func (r *LoginRepository) Varify(username, password string) (domain.User, error) {
	var userPO model.User
	firstErr := r.gorm.Table("user").Where("username = ?", username).First(&userPO).Error
	if firstErr != nil {
		return domain.User{}, firstErr
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(userPO.Password), []byte(password))
	if compareErr != nil {
		return domain.User{}, compareErr
	}

	userPO.Password = ""

	return userPO.ToDomain(), nil
}

func (r *LoginRepository) SucceedHandle(_ domain.User) error {
	return nil
}

func (r *LoginRepository) ErrorHandle(_ domain.User, _ error) error {
	return nil
}

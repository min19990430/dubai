package repository

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type UserRepository struct {
	gorm *gorm.DB
}

func NewUserRepository(gorm *gorm.DB) irepository.IUserRepository {
	return &UserRepository{gorm: gorm}
}

func (u *UserRepository) List(user domain.User) ([]domain.User, error) {
	userWherePO := model.User{}.FromDomain(user)
	userPOs := []model.User{}

	err := u.gorm.Where(userWherePO).Find(&userPOs).Error
	if err != nil {
		return []domain.User{}, err
	}

	var users []domain.User
	for _, userPO := range userPOs {
		users = append(users, userPO.ToDomain())
	}
	return users, nil
}

func (u *UserRepository) ListWithRoot(user domain.User, isRoot bool) ([]domain.User, error) {
	userWherePO := model.User{}.FromDomain(user)
	userPOs := []model.User{}

	query := u.gorm.Where(userWherePO)
	if !isRoot {
		subQuery := u.gorm.Table("authority").Select("user_uuid").Joins("JOIN user_authority ON authority_uuid = uuid").Where("id != ?", "root")
		query = query.Where("uuid in (?)", subQuery)
	}

	err := query.Find(&userPOs).Error
	if err != nil {
		return []domain.User{}, err
	}

	var users []domain.User
	for _, userPO := range userPOs {
		users = append(users, userPO.ToDomain())
	}
	return users, nil
}

func (u *UserRepository) Find(user domain.User) (domain.User, error) {
	userWherePO := model.User{}.FromDomain(user)
	userPO := model.User{}

	err := u.gorm.Where(userWherePO).First(&userPO).Error
	if err != nil {
		return domain.User{}, err
	}

	return userPO.ToDomain(), nil
}

func (u *UserRepository) UpdatePassword(user domain.User, password string) error {
	// 將明碼密碼轉換成雜湊密碼
	userPasswordHash, generateErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if generateErr != nil {
		return generateErr
	}

	// 更新密碼
	updateColumn := map[string]interface{}{
		"password": string(userPasswordHash),
	}

	update := u.gorm.Table("user").Where("uuid = ?", user.UUID).Omit("uuid").Updates(updateColumn)

	if err := update.Error; err != nil {
		return err
	}
	if update.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

type UserAuthRepository struct {
	gorm *gorm.DB
}

func NewUserAuthRepository(gorm *gorm.DB) irepository.IUserAuthRepository {
	return &UserAuthRepository{gorm: gorm}
}

func (u *UserAuthRepository) FindByUser(user domain.User) (domain.UserWithAuthority, error) {
	userWherePO := model.User{}.FromDomain(user)
	userPO := model.User{}

	err := u.gorm.Where(userWherePO).First(&userPO).Error
	if err != nil {
		return domain.UserWithAuthority{}, err
	}

	var authorityPO model.Authority

	subQuery := u.gorm.Table("user_authority").Select("authority_uuid").Where("user_uuid = ?", userPO.UUID).Limit(1)
	err = u.gorm.Where("uuid like (?)", subQuery).Find(&authorityPO).Error
	if err != nil {
		return domain.UserWithAuthority{}, err
	}

	var userWithAuthority model.UserWithAuthority
	userWithAuthority.User = userPO
	userWithAuthority.Authority = authorityPO

	return userWithAuthority.ToDomain(), nil
}

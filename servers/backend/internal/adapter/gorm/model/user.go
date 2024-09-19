package model

import (
	"time"

	"auto-monitoring/internal/domain"

	"gorm.io/gorm"
)

type User struct {
	UUID        string `gorm:"column:uuid;primary_key;type:char(36)"`
	IsEnable    bool   `gorm:"column:is_enable;not null;default:true;type:tinyint(1)" json:"is_enable"`
	Username    string `gorm:"column:username;not null;unique;type:varchar(128)" json:"username"`
	Password    string `gorm:"column:password;not null;type:varchar(128)" json:"password,omitempty"`
	FullName    string `gorm:"column:full_name;not null;type:varchar(100)" json:"full_name"`
	NickName    string `gorm:"column:nick_name;type:varchar(100)" json:"nick_name"`
	Phone       string `gorm:"column:phone;type:varchar(10)" json:"phone"`
	Email       string `gorm:"column:email;type:varchar(200)" json:"email"`
	Description string `gorm:"column:description;type:text" json:"description"`

	CreatedAt time.Time      `gorm:"autoCreateTime;column:created_at;not null;type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (User) TableName() string {
	return "user"
}

func (u User) ToDomain() domain.User {
	return domain.User{
		UUID:        u.UUID,
		IsEnable:    u.IsEnable,
		Username:    u.Username,
		Password:    u.Password,
		FullName:    u.FullName,
		NickName:    u.NickName,
		Phone:       u.Phone,
		Email:       u.Email,
		Description: u.Description,
	}
}

func (u User) FromDomain(user domain.User) User {
	return User{
		UUID:        user.UUID,
		IsEnable:    u.IsEnable,
		Username:    user.Username,
		Password:    user.Password,
		FullName:    user.FullName,
		NickName:    user.NickName,
		Phone:       user.Phone,
		Email:       user.Email,
		Description: user.Description,
	}
}

type UserWithAuthority struct {
	User

	Authority Authority `gorm:"many2many:user_authority;foreignKey:UUID;joinForeignKey:user_uuid;References:UUID;JoinReferences:authority_uuid" json:"authorities"`
}

func (UserWithAuthority) TableName() string {
	return "user"
}

func (u UserWithAuthority) ToDomain() domain.UserWithAuthority {
	return domain.UserWithAuthority{
		User:      u.User.ToDomain(),
		Authority: u.Authority.ToDomain(),
	}
}

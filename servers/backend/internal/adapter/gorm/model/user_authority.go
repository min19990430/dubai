package model

type UserAuthority struct {
	UserUUID      string `gorm:"column:user_uuid;primary_key;type:varchar(36)"`
	AuthorityUUID string `gorm:"column:authority_uuid;primary_key;type:varchar(36)"`
}

func (UserAuthority) TableName() string {
	return "user_authority"
}

package domain

type Authority struct {
	UUID        string `gorm:"column:uuid;primary_key;type:varchar(36)"`
	ID          string `gorm:"column:id;not null;unique;type:varchar(128)" json:"id"`
	Name        string `gorm:"column:name;not null;unique;type:varchar(128)" json:"name"`
	Description string `gorm:"column:description;type:text" json:"description"`
}

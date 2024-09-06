package model

import (
	"time"

	"auto-monitoring/internal/domain"
)

type AlarmRecord struct {
	ID               uint      `gorm:"column:id;not null;primary_key;auto_increment;type:bigint(20)" json:"-"`
	AlarmSettingUUID string    `gorm:"column:alarm_setting_uuid;type:char(36)" json:"alarm_setting_uuid"`
	Name             string    `gorm:"column:name;not null;index;type:varchar(45)" json:"device_name"`
	FullName         string    `gorm:"column:full_name;not null;type:varchar(45)" json:"full_name"`
	OccurTime        time.Time `gorm:"column:occur_time;not null;index;type:datetime" json:"occur_time"`
	Content          string    `gorm:"column:content;not null;type:text" json:"content"`
}

func (AlarmRecord) TableName() string {
	return "alarm_record"
}

func (a AlarmRecord) FromDomain(domain domain.AlarmRecord) AlarmRecord {
	return AlarmRecord{
		ID:               domain.ID,
		AlarmSettingUUID: domain.AlarmSettingUUID,
		Name:             domain.Name,
		FullName:         domain.FullName,
		OccurTime:        domain.OccurTime,
		Content:          domain.Content,
	}
}

func (a AlarmRecord) ToDomain() domain.AlarmRecord {
	return domain.AlarmRecord{
		ID:               a.ID,
		AlarmSettingUUID: a.AlarmSettingUUID,
		Name:             a.Name,
		FullName:         a.FullName,
		OccurTime:        a.OccurTime,
		Content:          a.Content,
	}
}

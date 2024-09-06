package domain

import "time"

type DeviceAlarmRecord struct {
	ID                     uint      `json:"-"`
	DeviceAlarmSettingUUID string    `json:"device_alarm_setting_uuid"`
	Name                   string    `json:"name"`
	FullName               string    `json:"full_name"`
	OccurTime              time.Time `json:"occur_time"`
	Content                string    `json:"content"`
}

package domain

import "time"

type DeviceAlarmSetting struct {
	UUID       string `json:"uuid"`
	DeviceUUID string `json:"device_uuid"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	Priority   int    `json:"priority"`

	BooleanExpression string `json:"boolean_expression"`

	IsEnable    bool `json:"is_enable"`
	IsNotify    bool `json:"is_notify"`
	IsActivated bool `json:"is_activated"`

	LastOccurTime *time.Time `json:"last_occur_time"`

	ContentSetting *string `json:"content_setting"`
}

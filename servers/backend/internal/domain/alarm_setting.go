package domain

import (
	"time"
)

type AlarmSetting struct {
	UUID                 string `json:"uuid,omitempty"`
	PhysicalQuantityUUID string `json:"physical_quantity_uuid"`
	Name                 string `json:"name"`
	FullName             string `json:"full_name"`
	Priority             int    `json:"priority"`

	IsEnable    bool `json:"is_enable"`
	IsNotify    bool `json:"is_notify"`
	IsActivated bool `json:"is_activated"`

	BooleanExpression string  `json:"boolean_expression,omitempty"`
	LastOccurValue    float64 `json:"last_occur_value"`

	LastAlarmOccurTime  *time.Time `json:"last_alarm_occur_time,omitempty"`
	AlarmContentSetting *string    `json:"alarm_content_setting,omitempty"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
}

type PQAlarmSettings struct {
	PhysicalQuantity `json:"physical_quantity"`
	AlarmSettings    []AlarmSetting `json:"alarm_settings"`
}

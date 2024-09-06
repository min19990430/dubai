package domain

import "time"

type AlarmRecord struct {
	ID               uint      `json:"-"`
	AlarmSettingUUID string    `json:"alarm_setting_uuid"`
	Name             string    `json:"name"`
	FullName         string    `json:"full_name"`
	OccurTime        time.Time `json:"occur_time"`
	Content          string    `json:"content"`
}

type AlarmRecordDetail struct {
	ID               uint      `json:"-"`
	AlarmSettingUUID string    `json:"alarm_setting_uuid"`
	Name             string    `json:"name"`
	FullName         string    `json:"full_name"`
	OccurTime        time.Time `json:"occur_time"`
	Content          string    `json:"content"`

	PhysicalQuantityUUID     string `json:"physical_quantity_uuid"`
	PhysicalQuantityName     string `json:"physical_quantity_name"`
	PhysicalQuantityFullName string `json:"physical_quantity_full_name"`

	DeviceUUID string `json:"device_uuid"`
	DeviceID   string `json:"device_id"`
	DeviceName string `json:"device_name"`

	StationUUID string `json:"station_uuid"`
	StationID   string `json:"station_id"`
	StationName string `json:"station_name"`
}

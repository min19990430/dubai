package convert

import "auto-monitoring/internal/domain"

type AlarmRecordCollection struct{}

func (AlarmRecordCollection) FromAlarmRecord(ar domain.AlarmRecord) domain.AlarmRecordCollection {
	return domain.AlarmRecordCollection{
		Name:      ar.Name,
		FullName:  ar.FullName,
		OccurTime: ar.OccurTime,
		Content:   ar.Content,
	}
}

func (AlarmRecordCollection) FromDeviceAlarmRecord(dar domain.DeviceAlarmRecord) domain.AlarmRecordCollection {
	return domain.AlarmRecordCollection{
		Name:      dar.Name,
		FullName:  dar.FullName,
		OccurTime: dar.OccurTime,
		Content:   dar.Content,
	}
}

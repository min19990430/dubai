package usecase

import (
	"time"

	"github.com/samber/lo"

	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type AlarmRecordUsecase struct {
	alarmRecord irepository.IAlarmRecordRepository

	alarmSetting     irepository.IAlarmSettingRepository
	physicalQuantity irepository.IPhysicalQuantityRepository
	device           irepository.IDeviceRepository
	Station          irepository.IStationRepository
}

func NewAlarmRecordUsecase(
	alarmRecord irepository.IAlarmRecordRepository,
	alarmSetting irepository.IAlarmSettingRepository,
	physicalQuantity irepository.IPhysicalQuantityRepository,
	device irepository.IDeviceRepository,
	station irepository.IStationRepository) *AlarmRecordUsecase {
	return &AlarmRecordUsecase{
		alarmRecord: alarmRecord,

		alarmSetting:     alarmSetting,
		physicalQuantity: physicalQuantity,
		device:           device,
		Station:          station,
	}
}

func (aru *AlarmRecordUsecase) List(startTime, endTime time.Time, alarmRecord domain.AlarmRecord, reverse bool) ([]domain.AlarmRecord, error) {
	return aru.alarmRecord.List(startTime, endTime, alarmRecord, reverse)
}

func (aru *AlarmRecordUsecase) ListByDeviceUUID(startTime, endTime time.Time, deviceUUID string, reverse bool) ([]domain.AlarmRecord, error) {
	physicalQuantities, listErr := aru.physicalQuantity.List(
		domain.PhysicalQuantity{
			DeviceUUID: deviceUUID,
		})
	if listErr != nil {
		return nil, listErr
	}

	var physicalQuantityUUIDs = make([]string, len(physicalQuantities))
	for i, pq := range physicalQuantities {
		physicalQuantityUUIDs[i] = pq.UUID
	}

	alarmSettings, listErr := aru.alarmSetting.ListInPhysicalQuantityUUIDs(
		domain.AlarmSetting{},
		physicalQuantityUUIDs)
	if listErr != nil {
		return nil, listErr
	}

	alarmSettingUUIDs := make([]string, len(alarmSettings))
	for i, alarmSettingPO := range alarmSettings {
		alarmSettingUUIDs[i] = alarmSettingPO.UUID
	}

	return aru.alarmRecord.ListIn(startTime, endTime, alarmSettingUUIDs, reverse)
}

func (aru *AlarmRecordUsecase) ListByStationUUID(startTime, endTime time.Time, stationUUID string, reverse bool) ([]domain.AlarmRecord, error) {
	physicalQuantities, listErr := aru.physicalQuantity.List(domain.PhysicalQuantity{StationUUID: stationUUID})
	if listErr != nil {
		return nil, listErr
	}

	var physicalQuantityUUIDs = make([]string, len(physicalQuantities))
	for i, pq := range physicalQuantities {
		physicalQuantityUUIDs[i] = pq.UUID
	}

	alarmSettings, listErr := aru.alarmSetting.ListInPhysicalQuantityUUIDs(domain.AlarmSetting{}, physicalQuantityUUIDs)
	if listErr != nil {
		return nil, listErr
	}

	var alarmSettingUUIDs = make([]string, len(alarmSettings))
	for i, as := range alarmSettings {
		alarmSettingUUIDs[i] = as.UUID
	}

	return aru.alarmRecord.ListIn(startTime, endTime, alarmSettingUUIDs, reverse)
}

func (aru *AlarmRecordUsecase) ListDetail(startTime, endTime time.Time, alarmRecord domain.AlarmRecord, reverse bool) ([]domain.AlarmRecordDetail, error) {
	var alarmRecordDetails []domain.AlarmRecordDetail

	alarmRecords, listErr := aru.alarmRecord.List(startTime, endTime, alarmRecord, reverse)
	if listErr != nil {
		return nil, listErr
	}

	// 轉換成 AlarmRecordDetail
	for _, alarmRecord := range alarmRecords {
		alarmRecordDetails = append(alarmRecordDetails, domain.AlarmRecordDetail{
			ID:               alarmRecord.ID,
			AlarmSettingUUID: alarmRecord.AlarmSettingUUID,
			Name:             alarmRecord.Name,
			FullName:         alarmRecord.FullName,
			OccurTime:        alarmRecord.OccurTime,
			Content:          alarmRecord.Content,
		})
	}

	// 收尋/填充 alarmSetting
	alarmRecordDetails, alarmSettings, fillErr := aru.fillAlarmSetting(alarmRecordDetails, alarmRecords)
	if fillErr != nil {
		return nil, fillErr
	}

	// 收尋/填充 physicalQuantity
	alarmRecordDetails, physicalQuantities, fillErr := aru.fillPhysicalQuantity(alarmRecordDetails, alarmSettings)
	if fillErr != nil {
		return nil, fillErr
	}

	// 收尋/填充 device
	alarmRecordDetails, fillErr = aru.fillDevice(alarmRecordDetails, physicalQuantities)
	if fillErr != nil {
		return nil, fillErr
	}

	// 收尋/填充 station
	alarmRecordDetails, fillErr = aru.fillStation(alarmRecordDetails, physicalQuantities)
	if fillErr != nil {
		return nil, fillErr
	}

	return alarmRecordDetails, nil
}

func (aru *AlarmRecordUsecase) fillAlarmSetting(alarmRecordDetails []domain.AlarmRecordDetail, alarmRecords []domain.AlarmRecord) ([]domain.AlarmRecordDetail, []domain.AlarmSetting, error) {
	uniquAlarmRecord := lo.UniqBy(alarmRecords, func(ar domain.AlarmRecord) string {
		return ar.AlarmSettingUUID
	})

	alarmSettingUniqUUIDs := make([]string, len(uniquAlarmRecord))
	for i, ar := range uniquAlarmRecord {
		alarmSettingUniqUUIDs[i] = ar.AlarmSettingUUID
	}

	alarmSettings, listErr := aru.alarmSetting.ListIn(domain.AlarmSetting{}, alarmSettingUniqUUIDs)
	if listErr != nil {
		return nil, nil, listErr
	}

	alarmSettingMap := lo.KeyBy(alarmSettings, func(as domain.AlarmSetting) string {
		return as.UUID
	})

	for i, ar := range alarmRecordDetails {
		alarmRecordDetails[i].PhysicalQuantityUUID = alarmSettingMap[ar.AlarmSettingUUID].PhysicalQuantityUUID
	}

	return alarmRecordDetails, alarmSettings, nil
}

func (aru *AlarmRecordUsecase) fillPhysicalQuantity(alarmRecordDetails []domain.AlarmRecordDetail, alarmSettings []domain.AlarmSetting) ([]domain.AlarmRecordDetail, []domain.PhysicalQuantity, error) {
	uniquAlarmSettingPOs := lo.UniqBy(alarmSettings, func(as domain.AlarmSetting) string {
		return as.PhysicalQuantityUUID
	})

	physicalQuantityUniqUUIDs := make([]string, len(alarmSettings))
	for i, as := range uniquAlarmSettingPOs {
		physicalQuantityUniqUUIDs[i] = as.PhysicalQuantityUUID
	}

	physicalQuantities, listErr := aru.physicalQuantity.ListIn(domain.PhysicalQuantity{}, physicalQuantityUniqUUIDs)
	if listErr != nil {
		return nil, nil, listErr
	}

	physicalQuantityMap := lo.KeyBy(physicalQuantities, func(pq domain.PhysicalQuantity) string {
		return pq.UUID
	})

	for i, ar := range alarmRecordDetails {
		alarmRecordDetails[i].PhysicalQuantityName = physicalQuantityMap[ar.PhysicalQuantityUUID].Name
		alarmRecordDetails[i].PhysicalQuantityFullName = physicalQuantityMap[ar.PhysicalQuantityUUID].FullName
		alarmRecordDetails[i].DeviceUUID = physicalQuantityMap[ar.PhysicalQuantityUUID].DeviceUUID
		alarmRecordDetails[i].StationUUID = physicalQuantityMap[ar.PhysicalQuantityUUID].StationUUID
	}

	return alarmRecordDetails, physicalQuantities, nil
}

func (aru *AlarmRecordUsecase) fillDevice(alarmRecordDetails []domain.AlarmRecordDetail, physicalQuantities []domain.PhysicalQuantity) ([]domain.AlarmRecordDetail, error) {
	uniquPhysicalQuantityPOs := lo.UniqBy(physicalQuantities, func(pq domain.PhysicalQuantity) string {
		return pq.DeviceUUID
	})

	deviceUniqUUIDs := make([]string, len(uniquPhysicalQuantityPOs))
	for i, pq := range uniquPhysicalQuantityPOs {
		deviceUniqUUIDs[i] = pq.DeviceUUID
	}

	devices, listErr := aru.device.ListIn(domain.Device{}, deviceUniqUUIDs)
	if listErr != nil {
		return nil, listErr
	}

	deviceMap := lo.KeyBy(devices, func(d domain.Device) string {
		return d.UUID
	})

	for i, ar := range alarmRecordDetails {
		alarmRecordDetails[i].DeviceID = deviceMap[ar.DeviceUUID].ID
		alarmRecordDetails[i].DeviceName = deviceMap[ar.DeviceUUID].Name
	}

	return alarmRecordDetails, nil
}

func (aru *AlarmRecordUsecase) fillStation(alarmRecordDetails []domain.AlarmRecordDetail, physicalQuantities []domain.PhysicalQuantity) ([]domain.AlarmRecordDetail, error) {
	uniquPhysicalQuantityPOs := lo.UniqBy(physicalQuantities, func(pq domain.PhysicalQuantity) string {
		return pq.StationUUID
	})

	stationUniqUUIDs := make([]string, len(uniquPhysicalQuantityPOs))
	for i, d := range uniquPhysicalQuantityPOs {
		stationUniqUUIDs[i] = d.StationUUID
	}

	stations, listErr := aru.Station.ListIn(domain.Station{}, stationUniqUUIDs)
	if listErr != nil {
		return nil, listErr
	}

	stationMap := lo.KeyBy(stations, func(s domain.Station) string {
		return s.UUID
	})

	for i, ar := range alarmRecordDetails {
		alarmRecordDetails[i].StationID = stationMap[ar.StationUUID].ID
		alarmRecordDetails[i].StationName = stationMap[ar.StationUUID].Name
	}

	return alarmRecordDetails, nil
}

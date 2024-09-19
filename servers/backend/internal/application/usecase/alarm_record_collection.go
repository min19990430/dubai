package usecase

import (
	"sort"
	"time"

	"github.com/samber/lo"

	"auto-monitoring/internal/application/convert"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type AlarmRecordCollectionUsecase struct {
	alarmRecord        irepository.IAlarmRecordRepository
	deviceAlarmRecord  irepository.IDeviceAlarmRecordRepository
	deviceAlarmSetting irepository.IDeviceAlarmSettingRepository

	physicalQuantity irepository.IPhysicalQuantityRepository
	alarmSetting     irepository.IAlarmSettingRepository
}

func NewAlarmRecordCollectionUsecase(
	alarmRecord irepository.IAlarmRecordRepository,
	deviceAlarmRecord irepository.IDeviceAlarmRecordRepository,
	deviceAlarmSetting irepository.IDeviceAlarmSettingRepository,
	physicalQuantity irepository.IPhysicalQuantityRepository,
	alarmSetting irepository.IAlarmSettingRepository) *AlarmRecordCollectionUsecase {
	return &AlarmRecordCollectionUsecase{
		alarmRecord:        alarmRecord,
		deviceAlarmRecord:  deviceAlarmRecord,
		deviceAlarmSetting: deviceAlarmSetting,

		physicalQuantity: physicalQuantity,
		alarmSetting:     alarmSetting,
	}
}

func (arcu *AlarmRecordCollectionUsecase) ListByDeviceUUID(startTime, endTime time.Time, deviceUUID string, reverse bool) ([]domain.AlarmRecordCollection, error) {
	return arcu.listByPhysicalQuantity(startTime, endTime, domain.PhysicalQuantity{DeviceUUID: deviceUUID}, reverse)
}

func (arcu *AlarmRecordCollectionUsecase) ListByStationUUID(startTime, endTime time.Time, stationUUID string, reverse bool) ([]domain.AlarmRecordCollection, error) {
	return arcu.listByPhysicalQuantity(startTime, endTime, domain.PhysicalQuantity{StationUUID: stationUUID}, reverse)
}

func (arcu *AlarmRecordCollectionUsecase) listByPhysicalQuantity(startTime, endTime time.Time, physicalQuantity domain.PhysicalQuantity, reverse bool) ([]domain.AlarmRecordCollection, error) {
	physicalQuantities, listErr := arcu.physicalQuantity.List(physicalQuantity)
	if listErr != nil {
		return nil, listErr
	}

	var alarmRecordCollections []domain.AlarmRecordCollection

	alarmSettings, listErr := arcu.getAlarmSettings(physicalQuantities)
	if listErr != nil {
		return nil, listErr
	}

	alarmRecords, listErr := arcu.getAlarmRecords(startTime, endTime, alarmSettings, reverse)
	if listErr != nil {
		return nil, listErr
	}

	arcu.appendAlarmRecordCollections(&alarmRecordCollections, alarmRecords)

	deviceAlarmSettings, listErr := arcu.getDeviceAlarmSettingsByPhysicalQuantities(physicalQuantities)
	if listErr != nil {
		return nil, listErr
	}

	deviceAlarmRecords, listErr := arcu.getDeviceAlarmRecords(startTime, endTime, deviceAlarmSettings, reverse)
	if listErr != nil {
		return nil, listErr
	}

	arcu.appendDeviceAlarmRecordCollections(&alarmRecordCollections, deviceAlarmRecords)

	return arcu.resortOccurTime(alarmRecordCollections, reverse), nil
}

func (arcu *AlarmRecordCollectionUsecase) getAlarmSettings(physicalQuantities []domain.PhysicalQuantity) ([]domain.AlarmSetting, error) {
	physicalQuantityUUIDs := make([]string, len(physicalQuantities))
	for i, pq := range physicalQuantities {
		physicalQuantityUUIDs[i] = pq.UUID
	}

	return arcu.alarmSetting.ListInPhysicalQuantityUUIDs(domain.AlarmSetting{}, physicalQuantityUUIDs)
}

func (arcu *AlarmRecordCollectionUsecase) getAlarmRecords(startTime, endTime time.Time, alarmSettings []domain.AlarmSetting, reverse bool) ([]domain.AlarmRecord, error) {
	alarmSettingUUIDs := make([]string, len(alarmSettings))
	for i, alarmSettingPO := range alarmSettings {
		alarmSettingUUIDs[i] = alarmSettingPO.UUID
	}

	return arcu.alarmRecord.ListIn(startTime, endTime, alarmSettingUUIDs, reverse)
}

func (arcu *AlarmRecordCollectionUsecase) appendAlarmRecordCollections(alarmRecordCollections *[]domain.AlarmRecordCollection, alarmRecords []domain.AlarmRecord) {
	for _, alarmRecord := range alarmRecords {
		*alarmRecordCollections = append(*alarmRecordCollections, convert.AlarmRecordCollection{}.FromAlarmRecord(alarmRecord))
	}
}

func (arcu *AlarmRecordCollectionUsecase) getDeviceAlarmRecords(startTime, endTime time.Time, deviceAlarmSettings []domain.DeviceAlarmSetting, reverse bool) ([]domain.DeviceAlarmRecord, error) {
	deviceAlarmSettingUUIDs := make([]string, len(deviceAlarmSettings))
	for i, deviceAlarmSettingPO := range deviceAlarmSettings {
		deviceAlarmSettingUUIDs[i] = deviceAlarmSettingPO.UUID
	}

	return arcu.deviceAlarmRecord.ListIn(startTime, endTime, deviceAlarmSettingUUIDs, reverse)
}

func (arcu *AlarmRecordCollectionUsecase) appendDeviceAlarmRecordCollections(alarmRecordCollections *[]domain.AlarmRecordCollection, deviceAlarmRecords []domain.DeviceAlarmRecord) {
	for _, deviceAlarmRecord := range deviceAlarmRecords {
		*alarmRecordCollections = append(*alarmRecordCollections, convert.AlarmRecordCollection{}.FromDeviceAlarmRecord(deviceAlarmRecord))
	}
}

func (arcu *AlarmRecordCollectionUsecase) getDeviceAlarmSettingsByPhysicalQuantities(physicalQuantities []domain.PhysicalQuantity) ([]domain.DeviceAlarmSetting, error) {
	uniquePhysicalQuantity := lo.UniqBy(physicalQuantities, func(pq domain.PhysicalQuantity) string {
		return pq.DeviceUUID
	})

	deviceUUIDs := make([]string, len(uniquePhysicalQuantity))
	for i, pq := range uniquePhysicalQuantity {
		deviceUUIDs[i] = pq.DeviceUUID
	}

	deviceAlarmSettings, listErr := arcu.deviceAlarmSetting.ListInDeviceUUIDs(domain.DeviceAlarmSetting{}, deviceUUIDs)
	if listErr != nil {
		return nil, listErr
	}

	return deviceAlarmSettings, nil
}

func (AlarmRecordCollectionUsecase) resortOccurTime(arc []domain.AlarmRecordCollection, reverse bool) []domain.AlarmRecordCollection {
	if reverse {
		sort.Slice(arc, func(i, j int) bool {
			return arc[i].OccurTime.After(arc[j].OccurTime)
		})
	} else {
		sort.Slice(arc, func(i, j int) bool {
			return arc[i].OccurTime.Before(arc[j].OccurTime)
		})
	}
	return arc
}

package usecase

import (
	"fmt"
	"time"

	"github.com/samber/lo"

	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type RecordUsecase struct {
	record           irepository.IRecordRepository
	device           irepository.IDeviceRepository
	station          irepository.IStationRepository
	physicalQuantity irepository.IPhysicalQuantityRepository
}

func NewRecordUsecase(
	record irepository.IRecordRepository,
	device irepository.IDeviceRepository,
	station irepository.IStationRepository,
	physicalQuantity irepository.IPhysicalQuantityRepository,
) *RecordUsecase {
	return &RecordUsecase{
		record:           record,
		device:           device,
		station:          station,
		physicalQuantity: physicalQuantity,
	}
}

func (ru *RecordUsecase) ListArrayByDevice(start, end time.Time, deviceUUID string, timeZone string, sourceStr string) (domain.TimeSeries, error) {
	source, err := SourceFromString(sourceStr)
	if err != nil {
		return domain.TimeSeries{}, err
	}

	physicalQuantities, pqErr := ru.physicalQuantity.List(domain.PhysicalQuantity{DeviceUUID: deviceUUID, IsEnable: true, Source: source.String()})
	if pqErr != nil {
		return domain.TimeSeries{}, pqErr
	}

	// 解析包含時區信息的時間字符串
	t, timeErr := time.Parse(time.RFC3339, "2006-01-02T15:04:05"+timeZone)
	if timeErr != nil {
		return domain.TimeSeries{}, timeErr
	}

	recordsArray, listErr := ru.record.ListArray(start, end, physicalQuantities)
	if listErr != nil {
		return domain.TimeSeries{}, listErr
	}

	// 轉換時間
	for i, record := range recordsArray.Data {
		tmpTime, parseErr := time.Parse(time.RFC3339, record[0])
		if parseErr != nil {
			return domain.TimeSeries{}, parseErr
		}

		recordsArray.Data[i][0] = tmpTime.In(t.Location()).Format(time.RFC3339)
	}

	return recordsArray, nil
}

func (ru *RecordUsecase) ListMapByDevice(start, end time.Time, deviceUUID string, timeZone string, sourceStr string) ([]map[string]string, error) {
	source, err := SourceFromString(sourceStr)
	if err != nil {
		return nil, err
	}

	device, dErr := ru.device.FindByUUID(deviceUUID)
	if dErr != nil {
		return nil, dErr
	}

	physicalQuantities, pqErr := ru.physicalQuantity.List(domain.PhysicalQuantity{DeviceUUID: deviceUUID, IsEnable: true, Source: source.String()})
	if pqErr != nil {
		return nil, pqErr
	}

	// 解析包含時區信息的時間字符串
	t, timeErr := time.Parse(time.RFC3339, "2006-01-02T15:04:05"+timeZone)
	if timeErr != nil {
		return nil, timeErr
	}

	recordsMap, listErr := ru.record.ListMap(start, end, physicalQuantities)
	if listErr != nil {
		return nil, listErr
	}

	// 轉換時間
	for i, recordMap := range recordsMap {
		tmpTime, parseErr := time.Parse(time.RFC3339, recordMap["times"])
		if parseErr != nil {
			return nil, parseErr
		}
		recordsMap[i]["datetime"] = tmpTime.In(t.Location()).Format(time.RFC3339)
		delete(recordsMap[i], "times")

		// 添加設備信息
		recordsMap[i]["device_id"] = device.ID
	}

	return recordsMap, nil
}

func (ru *RecordUsecase) ListArrayByStation(start, end time.Time, stationUUID string, timeZone string, sourceStr string) (domain.TimeSeries, error) {
	source, err := SourceFromString(sourceStr)
	if err != nil {
		return domain.TimeSeries{}, err
	}

	physicalQuantities, pqErr := ru.physicalQuantity.List(domain.PhysicalQuantity{StationUUID: stationUUID, IsEnable: true, Source: source.String()})
	if pqErr != nil {
		return domain.TimeSeries{}, pqErr
	}

	// 解析包含時區信息的時間字符串
	t, timeErr := time.Parse(time.RFC3339, "2006-01-02T15:04:05"+timeZone)
	if timeErr != nil {
		return domain.TimeSeries{}, timeErr
	}

	recordsArray, listErr := ru.record.ListArray(start, end, physicalQuantities)
	if listErr != nil {
		return domain.TimeSeries{}, listErr
	}

	// 轉換時間
	for i, record := range recordsArray.Data {
		tmpTime, parseErr := time.Parse(time.RFC3339, record[0])
		if parseErr != nil {
			return domain.TimeSeries{}, parseErr
		}

		recordsArray.Data[i][0] = tmpTime.In(t.Location()).Format(time.RFC3339)
	}

	return recordsArray, nil
}

func (ru *RecordUsecase) ListMapByStation(start, end time.Time, stationUUID string, timeZone string, sourceStr string) ([]map[string]string, error) {
	source, err := SourceFromString(sourceStr)
	if err != nil {
		return nil, err
	}

	station, sErr := ru.station.FindByUUID(stationUUID)
	if sErr != nil {
		return nil, sErr
	}

	physicalQuantities, pqErr := ru.physicalQuantity.List(domain.PhysicalQuantity{StationUUID: stationUUID, IsEnable: true, Source: source.String()})
	if pqErr != nil {
		return nil, pqErr
	}

	// 解析包含時區信息的時間字符串
	t, timeErr := time.Parse(time.RFC3339, "2006-01-02T15:04:05"+timeZone)
	if timeErr != nil {
		return nil, timeErr
	}

	recordsMap, listErr := ru.record.ListMap(start, end, physicalQuantities)
	if listErr != nil {
		return nil, listErr
	}

	// 轉換時間
	for i, recordMap := range recordsMap {
		tmpTime, parseErr := time.Parse(time.RFC3339, recordMap["times"])
		if parseErr != nil {
			return nil, parseErr
		}
		recordsMap[i]["datetime"] = tmpTime.In(t.Location()).Format(time.RFC3339)
		delete(recordsMap[i], "times")

		// 添加站點信息
		recordsMap[i]["station_id"] = station.ID
		recordsMap[i]["station_latitude"] = fmt.Sprintf("%f", station.Lat)
		recordsMap[i]["station_longitude"] = fmt.Sprintf("%f", station.Lon)
	}

	return recordsMap, nil
}

func (ru *RecordUsecase) List(start, end time.Time, deviceUUID string, timeZone string, sourceStr string) ([]domain.Record, error) {
	source, err := SourceFromString(sourceStr)
	if err != nil {
		return nil, err
	}

	physicalQuantities, pqErr := ru.physicalQuantity.List(domain.PhysicalQuantity{DeviceUUID: deviceUUID, IsEnable: true, Source: source.String()})
	if pqErr != nil {
		return nil, pqErr
	}

	// 解析包含時區信息的時間字符串
	t, timeErr := time.Parse(time.RFC3339, "2006-01-02T15:04:05"+timeZone)
	if timeErr != nil {
		return nil, timeErr
	}

	// 找出物理量對應的站點
	tempPhysicalQuantities := lo.UniqBy(physicalQuantities, func(pq domain.PhysicalQuantity) string {
		return pq.StationUUID
	})

	stationUUIDs := make([]string, len(tempPhysicalQuantities))
	for i, pq := range tempPhysicalQuantities {
		stationUUIDs[i] = pq.StationUUID
	}

	var records []domain.Record
	for _, stationUUID := range stationUUIDs {
		tempRecords, listErr := ru.record.List(stationUUID, start, end, domain.Record{})
		if listErr != nil {
			return nil, listErr
		}

		records = append(records, tempRecords...)
	}

	// 轉換時間
	for i, record := range records {
		records[i].Datetime = record.Datetime.In(t.Location())
	}

	return records, nil
}

func (ru *RecordUsecase) Last(uuid string, timeZone string) (domain.Record, error) {
	physicalQuantity, pqErr := ru.physicalQuantity.FindByUUID(uuid)
	if pqErr != nil {
		return domain.Record{}, pqErr
	}

	// 解析包含時區信息的時間字符串
	t, timeErr := time.Parse(time.RFC3339, "2006-01-02T15:04:05"+timeZone)
	if timeErr != nil {
		return domain.Record{}, timeErr
	}

	record, listErr := ru.record.Last(physicalQuantity.StationUUID, physicalQuantity)
	if listErr != nil {
		return domain.Record{}, listErr
	}

	// 轉換時間
	record.Datetime = record.Datetime.In(t.Location())

	return record, nil
}

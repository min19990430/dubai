package usecase

import (
	"errors"
	"time"

	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
	"auto-monitoring/pkg/linear"
)

type CatchInputUsecase struct {
	physicalQuantity            irepository.IPhysicalQuantityRepository
	physicalQuantityCatchDetail irepository.IPhysicalQuantityCatchDetailRepository
	device                      irepository.IDeviceRepository
	station                     irepository.IStationRepository
	record                      irepository.IRecordRepository

	alarmUsecase AlarmUsecase
}

func NewCatchInputUsecase(
	physicalQuantity irepository.IPhysicalQuantityRepository,
	physicalQuantityCatchDetail irepository.IPhysicalQuantityCatchDetailRepository,
	device irepository.IDeviceRepository,
	station irepository.IStationRepository,
	record irepository.IRecordRepository,
	alarmUsecase AlarmUsecase,
) *CatchInputUsecase {
	return &CatchInputUsecase{
		physicalQuantity:            physicalQuantity,
		physicalQuantityCatchDetail: physicalQuantityCatchDetail,
		device:                      device,
		station:                     station,
		record:                      record,
		alarmUsecase:                alarmUsecase,
	}
}

func (ciu *CatchInputUsecase) CatchInput(inputs []domain.InputsWithDeviceUUID) error {
	for _, input := range inputs {
		err := ciu.catchInputWithDeviceUUID(input.DeviceUUID, input.Inputs)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ciu *CatchInputUsecase) catchInputWithDeviceUUID(deviceUUID string, inputs []domain.Input) error {
	if len(inputs) == 0 {
		return errors.New("inputs is empty")
	}

	pqcd, listErr := ciu.physicalQuantityCatchDetail.List(
		domain.PhysicalQuantity{
			DeviceUUID: deviceUUID,
			IsEnable:   true,
		})
	if listErr != nil {
		return listErr
	}

	pqcdMap := mappingStruct(pqcd)

	// 更新裝置時間
	updatedDevice := domain.Device{}
	updatedDevice.UUID = deviceUUID
	updatedDevice.IsConnected = true
	updatedDevice.UpdateTime = new(time.Time)
	*updatedDevice.UpdateTime = time.Now()
	if updateErr := ciu.device.UpdateLastTime(updatedDevice); updateErr != nil {
		return updateErr
	}

	for _, input := range inputs {
		// 檢查UUID
		physicalQuantityCatchDetail, ok := pqcdMap[input.UUID]
		if !ok {
			return errors.New("uuid not found")
			// continue
		}

		// 指標為nil時，初始化
		if physicalQuantityCatchDetail.UpdateTime == nil {
			physicalQuantityCatchDetail.UpdateTime = &time.Time{}
			*physicalQuantityCatchDetail.UpdateTime = time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
		}

		// 回補資料視為正常狀態"10"
		if input.Datetime.Before(*physicalQuantityCatchDetail.UpdateTime) {
			physicalQuantityCatchDetail.StatusCode = "10"
		}

		insertValue := input.Value
		// 如果啟用校正，則進行校正
		if physicalQuantityCatchDetail.CalibrationEnable {
			insertValue = linear.ComputeTwoPointLinearRegression(physicalQuantityCatchDetail.CalibrationValue, physicalQuantityCatchDetail.CalibrationParameter, input.Value)
		}

		physicalQuantityCatchDetail.Value = insertValue
		physicalQuantityCatchDetail.Data = input.Value

		// 只有在正常狀態下，才檢查告警
		if physicalQuantityCatchDetail.StatusCode == "10" || physicalQuantityCatchDetail.StatusCode == "11" || physicalQuantityCatchDetail.StatusCode == "93" {
			// 如果狀態為"93"，則由狀態"10"開始判斷
			if physicalQuantityCatchDetail.StatusCode == "93" {
				physicalQuantityCatchDetail.StatusCode = "10"
			}

			// 告警動作
			status, alarmErr := ciu.alarmUsecase.Check(physicalQuantityCatchDetail.AlarmSettings, input.Datetime, *physicalQuantityCatchDetail.UpdateTime, physicalQuantityCatchDetail.Value)
			if alarmErr != nil {
				return alarmErr
			}

			physicalQuantityCatchDetail.StatusCode = status
		}

		// 如果是新資料，則更新資料
		if input.Datetime.After(*physicalQuantityCatchDetail.UpdateTime) {
			*physicalQuantityCatchDetail.UpdateTime = input.Datetime
			if updateErr := ciu.physicalQuantity.UpdateLast(physicalQuantityCatchDetail.PhysicalQuantity); updateErr != nil {
				return updateErr
			}
		}

		if createErr := ciu.record.Create(physicalQuantityCatchDetail.StationUUID,
			domain.Record{
				DeviceUUID:           physicalQuantityCatchDetail.DeviceUUID,
				PhysicalQuantityUUID: physicalQuantityCatchDetail.UUID,
				Datetime:             input.Datetime,
				Value:                physicalQuantityCatchDetail.Value,
				Data:                 physicalQuantityCatchDetail.Data,
				Status:               physicalQuantityCatchDetail.StatusCode,
			}); createErr != nil {
			return createErr
		}
	}
	return nil
}

func mappingStruct(pqcds []domain.PhysicalQuantityCatchDetail) map[string]domain.PhysicalQuantityCatchDetail {
	pqcdMap := make(map[string]domain.PhysicalQuantityCatchDetail)
	for _, pqcd := range pqcds {
		pqcdMap[pqcd.UUID] = pqcd
	}
	return pqcdMap
}

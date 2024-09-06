package usecase

import (
	"strings"
	"text/template"
	"time"

	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
	"auto-monitoring/pkg/convert"
	timeexpression "auto-monitoring/pkg/expression/time_expression"
)

type DeviceAlarmSettingUsecase struct {
	deviceAlarmSetting irepository.IDeviceAlarmSettingRepository
	deviceAlarmRecord  irepository.IDeviceAlarmRecordRepository

	device irepository.IDeviceRepository
}

func NewDeviceAlarmSettingUsecase(
	deviceAlarmSetting irepository.IDeviceAlarmSettingRepository,
	deviceAlarmRecord irepository.IDeviceAlarmRecordRepository,
	device irepository.IDeviceRepository,
) *DeviceAlarmSettingUsecase {
	return &DeviceAlarmSettingUsecase{
		deviceAlarmSetting: deviceAlarmSetting,
		deviceAlarmRecord:  deviceAlarmRecord,
		device:             device,
	}
}

func (dau *DeviceAlarmSettingUsecase) Check(device domain.Device) error {
	// 取得該裝置的所有告警設定
	deviceAlarmSettings, listErr := dau.deviceAlarmSetting.List(domain.DeviceAlarmSetting{DeviceUUID: device.UUID, IsEnable: true})
	if listErr != nil {
		return listErr
	}

	if len(deviceAlarmSettings) == 0 {
		return nil
	}

	// 轉換裝置資料為map
	deviceMap := convert.StructToMap(device)

	for _, deviceAlarmSetting := range deviceAlarmSettings {
		// 檢查規則是否觸發
		expressionResult, expressionErr := timeexpression.IsTrue(deviceAlarmSetting.BooleanExpression, deviceMap)
		if expressionErr != nil {
			continue
		}

		// 如果規則沒觸發，則跳過
		if !expressionResult {
			// 如果狀態為告警，則更新狀態為正常
			if deviceAlarmSetting.IsActivated {
				deviceAlarmSetting.IsActivated = false
				if updateErr := dau.deviceAlarmSetting.UpdateStatusAndTime(deviceAlarmSetting); updateErr != nil {
					return updateErr
				}
			}
			continue
		}

		// 如果規則觸發
		// 檢查是否已經告警過，如果是，則跳過
		if deviceAlarmSetting.IsActivated {
			continue
		}

		// 未告警過，則更新狀態為告警
		deviceAlarmSetting.IsActivated = true
		deviceAlarmSetting.LastOccurTime = new(time.Time)
		*deviceAlarmSetting.LastOccurTime = time.Now()
		if updateErr := dau.deviceAlarmSetting.UpdateStatusAndTime(deviceAlarmSetting); updateErr != nil {
			return updateErr
		}

		// 更新裝置狀態
		device.IsConnected = false
		if updateErr := dau.device.UpdateLastTime(device); updateErr != nil {
			return updateErr
		}

		// 建立告警內容
		content, contentErr := dau.parseAlarmContent(deviceAlarmSetting)
		if contentErr != nil {
			return contentErr
		}

		// 建立告警紀錄
		deviceAlarmRecord := domain.DeviceAlarmRecord{
			DeviceAlarmSettingUUID: deviceAlarmSetting.UUID,
			Name:                   deviceAlarmSetting.Name,
			FullName:               deviceAlarmSetting.FullName,
			OccurTime:              *deviceAlarmSetting.LastOccurTime,
			Content:                content,
		}
		if createErr := dau.deviceAlarmRecord.Create(deviceAlarmRecord); createErr != nil {
			return createErr
		}

		// 是否告警通知
		if deviceAlarmSetting.IsNotify {
			if notifyErr := dau.deviceAlarmSetting.Notify(deviceAlarmSetting, content); notifyErr != nil {
				return notifyErr
			}
		}
	}
	return nil
}

func (dau *DeviceAlarmSettingUsecase) parseAlarmContent(deviceAlarmSetting domain.DeviceAlarmSetting) (string, error) {
	if deviceAlarmSetting.ContentSetting == nil {
		return "設備告警", nil
	}

	tmpl, err := template.New("alarm").Parse(*deviceAlarmSetting.ContentSetting)
	if err != nil {
		return "", err
	}
	var messageBuilder strings.Builder
	err = tmpl.Execute(&messageBuilder, deviceAlarmSetting)
	if err != nil {
		return "", err
	}
	return messageBuilder.String(), nil
}

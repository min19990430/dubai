package usecase

import (
	"strings"
	"text/template"
	"time"

	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type AlarmUsecase struct {
	alarmSetting irepository.IAlarmSettingRepository
	alarmRecord  irepository.IAlarmRecordRepository
}

func NewAlarmUsecase(
	alarmSetting irepository.IAlarmSettingRepository,
	alarmRecord irepository.IAlarmRecordRepository,
) *AlarmUsecase {
	return &AlarmUsecase{
		alarmSetting: alarmSetting,
		alarmRecord:  alarmRecord,
	}
}

func (au *AlarmUsecase) Check(alarmSettings []domain.AlarmSetting, inputTime, updateTime time.Time, value float64) (string, error) {
	status := "10"
	for _, alarmSetting := range alarmSettings {
		// 如果告警未啟用，則跳過
		if !alarmSetting.IsEnable {
			continue
		}

		// 檢查規則是否觸發
		// expressionResult, expressionErr := expression.IsTrue(alarmSetting.BooleanExpression, value)
		// if expressionErr != nil {
		// 	continue
		// }

		expressionResult, expressionErr := au.alarmSetting.ExpressionCheck(alarmSetting, value)
		if expressionErr != nil {
			continue
		}

		// 如果規則沒觸發，則跳過
		if !expressionResult {
			// 如果狀態為告警，則更新狀態為正常
			if alarmSetting.IsActivated {
				alarmSetting.IsActivated = false
				if updateErr := au.alarmSetting.UpdateStatusAndTime(alarmSetting); updateErr != nil {
					return "", updateErr
				}
			}
			continue
		}

		// 如果規則觸發
		// 則為告警狀態
		status = "11"
		// 檢查是否事舊資料的告警，如果是，則跳過
		if inputTime.Before(updateTime) {
			continue
		}

		// 新資料的告警
		// 檢查是否已經告警過，如果是，則跳過
		if alarmSetting.IsActivated {
			continue
		}

		// 未告警過
		// 更新告警
		alarmSetting.IsActivated = true
		alarmSetting.LastAlarmOccurTime = new(time.Time)
		*alarmSetting.LastAlarmOccurTime = inputTime
		alarmSetting.LastOccurValue = value
		if updateErr := au.alarmSetting.UpdateStatusAndTime(alarmSetting); updateErr != nil {
			return "", updateErr
		}

		// 建立告警內容
		content, contentErr := au.parseAlarmContent(alarmSetting)
		if contentErr != nil {
			return "", contentErr
		}

		// 建立告警紀錄
		alarmRecord := domain.AlarmRecord{
			AlarmSettingUUID: alarmSetting.UUID,
			Name:             alarmSetting.Name,
			FullName:         alarmSetting.FullName,
			OccurTime:        inputTime,
			Content:          content,
		}
		if createErr := au.alarmRecord.Create(alarmRecord); createErr != nil {
			return "", createErr
		}

		// 是否告警通知
		if alarmSetting.IsNotify {
			if notifyErr := au.alarmSetting.Notify(alarmSetting, content); notifyErr != nil {
				return "", notifyErr
			}
		}
	}
	return status, nil
}

func (au *AlarmUsecase) parseAlarmContent(alarmSetting domain.AlarmSetting) (string, error) {
	if alarmSetting.AlarmContentSetting == nil {
		return "異常告警", nil
	}

	tmpl, err := template.New("alarm").Parse(*alarmSetting.AlarmContentSetting)
	if err != nil {
		return "", err
	}
	var messageBuilder strings.Builder
	err = tmpl.Execute(&messageBuilder, alarmSetting)
	if err != nil {
		return "", err
	}
	return messageBuilder.String(), nil
}

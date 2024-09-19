package job

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"

	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

func (daj *DeviceAlarmJob) Setup(s gocron.Scheduler) {
	_, err := s.NewJob(gocron.DurationJob(time.Minute), gocron.NewTask(daj.deviceAlarmCheck))
	if err != nil {
		log.Fatalf("new job error: %v", err)
	}
}

type DeviceAlarmJob struct {
	logger             *zap.Logger
	deviceAlarmSetting usecase.DeviceAlarmSettingUsecase
	device             usecase.DeviceUsecase
}

func NewDeviceAlarmJob(logger *zap.Logger, deviceAlarmSetting *usecase.DeviceAlarmSettingUsecase, device *usecase.DeviceUsecase) *DeviceAlarmJob {
	return &DeviceAlarmJob{
		logger:             logger,
		deviceAlarmSetting: *deviceAlarmSetting,
		device:             *device,
	}
}

func (daj *DeviceAlarmJob) deviceAlarmCheck() {
	// 取得所有裝置
	devices, _ := daj.device.List(domain.Device{IsEnable: true})
	for _, device := range devices {
		if err := daj.deviceAlarmSetting.Check(device); err != nil {
			daj.logger.Error("device alarm check error", zap.Error(err))
			continue
		}
	}
}

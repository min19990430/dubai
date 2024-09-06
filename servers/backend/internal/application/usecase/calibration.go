package usecase

import (
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
	"auto-monitoring/pkg/linear"
)

type CalibrationUsecase struct {
	calibration irepository.ICalibrationRepository
}

func NewCalibrationUsecase(calibration irepository.ICalibrationRepository) *CalibrationUsecase {
	return &CalibrationUsecase{
		calibration: calibration,
	}
}

func (cu *CalibrationUsecase) TrialCalculator(deviceUUID string) ([]domain.CalibrationDetail, error) {
	calibrationDetails, err := cu.calibration.FindDetailByDeviceUUID(deviceUUID)
	if err != nil {
		return nil, err
	}

	for i, calibrationDetail := range calibrationDetails {
		calibrationDetails[i].Value = cu.calculator(calibrationDetail.Calibration)
	}
	return calibrationDetails, nil
}

func (CalibrationUsecase) calculator(calibration domain.Calibration) float64 {
	return linear.ComputeTwoPointLinearRegression(calibration.CalibrationValue, calibration.CalibrationParameter, calibration.Data)
}

func (cu *CalibrationUsecase) Update(deviceUUID string, isCalibration bool, calibrationValue string, calibrationParameter string) error {
	return cu.calibration.Update(deviceUUID, isCalibration, calibrationValue, calibrationParameter)
}

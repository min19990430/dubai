package usecase

import (
	"errors"
	"time"

	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

const (
	timeCountLimit = 99999
)

type TimeSeriesUsecase struct {
	timeSeries       irepository.ITimeSeriesRepository
	device           irepository.IDeviceRepository
	physicalQuantity irepository.IPhysicalQuantityRepository
}

func NewTimeSeriesUsecase(physicalQuantity irepository.IPhysicalQuantityRepository, device irepository.IDeviceRepository, timeSeries irepository.ITimeSeriesRepository) *TimeSeriesUsecase {
	return &TimeSeriesUsecase{
		timeSeries:       timeSeries,
		device:           device,
		physicalQuantity: physicalQuantity,
	}
}

func (tsu *TimeSeriesUsecase) AggregateDataByStation(start, end time.Time, stationUUID string, intervalStr string, reverse bool) (domain.TimeSeries, error) {
	interval, convertErr := tsu.intervalToDuration(intervalStr)
	if convertErr != nil {
		return domain.TimeSeries{}, convertErr
	}

	if checkErr := tsu.inputCheck(start, end, interval); checkErr != nil {
		return domain.TimeSeries{}, checkErr
	}

	var physicalQuantities []domain.PhysicalQuantity
	physicalQuantities, pqListErr := tsu.physicalQuantity.List(domain.PhysicalQuantity{
		StationUUID: stationUUID,
		Source:      Sensor.String(),
		IsEnable:    true,
	})
	if pqListErr != nil {
		return domain.TimeSeries{}, pqListErr
	}

	timeSeriesList, aggregateErr := tsu.timeSeries.Aggregate(start, end, interval, physicalQuantities)
	if aggregateErr != nil {
		return domain.TimeSeries{}, aggregateErr
	}

	if reverse {
		for i, j := 0, len(timeSeriesList.Data)-1; i < j; i, j = i+1, j-1 {
			timeSeriesList.Data[i], timeSeriesList.Data[j] = timeSeriesList.Data[j], timeSeriesList.Data[i]
		}
	}

	return timeSeriesList, nil
}

func (tsu *TimeSeriesUsecase) AggregateMapDataByStation(start, end time.Time, stationUUID string, intervalStr string) ([]map[string]string, error) {
	interval, convertErr := tsu.intervalToDuration(intervalStr)
	if convertErr != nil {
		return nil, convertErr
	}

	if checkErr := tsu.inputCheck(start, end, interval); checkErr != nil {
		return nil, checkErr
	}

	physicalQuantities, pqListErr := tsu.physicalQuantity.List(domain.PhysicalQuantity{
		StationUUID: stationUUID,
		Source:      Sensor.String(),
		IsEnable:    true,
	})
	if pqListErr != nil {
		return nil, pqListErr
	}

	return tsu.timeSeries.AggregateMap(start, end, interval, physicalQuantities)
}

func (TimeSeriesUsecase) intervalToDuration(interval string) (time.Duration, error) {
	switch interval {
	case "1m":
		return time.Minute, nil
	case "5m":
		return 5 * time.Minute, nil
	case "10m":
		return 10 * time.Minute, nil
	case "30m":
		return 30 * time.Minute, nil
	case "1h":
		return time.Hour, nil
	case "1d":
		return 24 * time.Hour, nil
	default:
		return 0, errors.New("invalid interval")
	}
}

func (tsu *TimeSeriesUsecase) inputCheck(start, end time.Time, interval time.Duration) error {
	if start.After(end) {
		return errors.New("start time is after end time")
	}
	if interval == 0 {
		return errors.New("invalid interval")
	}

	if _, err := tsu.checkTimeCount(start, end, interval); err != nil {
		return err
	}

	return nil
}

func (tsu *TimeSeriesUsecase) checkTimeCount(start, end time.Time, interval time.Duration) (int, error) {
	count := tsu.timeCount(start, end, interval)
	if count > timeCountLimit {
		return 0, errors.New("too many data")
	}

	return count, nil
}

func (TimeSeriesUsecase) timeCount(start, end time.Time, interval time.Duration) int {
	return int(end.Sub(start) / interval)
}

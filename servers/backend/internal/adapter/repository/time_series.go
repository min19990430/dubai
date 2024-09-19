package repository

import (
	"database/sql"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"auto-monitoring/internal/application/convert"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type TimeSeriesRepository struct {
	gorm *gorm.DB
}

func NewTimeSeriesRepository(gorm *gorm.DB) irepository.ITimeSeriesRepository {
	return &TimeSeriesRepository{gorm: gorm}
}

func (t *TimeSeriesRepository) Aggregate(start time.Time, end time.Time, interval time.Duration, physicalQuantities []domain.PhysicalQuantity) (domain.TimeSeries, error) {
	// 建立時間序列
	timeSeriesSubQuery := t.createTimeQueryWithInterval(t.gorm, start, end, interval)

	// 建立查詢
	query := t.gorm.Table("(?) as time_series", timeSeriesSubQuery)

	// FIXME: 這裡的SELECT字串需要修改
	// selectStr := "FROM_UNIXTIME(time_series.times) as times"
	selectStr := "DATE_FORMAT(CONVERT_TZ(FROM_UNIXTIME(time_series.times), '+08:00', '+00:00'), '%Y-%m-%dT%TZ') as times"
	for _, pq := range physicalQuantities {
		// 建立SELECT字串
		selectStr += t.createPhysicalQuantitySelectString(pq)

		// 建立物理量子查詢
		pqSubQuery := t.createPhysicalQuantitySubQuery(pq, start, end, interval)

		// JOIN子查詢
		query = query.Joins("LEFT JOIN (?) as "+t.physicalQuantityTempTable(pq)+" ON time_series.times = "+t.physicalQuantityTempTable(pq)+".pqtime", pqSubQuery)
	}

	// 執行查詢
	rows, rowsErr := query.Select(selectStr).Order("times").Rows()
	if rowsErr != nil {
		return domain.TimeSeries{}, rowsErr
	}

	return t.dumpRowsToStringArray(physicalQuantities, rows)
}

func (t *TimeSeriesRepository) AggregateMap(start time.Time, end time.Time, interval time.Duration, physicalQuantities []domain.PhysicalQuantity) ([]map[string]string, error) {
	// 建立時間序列
	timeSeriesSubQuery := t.createTimeQueryWithInterval(t.gorm, start, end, interval)

	// 建立查詢
	query := t.gorm.Table("(?) as time_series", timeSeriesSubQuery)
	// selectStr := "FROM_UNIXTIME(time_series.times) as times"
	selectStr := "DATE_FORMAT(CONVERT_TZ(FROM_UNIXTIME(time_series.times), '+08:00', '+00:00'), '%Y-%m-%dT%TZ') as times"
	for _, pq := range physicalQuantities {
		// 建立SELECT字串
		selectStr += t.createPhysicalQuantitySelectString(pq)

		// 建立物理量子查詢
		pqSubQuery := t.createPhysicalQuantitySubQuery(pq, start, end, interval)

		// JOIN子查詢
		query = query.Joins("LEFT JOIN (?) as "+t.physicalQuantityTempTable(pq)+" ON time_series.times = "+t.physicalQuantityTempTable(pq)+".pqtime", pqSubQuery)
	}

	// 執行查詢
	rows, rowsErr := query.Select(selectStr).Order("times").Rows()
	if rowsErr != nil {
		return nil, rowsErr
	}

	_, rowsMap, mapErr := t.dumpRowsToMapArray(physicalQuantities, rows)
	return rowsMap, mapErr
}

func (TimeSeriesRepository) timeCount(start, end time.Time, interval time.Duration) int {
	return int(end.Sub(start)/interval) + 1
}

// ! 至多到99999
func (t *TimeSeriesRepository) createTimeQueryWithInterval(db *gorm.DB, startTime, endTime time.Time, interval time.Duration) *gorm.DB {
	timeCount := t.timeCount(startTime, endTime, interval)

	intervalSec := int(interval.Seconds())

	numberTableStr := "SELECT 0 AS n UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5 UNION ALL SELECT 6 UNION ALL SELECT 7 UNION ALL SELECT 8 UNION ALL SELECT 9"
	subQuery := t.gorm.Raw("(" + numberTableStr + ") as ones CROSS JOIN (" + numberTableStr + ") as tens CROSS JOIN (" + numberTableStr + ") as hundreds CROSS JOIN (" + numberTableStr + ") as thousands CROSS JOIN (" + numberTableStr + ") as ten_thousands")

	return db.Table("(?)", subQuery).
		Select("UNIX_TIMESTAMP(?) + "+strconv.Itoa(intervalSec)+" * (ones.n + 10 * tens.n + 100 * hundreds.n + 1000 * thousands.n + 10000 * ten_thousands.n) AS times", startTime).
		Limit(timeCount).
		Order("times")
}

func (t *TimeSeriesRepository) createPhysicalQuantitySelectString(pq domain.PhysicalQuantity) string {
	var selectStr string
	switch {
	case pq.PhysicalQuantityDataType == "Integer":
		selectStr += ", truncate(IFNULL(" + t.physicalQuantityTempTable(pq) + ".value,0),0) as " + t.physicalQuantityTempTable(pq) + "_value"
	case t.matchedDeciaml(pq.PhysicalQuantityDataType):
		selectStr += ", truncate(IFNULL(" + t.physicalQuantityTempTable(pq) + ".value,0)," + pq.PhysicalQuantityDataType[7:8] + ") as " + t.physicalQuantityTempTable(pq) + "_value"
	default:
		selectStr += ", IFNULL(" + t.physicalQuantityTempTable(pq) + ".value,0) as " + t.physicalQuantityTempTable(pq) + "_value"
	}
	return selectStr
}

func (t *TimeSeriesRepository) createPhysicalQuantitySubQuery(pq domain.PhysicalQuantity, startTime, endTime time.Time, interval time.Duration) *gorm.DB {
	intervalSec := int(interval.Seconds())

	pqSubQuery := t.gorm.
		Table(pq.StationUUID).
		Where("datetime between ? and ?", startTime, endTime).
		Where("physical_quantity_uuid = ?", pq.UUID).
		Group("pqtime")

	pqTimeSelect := "UNIX_TIMESTAMP(datetime) - UNIX_TIMESTAMP(datetime) % " + strconv.Itoa(intervalSec) + " as pqtime"
	switch pq.AggregateCalculationMethod {
	case "diff":
		pqSubQuery.Select(pqTimeSelect + ", MAX(value)-MIN(value) as value")
	case "avg":
		pqSubQuery.Select(pqTimeSelect + ", AVG(value) as value")
	case "max":
		pqSubQuery.Select(pqTimeSelect + ", MAX(value) as value")
	default:
		pqSubQuery.Select(pqTimeSelect + ", AVG(value) as value")
	}
	return pqSubQuery
}

func (TimeSeriesRepository) matchedDeciaml(input string) bool {
	r := regexp.MustCompile("^Decimal[1-6]$")
	return r.MatchString(input)
}

func (TimeSeriesRepository) dumpRowsToStringArray(physicalQuantities []domain.PhysicalQuantity, rows *sql.Rows) (domain.TimeSeries, error) {
	var recordTable domain.TimeSeries

	cols, err := rows.Columns()
	if err != nil {
		return domain.TimeSeries{}, err
	}

	recordTable.Columns = append(recordTable.Columns, convert.TimeSeriesColumn{}.CreateTimeColumn())
	for _, pq := range physicalQuantities {
		recordTable.Columns = append(recordTable.Columns, convert.TimeSeriesColumn{}.FromPhysicalQuantity(pq))
	}

	readCols := make([]interface{}, len(cols))
	writeCols := make([]string, len(cols))
	for i := range writeCols {
		readCols[i] = &writeCols[i]
	}

	for rows.Next() {
		err = rows.Scan(readCols...)
		if err != nil {
			return domain.TimeSeries{}, err
		}
		copiedSlice := make([]string, len(cols))
		copy(copiedSlice, writeCols)
		recordTable.Data = append(recordTable.Data, copiedSlice)
	}
	return recordTable, nil
}

func (TimeSeriesRepository) dumpRowsToMapArray(physicalQuantities []domain.PhysicalQuantity, rows *sql.Rows) ([]string, []map[string]string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}

	readCols := make([]interface{}, len(cols))
	writeCols := make([]string, len(cols))
	for i := range writeCols {
		readCols[i] = &writeCols[i]
	}

	var data []map[string]string

	for rows.Next() {
		err = rows.Scan(readCols...)
		if err != nil {
			return nil, nil, err
		}

		tempMap := make(map[string]string)

		for i, v := range cols {
			if i == 0 {
				tempMap[v] = writeCols[i]
				continue
			}
			tempMap[physicalQuantities[i-1].FullName] = writeCols[i]
		}

		data = append(data, tempMap)
	}
	return cols, data, nil
}

func (TimeSeriesRepository) physicalQuantityTempTable(pq domain.PhysicalQuantity) string {
	return pq.Name + "_" + strings.ReplaceAll(pq.DeviceUUID, "-", "_")
}

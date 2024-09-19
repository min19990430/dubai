package repository

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"auto-monitoring/internal/adapter/gorm/model"
	"auto-monitoring/internal/domain"
	"auto-monitoring/internal/domain/irepository"
)

type RecordRepository struct {
	gorm *gorm.DB
}

func NewRecordRepository(gorm *gorm.DB) irepository.IRecordRepository {
	return &RecordRepository{gorm: gorm}
}

func (r *RecordRepository) Create(tableName string, record domain.Record) error {
	recordsPO := model.Record{}.FromDomain(record)

	err := r.gorm.Table(tableName).Create(&recordsPO).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		// if duplicate key, not error
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil
		}

		// if table is not created
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1146 {
			if createErr := r.createTable(tableName); createErr != nil {
				return createErr
			}
			err = r.gorm.Table(tableName).Create(&recordsPO).Error
		}

		return err
	}

	return nil
}

func (r *RecordRepository) CreateMany(tableName string, records []domain.Record) error {
	for _, record := range records {
		if err := r.Create(tableName, record); err != nil {
			return err
		}
	}

	return nil
}

func (r *RecordRepository) List(tableName string, start, end time.Time) ([]domain.Record, error) {
	var recordsPO []model.Record
	if err := r.gorm.Table(tableName).
		Where("datetime BETWEEN ? AND ?", start, end).
		Find(&recordsPO).Error; err != nil {
		return nil, err
	}

	records := make([]domain.Record, 0, len(recordsPO))
	for _, recordPO := range recordsPO {
		records = append(records, recordPO.ToDomain())
	}

	return records, nil
}

func (r *RecordRepository) createTable(tableName string) error {
	if createErr := r.gorm.Migrator().
		CreateTable(&model.Record{}); createErr != nil {
		return createErr
	}
	if renameErr := r.gorm.Migrator().
		RenameTable(model.Record{}.TableName(), tableName); renameErr != nil {
		return renameErr
	}
	return nil
}

func (r *RecordRepository) ListMap(start, end time.Time, physicalQuantities []domain.PhysicalQuantity) ([]map[string]string, error) {
	if len(physicalQuantities) == 0 {
		return nil, errors.New("empty physical quantities")
	}

	// 建立時間序列
	timeSeriesSubQuery := r.createTimeQuery(physicalQuantities[0].StationUUID, physicalQuantities[0], start, end)

	// 建立查詢
	query := r.gorm.Table("(?) as time_series", timeSeriesSubQuery)

	selectStr := "time_series.times as times"

	for _, pq := range physicalQuantities {
		// 建立SELECT字串
		selectStr += r.createPhysicalQuantitySelectString(pq)

		// 建立物理量子查詢
		pqSubQuery := r.createPhysicalQuantitySubQuery(pq.StationUUID, pq, start, end)

		// JOIN子查詢
		query = query.Joins("LEFT JOIN (?) as "+r.physicalQuantityTempTable(pq)+" ON time_series.times = "+r.physicalQuantityTempTable(pq)+".pqtime", pqSubQuery)
	}

	// 執行查詢
	rows, rowsErr := query.Select(selectStr).Order("times").Rows()
	if rowsErr != nil {
		return nil, rowsErr
	}

	_, rowsMap, mapErr := r.dumpRowsToMapArray(physicalQuantities, rows)
	return rowsMap, mapErr
}

func (r *RecordRepository) Last(tableName string, physicalQuantity domain.PhysicalQuantity) (domain.Record, error) {
	var recordPO model.Record

	err := r.gorm.
		Table(tableName).
		Where("physical_quantity_uuid = ?", physicalQuantity.UUID).
		Order("datetime DESC").
		First(&recordPO).Error
	if err != nil {
		return domain.Record{}, err
	}

	return recordPO.ToDomain(), nil
}

func (r *RecordRepository) createTimeQuery(tableName string, physicalQuantity domain.PhysicalQuantity, startTime, endTime time.Time) *gorm.DB {
	return r.gorm.Table(tableName).
		Select("datetime as times").
		Where("physical_quantity_uuid = ?", physicalQuantity.UUID).
		Where("datetime BETWEEN ? AND ?", startTime, endTime).
		Order("datetime")
}

func (r *RecordRepository) createPhysicalQuantitySelectString(pq domain.PhysicalQuantity) string {
	return ", " + r.physicalQuantityTempTable(pq) + ".value as " + pq.Name + ", " + r.physicalQuantityTempTable(pq) + ".status as " + pq.Name + "_status"
}

func (r *RecordRepository) createPhysicalQuantitySubQuery(tableName string, pq domain.PhysicalQuantity, startTime, endTime time.Time) *gorm.DB {
	return r.gorm.Table(tableName).
		Select("datetime as pqtime", "value", "status").
		Where("datetime BETWEEN ? AND ?", startTime, endTime).
		Where("physical_quantity_uuid = ?", pq.UUID).
		Order("datetime")
}

func (RecordRepository) physicalQuantityTempTable(pq domain.PhysicalQuantity) string {
	return pq.Name + "_" + strings.ReplaceAll(pq.DeviceUUID, "-", "_")
}

func (RecordRepository) dumpRowsToMapArray(physicalQuantities []domain.PhysicalQuantity, rows *sql.Rows) ([]string, []map[string]string, error) {
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

		// 第一格必為時間
		tempMap[cols[0]] = writeCols[0]
		// 填充物理量與狀態
		j := 0
		for i := 1; i < len(cols); i += 2 {
			tempMap[physicalQuantities[j].Name] = writeCols[i]
			tempMap[physicalQuantities[j].Name+"_status"] = writeCols[i+1]
			j++
		}

		data = append(data, tempMap)
	}
	return cols, data, nil
}

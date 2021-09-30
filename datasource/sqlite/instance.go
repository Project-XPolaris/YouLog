package sqlite

import (
	"fmt"
	"github.com/projectxpolaris/youlog/datasource"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
)

var DefaultSqliteDataSource datasource.LogDataSource = &DataSource{}

type DataSource struct {
	Instance *gorm.DB
}

func (s *DataSource) ReadLogs(queryBuilder datasource.LogListQueryBuilder) (int64, []datasource.Log, error) {
	var list []*Log
	var count int64
	queryDB := s.Instance
	if queryBuilder.LogLevels != nil {
		queryDB = queryDB.Where("level IN ?", queryBuilder.LogLevels)
	}
	if queryBuilder.StartTime != nil {
		queryDB = queryDB.Where("time >= ?", queryBuilder.StartTime.Format("2006-01-02 15:04:05"))
	}
	if queryBuilder.EndTime != nil {
		queryDB = queryDB.Where("time <= ?", queryBuilder.EndTime.Format("2006-01-02 15:04:05"))
	}
	if len(queryBuilder.Application) > 0 {
		queryDB = queryDB.Where("application = ?", queryBuilder.Application)
	}
	if queryBuilder.Orders != nil {
		for _, order := range queryBuilder.Orders {
			if strings.HasPrefix(order, "-") {
				orderKey := order[1:]
				queryDB = queryDB.Order(fmt.Sprintf("%s DESC", orderKey))
			} else {
				queryDB = queryDB.Order(order)
			}
		}
	}
	err := queryDB.Offset((queryBuilder.Page - 1) * queryBuilder.PageSize).Limit(queryBuilder.PageSize).Find(&list).Count(&count).Error
	if err != nil {
		return count, nil, err
	}
	result := make([]datasource.Log, 0)
	for _, logRecord := range list {
		result = append(result, logRecord)
	}
	return count, result, nil
}

func (s *DataSource) Init() error {
	var err error
	s.Instance, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = s.Instance.AutoMigrate(&Log{})
	if err != nil {
		return err
	}
	return nil
}

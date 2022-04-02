package database

import (
	"fmt"
	"github.com/projectxpolaris/youlog/datasource"
	"gorm.io/gorm"
	"strings"
)

var DatabaseDataSourceList = make([]*DatabaseDataSource, 0)

type DatabaseDataSource struct {
	Name     string
	Instance *gorm.DB
}

func (s *DatabaseDataSource) ReadLogs(queryBuilder datasource.LogListQueryBuilder) (int64, []datasource.Log, error) {
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
	if len(queryBuilder.ApplicationSearch) > 0 {
		queryDB = queryDB.Where("application like ?", fmt.Sprintf("%%%s%%", queryBuilder.ApplicationSearch))
	}
	if queryBuilder.DistinctApp {
		queryDB = queryDB.Distinct("Application")
	}
	err := queryDB.Offset((queryBuilder.Page - 1) * queryBuilder.PageSize).Limit(queryBuilder.PageSize).Find(&list).Offset(-1).Count(&count).Error
	if err != nil {
		return count, nil, err
	}
	result := make([]datasource.Log, 0)
	for _, logRecord := range list {
		result = append(result, logRecord)
	}
	return count, result, nil
}

func (s *DatabaseDataSource) Init() error {
	return nil
}

func GetDefaultDataSource() datasource.LogDataSource {
	return DatabaseDataSourceList[0]
}

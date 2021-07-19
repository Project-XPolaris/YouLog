package service

import "github.com/projectxpolaris/youlog/database"

func GetLogList(page int, pageSize int) (int64, []*database.Log, error) {
	var list []*database.Log
	var count int64
	err := database.Instance.Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Count(&count).Error
	if err != nil {
		return count, nil, err
	}
	return count, list, nil
}

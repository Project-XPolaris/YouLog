package database

import (
	"github.com/allentom/harukap/datasource"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DefaultDatabasePlugin = &datasource.Plugin{
	OnConnected: func(db *gorm.DB) {
		err := db.AutoMigrate(&Log{})
		if err != nil {
			logrus.Error(err)
		}
	},
}

package service

import (
	"github.com/project-xpolaris/youplustoolkit/youlog/logservice"
	"github.com/projectxpolaris/youlog/datasource/database"
	"gorm.io/gorm"
	"time"
)

var DefaultDBOutputs = make([]*DatabaseOutput, 0)

type DatabaseOutput struct {
	Name string
	DB   *gorm.DB
}

func (o *DatabaseOutput) writeLog(data *logservice.LogData) error {
	logData := database.Log{
		Application: data.Application,
		Instance:    data.Instance,
		Scope:       data.Scope,
		Level:       data.Level,
		Message:     data.Message,
		Extra:       data.Extra,
		Time:        time.Unix(0, data.Time*int64(time.Millisecond)),
	}
	return o.DB.Save(&logData).Error
}

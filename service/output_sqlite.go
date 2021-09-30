package service

import (
	"github.com/projectxpolaris/youlog/datasource/sqlite"
	"github.com/projectxpolaris/youlog/pb"
	"time"
)

type SqliteLogOutput struct {
}

func (o *SqliteLogOutput) writeLog(data *pb.LogData) error {
	logData := sqlite.Log{
		Application: data.Application,
		Instance:    data.Instance,
		Scope:       data.Scope,
		Level:       data.Level,
		Message:     data.Message,
		Extra:       data.Extra,
		Time:        time.Unix(0, data.Time*int64(time.Millisecond)),
	}
	return sqlite.DefaultSqliteDataSource.(*sqlite.DataSource).Instance.Save(&logData).Error
}

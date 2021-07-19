package service

import (
	"context"
	"github.com/projectxpolaris/youlog/database"
	"github.com/projectxpolaris/youlog/pb"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	DefaultLogWriter *LogWriter
	writeLogger      *logrus.Entry
)

func init() {
	DefaultLogWriter = &LogWriter{
		In: make(chan *pb.LogData),
	}
	writeLogger = logrus.WithField("scope", "LogWriter")
}

type LogWriter struct {
	In chan *pb.LogData
}

func (w *LogWriter) Run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case in := <-w.In:
				err := w.writeLog(in)
				if err != nil {
					writeLogger.Error(err)
				}
			}

		}
	}()
	writeLogger.Info("writer is running")
}
func (w *LogWriter) writeLog(data *pb.LogData) error {
	logData := database.Log{
		Application: data.Application,
		Instance:    data.Instance,
		Scope:       data.Scope,
		Level:       data.Level,
		Message:     data.Message,
		Extra:       data.Extra,
		Time:        time.Unix(0, data.Time*int64(time.Millisecond)),
	}
	return database.Instance.Save(&logData).Error
}

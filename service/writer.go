package service

import (
	"context"
	"github.com/project-xpolaris/youplustoolkit/youlog/logservice"
	"github.com/sirupsen/logrus"
)

var (
	DefaultLogWriter *LogWriter
	writeLogger      *logrus.Entry
)

func init() {
	DefaultLogWriter = &LogWriter{
		In: make(chan *logservice.LogData),
	}
	writeLogger = logrus.WithField("scope", "LogWriter")
}

type LogWriter struct {
	In chan *logservice.LogData
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
func (w *LogWriter) writeLog(data *logservice.LogData) error {
	for _, output := range GetLogOutputs() {
		err := output.writeLog(data)
		if err != nil {
			return err
		}
	}
	return nil
}

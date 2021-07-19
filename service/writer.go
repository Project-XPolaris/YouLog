package service

import (
	"context"
	"github.com/projectxpolaris/youlog/pb"
	"github.com/sirupsen/logrus"
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
	for _, output := range GetLogOutputs() {
		err := output.writeLog(data)
		if err != nil {
			return err
		}
	}
	return nil
}

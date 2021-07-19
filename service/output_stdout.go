package service

import (
	"encoding/json"
	"github.com/projectxpolaris/youlog/pb"
	"github.com/sirupsen/logrus"
	"time"
)

var defaultTimeFormat = "2006-01-02 15:04:05"

type StdoutLogOutput struct {
}

func (o *StdoutLogOutput) writeLog(data *pb.LogData) error {
	extra := map[string]interface{}{}
	err := json.Unmarshal([]byte(data.Extra), &extra)
	if err != nil {
		return err
	}
	fields := logrus.Fields{
		"Application": data.Application,
		"Instance":    data.Instance,
		"Scope":       data.Scope,
		"Time":        time.Unix(0, data.Time*int64(time.Millisecond)).Format(defaultTimeFormat),
	}
	for key, value := range extra {
		fields[key] = value
	}
	entity := logrus.WithFields(fields)
	switch data.Level {
	case LEVEL_DEBUG:
		entity.Debug(data.Message)
	case LEVEL_ERROR:
		entity.Error(err)
	case LEVEL_FATAL:
		entity.Error(err)
	case LEVEL_WARN:
		entity.Warn(data.Message)
	case LEVEL_INFO:
		entity.Info(data.Message)
	default:
		return UnknownLogLevel
	}
	return nil
}

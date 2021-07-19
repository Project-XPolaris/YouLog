package api

import (
	"encoding/json"
	"github.com/projectxpolaris/youlog/database"
	"github.com/sirupsen/logrus"
)

var defaultTimeFormat = "2006-01-02 15:04:05"

type BaseLogTemplate struct {
	Application string      `json:"application"`
	Instance    string      `json:"instance"`
	Scope       string      `json:"scope"`
	Level       int64       `json:"level"`
	Message     string      `json:"message"`
	Extra       interface{} `json:"extra"`
	Time        string      `json:"time"`
}

func (t *BaseLogTemplate) Assign(log *database.Log) {
	t.Instance = log.Instance
	t.Application = log.Application
	t.Scope = log.Scope
	t.Level = log.Level
	t.Message = log.Message
	extra := map[string]interface{}{}
	err := json.Unmarshal([]byte(log.Extra), &extra)
	if err != nil {
		logrus.Error(err)
	}
	t.Extra = extra
	t.Time = log.Time.Format(defaultTimeFormat)
}

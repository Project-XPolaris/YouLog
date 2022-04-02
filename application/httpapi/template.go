package httpapi

import (
	"encoding/json"
	"github.com/projectxpolaris/youlog/datasource"
	"github.com/sirupsen/logrus"
)

var defaultTimeFormat = "2006-01-02 15:04:05"

type BaseLogTemplate struct {
	Id          string      `json:"id"`
	Application string      `json:"application"`
	Instance    string      `json:"instance"`
	Scope       string      `json:"scope"`
	Level       int64       `json:"level"`
	Message     string      `json:"message"`
	Extra       interface{} `json:"extra"`
	Time        string      `json:"time"`
}

func (t *BaseLogTemplate) Assign(log datasource.Log) {
	t.Id = log.GetId()
	t.Instance = log.GetInstance()
	t.Application = log.GetApplication()
	t.Scope = log.GetScope()
	t.Level = log.GetLevel()
	t.Message = log.GetMessage()
	if len(log.GetExtra().(string)) != 0 {
		extra := map[string]interface{}{}
		err := json.Unmarshal([]byte(log.GetExtra().(string)), &extra)
		if err != nil {
			logrus.Error(err)
		}
		t.Extra = extra
	}
	t.Time = log.GetTime().Format(defaultTimeFormat)
}

type BaseApplicationTemplate struct {
	Name string `json:"name"`
}

func (t *BaseApplicationTemplate) Assign(log datasource.Log) {
	t.Name = log.GetApplication()
}

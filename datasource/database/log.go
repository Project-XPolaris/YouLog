package database

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Log struct {
	gorm.Model
	Application string
	Instance    string
	Scope       string
	Level       int64
	Message     string
	Extra       string
	Time        time.Time
}

func (l *Log) GetApplication() string {
	return l.Application
}

func (l *Log) GetInstance() string {
	return l.Instance
}

func (l *Log) GetLevel() int64 {
	return l.Level
}

func (l *Log) GetScope() string {
	return l.Scope
}

func (l *Log) GetMessage() string {
	return l.Message
}

func (l *Log) GetExtra() interface{} {
	return l.Extra
}

func (l *Log) GetTime() *time.Time {
	return &l.Time
}

func (l *Log) GetId() string {
	return fmt.Sprintf("%d", l.ID)
}

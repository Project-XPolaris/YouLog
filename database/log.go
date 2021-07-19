package database

import (
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

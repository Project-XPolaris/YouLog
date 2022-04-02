package service

import (
	"github.com/project-xpolaris/youplustoolkit/youlog/logservice"
)

var (
	StdoutOutput = "stdout"
)

type LogOutput interface {
	writeLog(data *logservice.LogData) error
}

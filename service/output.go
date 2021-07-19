package service

import "github.com/projectxpolaris/youlog/pb"

var (
	SqliteOutput = "sqlite"
	StdoutOutput = "stdout"
)

type LogOutput interface {
	writeLog(data *pb.LogData) error
}

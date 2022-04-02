package service

import (
	"github.com/projectxpolaris/youlog/config"
)

func GetLogOutputs() []LogOutput {
	outputs := make([]LogOutput, 0)
	for _, outputConf := range config.Instance.Output {
		switch outputConf {
		case StdoutOutput:
			outputs = append(outputs, &StdoutLogOutput{})
		default:
			for _, db := range DefaultDBOutputs {
				if outputConf == db.Name {
					outputs = append(outputs, db)
				}
			}
		}
	}
	return outputs
}

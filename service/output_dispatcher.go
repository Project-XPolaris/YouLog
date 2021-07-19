package service

import "github.com/projectxpolaris/youlog/config"

func GetLogOutputs() []LogOutput {
	outputs := make([]LogOutput, 0)
	for _, outputConf := range config.Instance.Output {
		switch outputConf {
		case SqliteOutput:
			outputs = append(outputs, &SqliteLogOutput{})
		case StdoutOutput:
			outputs = append(outputs, &StdoutLogOutput{})
		}
	}
	return outputs
}

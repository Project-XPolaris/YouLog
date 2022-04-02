package config

import (
	"github.com/allentom/harukap/config"
	"os"
)

var DefaultConfigProvider *config.Provider
var Instance Config

func InitConfigProvider() error {
	var err error
	customConfigPath := os.Getenv("YOULOG_CONFIG_PATH")
	DefaultConfigProvider, err = config.NewProvider(func(provider *config.Provider) {
		ReadConfig(provider)
	}, customConfigPath)
	return err
}

type Config struct {
	Addr   string
	Output []string
}

func ReadConfig(provider *config.Provider) error {
	provider.Manager.SetDefault("addr", ":50052")
	provider.Manager.SetDefault("api_addr", ":8401")
	provider.Manager.SetDefault("output", []string{"stdout"})
	Instance = Config{
		Addr:   provider.Manager.GetString("rpc_addr"),
		Output: provider.Manager.GetStringSlice("output"),
	}
	return nil
}

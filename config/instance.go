package config

import viper "github.com/spf13/viper"

var Instance Config

type Config struct {
	Addr    string
	ApiAddr string
}

func ReadConfig() error {
	configer := viper.New()
	configer.AddConfigPath("./")
	configer.AddConfigPath("../")
	configer.SetConfigType("yaml")
	configer.SetConfigName("config")
	err := configer.ReadInConfig()
	if err != nil {
		return err
	}
	configer.SetDefault("addr", ":50052")
	configer.SetDefault("api_addr", ":8401")

	Instance = Config{
		Addr:    configer.GetString("addr"),
		ApiAddr: configer.GetString("api_addr"),
	}
	return nil
}

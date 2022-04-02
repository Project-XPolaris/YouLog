package main

import (
	"github.com/allentom/harukap"
	"github.com/allentom/harukap/cli"
	"github.com/projectxpolaris/youlog/application/httpapi"
	"github.com/projectxpolaris/youlog/config"
	"github.com/projectxpolaris/youlog/datasource/database"
	"github.com/projectxpolaris/youlog/plugins"
	"github.com/sirupsen/logrus"
)

func main() {
	err := config.InitConfigProvider()
	if err != nil {
		logrus.Fatal(err)
	}
	appEngine := harukap.NewHarukaAppEngine()
	appEngine.ConfigProvider = config.DefaultConfigProvider
	appEngine.UsePlugin(database.DefaultDatabasePlugin)
	appEngine.UsePlugin(&plugins.InitPlugin{})
	appEngine.HttpService = httpapi.GetEngine()
	if err != nil {
		logrus.Fatal(err)
	}
	appWrap, err := cli.NewWrapper(appEngine)
	if err != nil {
		logrus.Fatal(err)
	}
	appWrap.RunApp()
}

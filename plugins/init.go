package plugins

import (
	"context"
	"github.com/allentom/harukap"
	"github.com/projectxpolaris/youlog/datasource/database"
	"github.com/projectxpolaris/youlog/rpc"
	"github.com/projectxpolaris/youlog/service"
)

type InitPlugin struct {
}

func (p *InitPlugin) OnInit(e *harukap.HarukaAppEngine) error {
	databaseSourceList := e.ConfigProvider.Manager.GetStringMap("datasource")
	for source := range databaseSourceList {
		service.DefaultDBOutputs = append(service.DefaultDBOutputs, &service.DatabaseOutput{
			Name: source,
			DB:   database.DefaultDatabasePlugin.DBS[source],
		})
		database.DatabaseDataSourceList = append(database.DatabaseDataSourceList, &database.DatabaseDataSource{
			Name:     source,
			Instance: database.DefaultDatabasePlugin.DBS[source],
		})
	}
	service.DefaultLogWriter.Run(context.Background())
	go rpc.DefaultLogServer.Run()
	return nil
}

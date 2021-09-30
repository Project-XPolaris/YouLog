package api

import (
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/youlog/datasource"
	"github.com/projectxpolaris/youlog/datasource/sqlite"
	"time"
)

type LogListRequestInput struct {
	LogLevels   []string   `hsource:"query" hname:"level"`
	StartTime   *time.Time `hsource:"query" hname:"startTime"`
	EndTime     *time.Time `hsource:"query" hname:"endTime"`
	Application string     `hsource:"query" hname:"application"`
	Orders      []string   `hsource:"query" hname:"order"`
}

var logListHandler haruka.RequestHandler = func(context *haruka.Context) {
	requestInput := LogListRequestInput{}
	err := context.BindingInput(&requestInput)
	if err != nil {
		AbortError(context, err, 400)
		return
	}
	queryBuilder := datasource.LogListQueryBuilder{}
	queryBuilder.WithPage(context.Param["page"].(int)).WithPageSize(context.Param["pageSize"].(int))
	if requestInput.LogLevels != nil && len(requestInput.LogLevels) != 0 {
		queryBuilder.InLevels(requestInput.LogLevels)
	}
	if requestInput.StartTime != nil {
		queryBuilder.AfterTime(requestInput.StartTime)
	}
	if requestInput.EndTime != nil {
		queryBuilder.BeforeTime(requestInput.EndTime)
	}
	if len(requestInput.Application) > 0 {
		queryBuilder.OfApplication(requestInput.Application)
	}
	if requestInput.Orders != nil {
		queryBuilder.WithOrder(requestInput.Orders)
	}
	count, logList, err := sqlite.DefaultSqliteDataSource.ReadLogs(queryBuilder)
	if err != nil {
		AbortError(context, err, 500)
		return
	}
	data := make([]*BaseLogTemplate, 0)
	for _, logData := range logList {
		template := BaseLogTemplate{}
		template.Assign(logData)
		data = append(data, &template)
	}
	context.JSON(haruka.JSON{
		"success":  true,
		"count":    count,
		"page":     context.Param["page"].(int),
		"pageSize": context.Param["pageSize"].(int),
		"result":   data,
	})
}

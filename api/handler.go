package api

import (
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/youlog/service"
)

var logListHandler haruka.RequestHandler = func(context *haruka.Context) {
	count, logList, err := service.GetLogList(context.Param["page"].(int), context.Param["pageSize"].(int))
	if err != nil {
		AbortError(context, err, 500)
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

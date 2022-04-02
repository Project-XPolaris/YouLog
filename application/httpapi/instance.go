package httpapi

import (
	"github.com/allentom/haruka"
	"github.com/allentom/haruka/middleware"
	"github.com/rs/cors"
)

func GetEngine() *haruka.Engine {
	e := haruka.NewEngine()
	e.UseCors(cors.AllowAll())
	e.UseMiddleware(middleware.NewLoggerMiddleware())
	e.UseMiddleware(middleware.NewPaginationMiddleware("page", "pageSize", 1, 20))
	e.Router.GET("/log", logListHandler)
	e.Router.GET("/log/application", applicationListHandler)
	e.Router.GET("/info", serviceInfoHandler)
	return e
}

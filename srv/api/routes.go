package api

import (
	docs "github.com/ranjbar-dev/backup-api/.swagger/docs"
	apicontroller "github.com/ranjbar-dev/backup-api/srv/api/controllers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *Api) registerSwagger() {

	docs.SwaggerInfo.Title = "Backup api"
	docs.SwaggerInfo.Description = "Responses indicate the success or failure of the request, with HTTP status codes of 200 for success, 400 for bad input, and 500 for errors."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "127.0.0.1"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"https"}

	router := a.hs.GetRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}

func (a *Api) registerRoutes() {

	controller := apicontroller.NewController()

	// register middlewares //

	a.registerMiddlewares()

	// swagger //

	a.registerSwagger()

	// test //

	a.hs.RegisterPostRoute("/store-backup", controller.StoreBackup)

	a.hs.RegisterGetRoute("/test", controller.Test)

}

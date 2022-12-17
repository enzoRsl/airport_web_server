package routes

import (
	"airport_web_server/internal/rest_api/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Set routes
	router := gin.Default()
	apiRouter := router.Group("/api")

	dataTypeRouter := apiRouter.Group("/datatype")
	dataTypeRouter.GET("/", controllers.GetDataTypes)

	airportRouter := apiRouter.Group("/airport")
	airportRouter.GET("/", controllers.GetAirports)

	airportIDRouter := airportRouter.Group("/:airport")
	airportIDRouter.Use(AirportMiddleware)
	airportRangeRouter := airportIDRouter.Group("/range")
	airportRangeRouter.Use(RangeMiddleware)
	return router
}

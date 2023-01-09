package routes

import (
	"airport_web_server/internal/rest_api/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Set routes
	router := gin.Default()
	router.Use(cors.Default())

	apiRouter := router.Group("/api")

	listDataTypeRouter := apiRouter.Group("/datatype")
	listDataTypeRouter.GET("/", controllers.GetDataTypes)

	airportRouter := apiRouter.Group("/airport")
	airportRouter.GET("/", controllers.GetAirports)

	airportIDRouter := airportRouter.Group("/:airport")
	airportIDRouter.Use(AirportMiddleware)

	averageDailyRouter := airportIDRouter.Group("/average")
	averageDailyRouter.Use(DateMiddleware)
	averageDailyRouter.GET("/", controllers.GetAverageMetricsByDay)

	dataTypeRouter := airportIDRouter.Group("/datatype/:datatype")
	dataTypeRouter.Use(DataTypeMiddleware)

	airportRangeRouter := dataTypeRouter.Group("/range")
	airportRangeRouter.Use(RangeMiddleware)
	airportRangeRouter.GET("/", controllers.GetMetricsListInRange)

	return router
}

package controllers

import (
	"airport_web_server/internal/rest_api/models"
	"airport_web_server/internal/rest_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetMetricsListInRange(c *gin.Context) {
	dateDebut := c.MustGet("dateDebut").(time.Time)
	dateFin := c.MustGet("dateFin").(time.Time)
	airport := c.MustGet("airport").(string)
	dataType := c.MustGet("dataType").(string)

	metricsList, err := services.GetMetricsListInRange(dataType, airport, dateDebut, dateFin)
	if err != nil {
		c.JSON(err.Code, err.ErrorMessage)
		return
	}
	if metricsList == nil {
		metricsList = make([]models.Metrics, 0)
	}
	c.JSON(200, metricsList)
}

func GetAverageMetricsByDay(c *gin.Context) {
	dayDate := c.MustGet("date").(time.Time)
	airport := c.MustGet("airport").(string)
	listAverages, err := services.GetAverageMetricsByDay(dayDate, airport)
	if err != nil {
		c.JSON(err.Code, err.ErrorMessage)
		return
	}
	c.JSON(200, listAverages)
}

func GetDataTypes(c *gin.Context) {
	dataTypes, err := services.GetAllDataTypes()
	if err != nil {
		c.JSON(err.Code, err.ErrorMessage)
		return
	}
	c.JSON(http.StatusOK, dataTypes)
}

func GetAirports(c *gin.Context) {
	airports, responseError := services.GetAllAirports()
	if responseError != nil {
		c.JSON(responseError.Code, responseError.ErrorMessage)
		return
	}
	c.JSON(200, airports)
}

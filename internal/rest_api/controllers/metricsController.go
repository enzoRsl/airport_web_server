package controllers

import (
	"airport_web_server/internal/rest_api/services"
	"github.com/gin-gonic/gin"
)

func GetDataTypes(c *gin.Context) {
	dataTypes, responseError := services.GetAllDataTypes()
	if responseError != nil {
		c.JSON(responseError.Code, responseError.ErrorMessage)
		return
	}
	c.JSON(200, dataTypes)
}

func GetAirports(c *gin.Context) {
	airports, responseError := services.GetAllAirports()
	if responseError != nil {
		c.JSON(responseError.Code, responseError.ErrorMessage)
		return
	}
	c.JSON(200, airports)
}

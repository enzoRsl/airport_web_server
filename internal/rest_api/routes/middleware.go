package routes

import (
	"airport_web_server/internal/rest_api/services"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

func AirportMiddleware(c *gin.Context) {
	println("airport middleware")
	airport := c.Param("airport")
	if airport == "" {
		c.JSON(400, gin.H{
			"error": "airport is required",
		})
		c.Abort()
		return
	}

	if len(airport) != 3 {
		c.JSON(400, gin.H{
			"error": "airport must be 3 characters long",
		})
		c.Abort()

		return
	}
	c.Set("airport", airport)
	c.Next()
}

func DataTypeMiddleware(c *gin.Context) {
	println("datatype middleware")
	dataType := c.Param("datatype")
	if dataType == "" {
		c.JSON(400, gin.H{
			"error": "dataType is required",
		})
		c.Abort()
		return
	}
	dataTypeFound, err := services.GetDataType(dataType)
	if err != nil {
		c.JSON(err.Code, err.ErrorMessage)
		c.Abort()
		return
	}
	c.Set("dataType", dataTypeFound)
	c.Next()

}

func RangeMiddleware(c *gin.Context) {
	println("range middleware")
	println(c.Query("dateDebut"))
	dateDebut, err := formatDateTime(c.Query("dateDebut"))
	if err != nil {
		println(err.Error())
		c.JSON(400, gin.H{
			"error": "dateDebut format must be dd/mm/yyyy hh",
		})
		c.Abort()
		return
	}

	dateFin, err := formatDateTime(c.Query("dateFin"))
	if err != nil {
		println(err.Error())
		c.JSON(400, gin.H{
			"error": "dateFin format must be dd/mm/yyyy hh",
		})
		c.Abort()

		return
	}
	c.Set("dateDebut", dateDebut)
	c.Set("dateFin", dateFin)
	c.Next()
}

func DateMiddleware(c *gin.Context) {
	println("date middleware")
	date, err := formatDate(c.Query("date"))
	if err != nil {
		println(err)
		c.JSON(400, gin.H{
			"error": "date format must be dd/mm/yyyy",
		})
		c.Abort()

		return
	}
	c.Set("date", date)
	c.Next()
}

func formatDate(stringDate string) (time.Time, error) {
	println("format date middleware")
	println(stringDate)
	if stringDate == "" {
		return time.Time{}, errors.New("date is required")
	}
	date, err := time.Parse("02/01/2006", stringDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func formatDateTime(stringDate string) (time.Time, error) {
	if stringDate == "" {
		return time.Time{}, errors.New("date is required")
	}
	date, err := time.Parse("02/01/2006 15", stringDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

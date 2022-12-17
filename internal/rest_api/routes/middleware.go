package routes

import "github.com/gin-gonic/gin"

func AirportMiddleware(c *gin.Context) {
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
	c.Next()
}

func RangeMiddleware(c *gin.Context) {
	dateDebut := c.Query("dateDebut")
	if dateDebut == "" {
		c.JSON(400, gin.H{
			"error": "dateDebut is required",
		})
		c.Abort()
		return
	}

	dateFin := c.Query("dateFin")
	if dateFin == "" {
		c.JSON(400, gin.H{
			"error": "dateFin is required",
		})
		c.Abort()
		return
	}
	c.Set("dateDebut", dateDebut)
	c.Set("dateFin", dateFin)
	c.Next()
}

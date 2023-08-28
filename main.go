package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPlants(c *gin.Context) {
	plants := models.getPlants()

	if plants == nil || len(plants) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, plants)
	}
}

func getPlant(c *gin.Context) {
	code := c.Param("code")

	plant := models.GetProduct(code)

	if plant == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, plant)
	}
}

func addPlant(c *gin.Context) {
	var plant models.Plant

	if err := c.BindJSON(&plant); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddPlant(plant)
		c.IndentedJSON(http.StatusCreated, plant)
	}
}

func main() {
	router := gin.Default()
	router.GET("/plants", getPlants)
	router.GET("/plant/:code", getPlant)
	router.POST("/plants", addPlant)
	router.Run("localhost:8083")
}

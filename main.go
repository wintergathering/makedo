package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wintergathering/makedo/controller"
	"github.com/wintergathering/makedo/reviewer"
)

var (
	bathroomReview     reviewer.BathroomReviewer     = reviewer.New()
	bathroomController controller.BathroomController = controller.New(bathroomReview)
)

func main() {
	r := gin.Default()

	r.GET("/bathrooms", func(c *gin.Context) {
		c.JSON(200, bathroomController.FindAll())
	})

	r.POST("/bathrooms", func(c *gin.Context) {
		c.JSON(200, bathroomController.Save(c))
	})

	r.Run()
}

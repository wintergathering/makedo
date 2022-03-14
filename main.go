package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wintergathering/makedo/controller"
	"github.com/wintergathering/makedo/middlewares"
	"github.com/wintergathering/makedo/reviewer"
)

var (
	bathroomReview     reviewer.BathroomReviewer     = reviewer.New()
	bathroomController controller.BathroomController = controller.New(bathroomReview)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	r := gin.New()

	r.Use(gin.Recovery(), middlewares.Logger())

	r.GET("/bathrooms", func(c *gin.Context) {
		c.JSON(200, bathroomController.FindAll())
	})

	r.POST("/bathrooms", func(c *gin.Context) {
		c.JSON(200, bathroomController.Save(c))
	})

	r.Run()
}

//RESUME @8:45 IN MIDDLEWARE VIDEO

package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
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

	r.Static("/css", "./templates/css")

	r.LoadHTMLGlob("templates/*.html")

	r.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	r.GET("/bathrooms", func(c *gin.Context) {
		c.JSON(200, bathroomController.FindAll())
	})

	r.POST("/bathrooms", func(c *gin.Context) {
		err := bathroomController.Save(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Review is valid"})
		}
	})

	r.Run("localhost:8080")
}

//RESUME @ ~4:00 OF HTML TEMPLATES VIDEO

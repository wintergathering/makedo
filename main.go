package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		//RESUME HERE AT LIKE 14:45 IN INTRO GIN VIDEO
	})

	r.Run()
}

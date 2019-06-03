package main

/*
  Example of gin web framework
*/

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8000") // Listen and serve on 0.0.0.0:8000

}

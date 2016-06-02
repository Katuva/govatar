package main

import (
	"github.com/gin-gonic/gin"
	"github.com/katuva/govatar/config"
	"github.com/katuva/govatar/govatar"
)

func main() {
	govatar.Conf = config.LoadConfig()

	govatar.InitDb()
	govatar.CreateUser("test2@test.io", "beebop2")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello!",
		})
	})

	r.Run()
}

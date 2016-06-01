package main

import (
	"github.com/katuva/govatar/govatar"
	"github.com/katuva/govatar/config"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.LoadConfig("")

	govatar.Conf = conf

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

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/katuva/govatar/config"
	"github.com/katuva/govatar/govatar"
)

func main() {
	govatar.Conf = config.LoadConfig()

	govatar.InitDb()
	govatar.CreateUser("test2@test.io", "beebop2")

	temp := govatar.GetUserByHash("1958e9b00a8319f05cb46cf06b397e724c5c66a2ff2633c7f0884d6a57432af8")

	fmt.Printf("%v", temp)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello!",
		})
	})

	r.Run()
}

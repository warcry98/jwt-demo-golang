package main

import (
	"fmt"
	"net/http"
	"warcry98/jwt-demo-golang/controller"
	"warcry98/jwt-demo-golang/middleware"
	"warcry98/jwt-demo-golang/model"

	"github.com/gin-gonic/gin"
)

func init() {
	model.SetDBClient()
}

func main() {
	fmt.Println("Welcome to Go authorization with Go")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Home router",
		})
	})

	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)
	r.GET("/api/v1", middleware.Authorize, controller.Resources)

	r.Run(":5000")
}

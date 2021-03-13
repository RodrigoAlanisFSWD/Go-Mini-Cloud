package main

import (
	"cloud-golang/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/public", "./uploads")

	router.GET("/", func(c *gin.Context) {
		fmt.Fprintln(c.Writer, "Hello!")
	})

	router.POST("/upload", controllers.UploadBreackPoint)
	router.POST("/createdir/:path", controllers.CreateDirBreakPoint)
	router.GET("/read/*path", controllers.ReadDirBreakPoint)

	router.Run()
}

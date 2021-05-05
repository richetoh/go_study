package main

import (
	"github.com/gin-gonic/gin"

	"github.com/getveryrichet/gin_rest_test/controllers" // new
	"github.com/getveryrichet/gin_rest_test/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}

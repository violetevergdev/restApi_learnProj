package app

import (
	db_restapi_dev "restAPI/internal/database/postgres"

	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func init() {
	db_restapi_dev.RestAPIAuth()
}

func Run()  {
	//Серверная часть
	Router.Run("localhost:8080")

	Router.GET("/albums", db_restapi_dev.GetAlbums)
	Router.POST("/albums", db_restapi_dev.CreateAlbum)
}
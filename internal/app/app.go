package app

import (
	db_restapi_dev "restAPI/internal/database/postgres"

	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func Run()  {
	//Серверная часть
	Router.GET("/albums", db_restapi_dev.GetAlbums)
	Router.POST("/albums", db_restapi_dev.CreateAlbum)
	
	Router.Run("localhost:8080")
}
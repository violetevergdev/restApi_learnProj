package db_restapi_dev

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"restAPI/internal/models"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func RestAPIAuth() {
	config := NewConfig()

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBname)


    db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
        panic(err)
    }
    defer db.Close()
  
    err = db.Ping()
    if err != nil {
        panic(err)
    }
  
    fmt.Println("Successfully connected!")

	//Серверная часть 

	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.POST("/albums", CreateAlbum)

	router.Run("localhost:8080")
}

func GetAlbums(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var albums []models.Album

	for rows.Next() {
		var a models.Album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			log.Fatal(err)
		}
		albums = append(albums, a)
	}

	err = rows.Err()
	if err!= nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func CreateAlbum(c *gin.Context) {
	var a models.Album

	if err := c.BindJSON(&a); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalud requesr payload"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO albums(id, title, artist, price) values($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(a.ID, a.Title, a.Artist, a.Price); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, a)
}
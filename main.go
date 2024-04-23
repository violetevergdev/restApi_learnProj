package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Album struct {
	ID     int  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "4228"
    dbname   = "restapi_dev"
)

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
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

	var albums []Album

	for rows.Next() {
		var a Album
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
	var a Album

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
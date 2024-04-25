package db_restapi_dev

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"restAPI/internal/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Инициализируем тип sql.DB
var db *sql.DB

func RestAPIAuth() {
	//Получаем конфигурацию для работы с БД
	config := NewConfig()

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBname)


    db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
       log.Fatal(err)
    }
    defer db.Close()

  // Проверяем соединение с БД
    err = db.Ping()
    if err != nil {
		log.Fatal(err)
    }
  
    fmt.Println("Successfully connected!")
}

func GetAlbums(c *gin.Context) {
	//Указывает формат response 
	c.Header("Content-Type", "application/json")
 fmt.Println(1)
	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
fmt.Println(2)
	var albums []models.Album

	// Итерация по полученым строкам с запроса
	for rows.Next() {
		var a models.Album
		// получаем значения одной строки
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
	// Конвертируем массив в JSON и возвращаем HTTP res
	c.IndentedJSON(http.StatusOK, albums)
}

func CreateAlbum(c *gin.Context) {
	var a models.Album
	
	// Обрабатываем сериализацию JSON для структуры
	if err := c.BindJSON(&a); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Инициализуем запрос
	stmt, err := db.Prepare("INSERT INTO albums(id, title, artist, price) values(default, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// Обрабатываем запрос в БД
	if _, err := stmt.Exec(a.ID, a.Title, a.Artist, a.Price); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, a)
}
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
var DB *sql.DB

// Коннектимся к БД
func DBConnect() {
	//Получаем конфигурацию для работы с БД
	config := NewConfig()

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBname)

    DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
       log.Fatalf("Failed to open db: %s", err)
    }
	
  // Проверяем соединение с БД
    err = DB.Ping()
    if err != nil {
	 	log.Fatalf("Failed to ping db: %s", err)
    }
  
    fmt.Println("Successfully connected!")
}
// Определяем методы работы с БД
func GetAlbums(c *gin.Context) {
	//Указывает формат response 
	c.Header("Content-Type", "application/json")

	rows, err := DB.Query("SELECT * FROM albums")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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
	c.JSON(http.StatusOK, albums)
}

func CreateAlbum(c *gin.Context) {
	var a models.Album
	
	// Обрабатываем сериализацию JSON для структуры
	if err := c.BindJSON(&a); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Инициализуем запрос
	stmt, err := DB.Prepare("INSERT INTO albums(id, title, artist, price) values(default, $1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// Обрабатываем запрос в БД
	if _, err := stmt.Exec(a.Title, a.Artist, a.Price); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, a)
}
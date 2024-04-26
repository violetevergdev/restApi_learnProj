package main

import (
	"restAPI/internal/app"
	db_restapi_dev "restAPI/internal/database/postgres"

	_ "github.com/lib/pq"
)

func init() {
	db_restapi_dev.DBConnect()
}

func main() {
	app.Run()
}

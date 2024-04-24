package main

import (
	db_restapi_dev "restAPI/internal/database/postgres"

	_ "github.com/lib/pq"
)

func init() {
	db_restapi_dev.RestAPIAuth()
}

func main() {
	//

}


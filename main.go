package main

import (
	"encoding/json"

	"github.com/nikhosagala/soapi/db"
	"github.com/nikhosagala/soapi/urls"
)

func main() {
	db.CreateConnection()

	defer db.CloseConnection()

	db.Seed()

	urls.InitializeEngine()

	urls.InitializeRoutes()

	urls.Run() // listen and serve on 0.0.0.0:8080
}

func getJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

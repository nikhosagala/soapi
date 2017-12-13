package main

import (
	"encoding/json"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nikhosagala/soapi/conf"
	"github.com/nikhosagala/soapi/db"
)

func main() {
	db.CreateConnection()

	defer db.CloseConnection()

	db.Seed()

	conf.InitializeEngine()

	conf.InitializeRoutes()

	conf.Run() // listen and serve on 0.0.0.0:8080
}

func getJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

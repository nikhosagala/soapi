package db

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// pg driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is a var to define database from gorm
var DB *gorm.DB

// CreateConnection to database,
// if development, use database name soapi
// if production, use database_url env from Heroku
func CreateConnection() *gorm.DB {
	var migrate = true
	var url = "host=localhost user=soapiuser dbname=soapi sslmode=disable password=soapipass"
	if strings.Compare(os.Getenv("GIN_MODE"), gin.ReleaseMode) == 0 {
		url = os.Getenv("DATABASE_URL")
		migrate = false
	}
	db, err := gorm.Open("postgres", url)

	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	Migrate(migrate)

	return db
}

// CloseConnection function to close connection from database
func CloseConnection() error {
	err := DB.Close()
	return err
}

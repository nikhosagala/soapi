package urls

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// InitializeEngine is function to construct Gin engine
func InitializeEngine() {

	router = gin.Default()

	router.Use(cors.Default())

	router.Use(gzip.Gzip(gzip.BestCompression))
}

// Router return Gin engine
func Router() *gin.Engine {
	return router
}

// Run function to run Gin engine
func Run() {
	router.Run()
}

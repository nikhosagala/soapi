package controllers

import (
	"net/http"

	"github.com/nikhosagala/soapi/utils"

	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/soapi/models"
)

var payload gin.H

// Index to return
func Index(c *gin.Context) {
	utils.Render(c, gin.H{
		"payload": models.Response{Code: http.StatusOK, Data: "API Running"},
	})
}

// Pong to return
func Pong(c *gin.Context) {
	utils.Render(c, gin.H{
		"payload": models.Response{Code: http.StatusOK, Data: "Pong"},
	})
}

// NotFound is function to return if routes not found
func NotFound(c *gin.Context) {
	utils.Render(c, utils.NotFound(" Routes not Found"))
}

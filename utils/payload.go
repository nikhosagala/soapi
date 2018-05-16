package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/soapi/models"
)

// NotFoundPayload return json payload for gin.H
func NotFoundPayload(what string) gin.H {
	return gin.H{"payload": models.Response{Code: http.StatusNotFound, Error: what + " not found", Limit: 0, Offset: 0, Message: "", Data: []models.Response{}}, "code": http.StatusNotFound}
}

// NotFound ...
func NotFound(err string) gin.H {
	return gin.H{"payload": models.Response{Code: http.StatusNotFound, Error: err, Limit: 0, Offset: 0, Message: "", Data: []models.Response{}}, "code": http.StatusNotFound}
}

// Success return defaul json payload for response
func Success(limit int, offset int, data interface{}, message string) gin.H {
	return gin.H{"payload": models.Response{Code: http.StatusOK, Error: "", Limit: limit, Offset: offset, Message: message, Data: data}, "code": http.StatusOK}
}

// SuccessMessage ...
func SuccessMessage(limit int, offset int, message string) gin.H {
	return gin.H{"payload": models.Response{Code: http.StatusOK, Error: "", Limit: limit, Offset: offset, Message: message, Data: []models.Response{}}, "code": http.StatusOK}
}

// Error ...
func Error(err interface{}) gin.H {
	return gin.H{"payload": models.Response{Code: http.StatusBadRequest, Error: err, Limit: 0, Offset: 0, Message: "", Data: []models.Response{}}, "code": http.StatusBadRequest}
}

// BadRequest ...
func BadRequest() gin.H {
	return gin.H{"payload": models.Response{Code: http.StatusBadRequest, Error: []models.Response{}, Limit: 0, Offset: 0, Message: "", Data: []models.Response{}}, "code": http.StatusBadRequest}
}

// ErrorMessage ...
func ErrorMessage(err interface{}, message string) gin.H {
	return gin.H{"payload": models.Response{Code: http.StatusBadRequest, Error: err, Limit: 0, Offset: 0, Message: message, Data: []models.Response{}}, "code": http.StatusBadRequest}
}

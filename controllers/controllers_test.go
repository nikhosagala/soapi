package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/soapi/db"
	"github.com/nikhosagala/soapi/models"
)

func connectDb() {
	db.CreateConnection()
	db.Seed()
}

func closeDb(){
	db.CloseConnection()
}

func TestContextParamsByName(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/test/nikhosagala", nil)
	w := httptest.NewRecorder()
	name := ""

	r := gin.Default()
	r.GET("/api/test/:name", func(c *gin.Context) {
		name = c.Params.ByName("name")
	})

	r.ServeHTTP(w, req)

	if name != "nikhosagala" {
		t.Errorf("Url parameter was not correctly parsed. Should be nikhosagala, was %s.", name)
	}
}

func TestContextJSON(t *testing.T) {
	connectDb()
	defer closeDb()
	req, _ := http.NewRequest("GET", "/api/test/admins", nil)
	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/api/test/admins", ReadAdmins)
	r.ServeHTTP(w, req)
	if w.HeaderMap.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Errorf("Content-Type should be application/json; charset=utf-8, was %s", w.HeaderMap.Get("Content-Type"))
	}
}

func TestInsert(t *testing.T) {
	connectDb()
	defer closeDb()
	admin := models.Admin{Username: "nikhosagala", Password: "1234"}
	json, _ := json.Marshal(admin)
	body := bytes.NewBuffer(json)
	r := gin.Default()
	r.POST("/api/test/admin/create", CreateAdmin)
	req, _ := http.NewRequest("POST", "/api/test/admin/create", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code should be Ok, was: %s", w.Code)
	}

	if w.HeaderMap.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Errorf("Content-Type should be application/json; charset=utf-8, was %s", w.HeaderMap.Get("Content-Type"))
	}
}

func TestUpdate(t *testing.T) {
	connectDb()
	defer closeDb()
	baseModel := models.BaseModel{ID: 1}
	admin := models.Admin{BaseModel: baseModel, Username: "nikhosagala", Password: "1234"}

	json, _ := json.Marshal(admin)
	body := bytes.NewBuffer(json)

	r := gin.Default()
	r.PUT("/api/test/admin/update", UpdateAdmin)

	req, _ := http.NewRequest("PUT", "/api/test/admin/update", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Response code should be NotFound, was: %d", w.Code)
	}

	if w.HeaderMap.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Errorf(w.Body.String())
		t.Errorf("Content-Type should be application/json; charset=utf-8, was %s", w.HeaderMap.Get("Content-Type"))
	}
}

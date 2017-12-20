package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/soapi/db"
	"github.com/nikhosagala/soapi/models"
	"github.com/nikhosagala/soapi/utils"
)

// CreateService ..
func CreateService(c *gin.Context) {
	var service models.Service
	err := c.ShouldBindJSON(&service)
	if err != nil {
		payload = utils.Error(err)
	} else {
		size := len(db.ReadServices(-1, 0))
		service.Code = utils.Code(service.Name, size)
		var serviceSave *models.Service
		serviceSave = db.CreateService(&service)
		payload = utils.Success(0, 0, serviceSave.BaseModel, "Service created.")
	}
	utils.Render(c, payload)
}

// ReadServices ...
func ReadServices(c *gin.Context){
	limit, offset := utils.LimitAndOffset(c.Query("limit"), c.Query("offset"))
	payload = utils.Success(limit, offset, db.ReadServices(limit, offset), "")
	utils.Render(c, payload)
}

// ReadService ...
func ReadService(c *gin.Context){
	id := utils.ParseParam2Int(c.Param("id"))
	service := db.ReadService(id)
	if service.ID == 0 || service.DeletedAt != nil {
		payload = utils.NotFoundPayload("Service")
	}else{
		payload = utils.Success(0, 0, service, "")
	}
	utils.Render(c, payload)
}

// UpdateService ...
func UpdateService(c *gin.Context){
	var service *models.Service
	err := c.ShouldBindJSON(&service)
	if err != nil {
		payload = utils.Error(err)
	} else {
		service = db.UpdateService(service)
		service := db.ReadService(service.ID)
		payload = utils.Success(0, 0, service.BaseModel, "Service updated.")
	}
	utils.Render(c, payload)
}

// DeleteService ...
func DeleteService(c *gin.Context){
	id := utils.ParseParam2Int(c.Param("id"))
	service := db.ReadService(id)
	if service.ID == 0 || service.DeletedAt != nil {
		payload = utils.NotFoundPayload("Service")
	} else {
		db.DeleteService(&service)
		payload = utils.Success(0, 0, service.BaseModel, "Service deleted.")
	}
	utils.Render(c, payload)
}

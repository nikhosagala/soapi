package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/soapi/db"
	"github.com/nikhosagala/soapi/models"
	"github.com/nikhosagala/soapi/utils"
)

// CreateAdmin ...
func CreateAdmin(c *gin.Context) {
	var admin models.Admin
	err := c.BindJSON(&admin)
	if err != nil {
		payload = utils.Error(err)
	} else {
		var adminSave *models.Admin
		adminSave, err2 := db.CreateAdmin(&admin)
		if err2 != nil {
			payload = utils.ErrorMessage(err2.Error(), "Username "+admin.Username+" already exist.")
		} else {
			payload = utils.Success(0, 0, adminSave.BaseModel, "Admin created")
		}
	}
	utils.Render(c, payload)
}

// ReadAdmins ...
func ReadAdmins(c *gin.Context) {
	limit, offset := utils.LimitAndOffset(c.Query("limit"), c.Query("offset"))
	payload = utils.Success(limit, offset, db.ReadAllAdmins(limit, offset), "")
	utils.Render(c, payload)
}

// ReadAdmin ...
func ReadAdmin(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	admin := db.ReadAdmin(id)
	if utils.CheckIDAndDeleted(admin.BaseModel) {
		payload = utils.NotFoundPayload("Admin")
	} else {
		payload = utils.Success(0, 0, admin, "")
	}
	utils.Render(c, payload)
}

// UpdateAdmin ...
func UpdateAdmin(c *gin.Context) {
	var admin models.Admin
	err := c.BindJSON(&admin)
	if err != nil {
		payload = utils.Error(err)
	} else {
		var adminUpdate *models.Admin
		adminUpdate, err2 := db.UpdateAdmin(&admin)
		if utils.CheckIDAndDeleted(adminUpdate.BaseModel) {
			payload = utils.NotFoundPayload("Admin")
		} else if err2 != nil {
			payload = utils.ErrorMessage(err2.Error(), "Username "+admin.Username+" already exist.")
		} else {
			// adminUpdate := db.ReadAdmin(adminUpdate.ID)
			payload = utils.Success(0, 0, adminUpdate.BaseModel, "Admin updated.")
		}
	}
	utils.Render(c, payload)
}

// DeleteAdmin ...
func DeleteAdmin(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	admin := db.ReadAdmin(id)
	if admin.ID == 0 || admin.DeletedAt != nil {
		payload = utils.NotFoundPayload("Admin")
	} else {
		db.DeleteAdmin(&admin)
		payload = utils.Success(0, 0, admin.BaseModel, "Admin deleted.")
	}
	utils.Render(c, payload)
}

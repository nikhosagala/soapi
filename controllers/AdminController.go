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
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		payload = utils.Error(err)
	} else {
		var adminSave *models.Admin
		admin.Password = utils.CreatePassword(admin.Password)
		admin.Token = utils.Token()
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
	if admin.ID == 0 || admin.DeletedAt != nil {
		payload = utils.NotFoundPayload("Admin")
	} else {
		payload = utils.Success(0, 0, admin, "")
	}
	utils.Render(c, payload)
}

// UpdateAdmin ...
func UpdateAdmin(c *gin.Context) {
	var admin *models.Admin
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		payload = utils.Error(err)
	} else {
		admin, err2 := db.UpdateAdmin(admin)
		if err2 != nil {
			payload = utils.ErrorMessage(err2.Error(), "Username "+admin.Username+" already exist.")
		} else {
			adminUpdate := db.ReadAdmin(admin.ID)
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

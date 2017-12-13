package db

import (
	"errors"

	"github.com/nikhosagala/soapi/models"
)

// CreateAdmin ...
func CreateAdmin(admin *models.Admin) (*models.Admin, error) {
	var checkAdmin models.Admin
	DB.NewRecord(admin)
	DB.Where("username = ?", admin.Username).First(&checkAdmin)
	if checkAdmin.ID != 0 {
		return admin, errors.New("username already exist")
	}
	DB.Create(&admin)
	return admin, nil
}

// ReadAllAdmins ...
func ReadAllAdmins(limit int, offset int) []models.Admin {
	admins := []models.Admin{}
	DB.Limit(limit).Offset(offset).Find(&admins)
	return admins
}

// ReadAdmin ...
func ReadAdmin(id uint) models.Admin {
	admin := models.Admin{}
	DB.Find(&admin, id)
	return admin
}

// UpdateAdmin ...
func UpdateAdmin(admin *models.Admin) (*models.Admin, error) {
	var update models.Admin
	search := ReadAdmin(admin.ID)
	DB.Where("username = ?", admin.Username).First(&update)
	if update.ID != 0 && search.Username != update.Username {
		return admin, errors.New("username already exist")
	}
	DB.Model(&admin).Updates(admin)
	return admin, nil
}

// DeleteAdmin ...
func DeleteAdmin(admin *models.Admin) *models.Admin {
	DB.Delete(&admin)
	return admin
}

package db

import (
	"errors"

	"github.com/nikhosagala/soapi/models"
	"github.com/nikhosagala/soapi/utils"
)

// CreateAdmin ...
func CreateAdmin(admin *models.Admin) (*models.Admin, error) {
	var checkAdmin models.Admin
	DB.NewRecord(admin)
	DB.Where("username = ?", admin.Username).First(&checkAdmin)
	if checkAdmin.ID != 0 {
		return admin, errors.New("username already exist")
	} else if admin.Password == "" || admin.Username == "" {
		return admin, errors.New("password or username can not be empty")
	}
	admin.Password = utils.CreatePassword(admin.Password)
	admin.Token = utils.Token()
	err := DB.Create(&admin).Error
	if err != nil {
		return admin, err
	}
	return admin, nil
}

// ReadAllAdmins ...
func ReadAllAdmins(limit int, offset int) []models.Admin {
	if limit == 0 {
		limit = -1
	}
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
	if search.ID == 0 {
		return &search, errors.New("admin not found")
	} else if update.ID != 0 {
		return &search, errors.New("username already exist")
	}
	if admin.Password != "" {
		admin.Password = utils.CreatePassword(admin.Password)
	}
	DB.Model(&admin).Updates(admin)
	return admin, nil
}

// DeleteAdmin ...
func DeleteAdmin(admin *models.Admin) *models.Admin {
	DB.Delete(&admin)
	return admin
}

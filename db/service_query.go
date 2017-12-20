package db

import "github.com/nikhosagala/soapi/models"

// CreateService ...
func CreateService(service *models.Service) *models.Service {
	DB.NewRecord(service)
	DB.Create(&service)
	return service
}

// ReadServices ...
func ReadServices(limit int, offset int) []models.Service {
	services := []models.Service{}
	DB.Limit(limit).Offset(offset).Find(&services)
	return services
}

// ReadService ...
func ReadService(id uint) models.Service {
	service := models.Service{}
	DB.Find(&service, id)
	return service
}

// UpdateService ...
func UpdateService(service *models.Service) *models.Service {
	DB.Model(&service).Updates(service)
	return service
}

// DeleteService ...
func DeleteService(service *models.Service) *models.Service {
	DB.Delete(&service)
	return service
}

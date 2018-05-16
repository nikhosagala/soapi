package db

import "github.com/nikhosagala/soapi/models"
import "log"

// Migrate function to migrate model from model package
func Migrate(migrate bool) {
	DropTables(migrate)
	log.Println("Create tables")
	DB.AutoMigrate(&models.Admin{}, &models.Branch{}, &models.OrderDetails{}, &models.Company{})
	DB.AutoMigrate(&models.Customer{}, &models.Order{}, &models.Service{}, &models.Session{})
}

// DropTables ...
func DropTables(migrate bool) {
	if migrate {
		log.Println("Migrating")
		log.Println("Drop tables")
		DB.DropTableIfExists(&models.Admin{})
		DB.DropTableIfExists(&models.Branch{})
		DB.DropTableIfExists(&models.OrderDetails{})
		DB.DropTableIfExists(&models.Company{})
		DB.DropTableIfExists(&models.Customer{})
		DB.DropTableIfExists(&models.Order{})
		DB.DropTableIfExists(&models.Service{})
		DB.DropTableIfExists(&models.Session{})
	}
}

// Seed ...
func Seed() {
	var company models.Company
	var branch models.Branch
	branch = models.Branch{BasicInfo: models.BasicInfo{Name: "Liquid 1", Address: "Sukabirus", Phone: "0221230424"}}
	DB.Create(&branch)
	company = models.Company{BasicInfo: models.BasicInfo{Name: "Liquid", Address: "Sukapura", Phone: "0221230423"}, Branches: []models.Branch{branch}}
	DB.Create(&company)
	branch = models.Branch{BasicInfo: models.BasicInfo{Name: "Liquid 2", Address: "Dayeuhkolot", Phone: "0221230425"}, CompanyID: company.ID}
	DB.Create(&branch)
}

package db

import "github.com/nikhosagala/soapi/models"

// ReadAllCompanies ...
func ReadAllCompanies(limit int, offset int) []models.Company {
	companies := []models.Company{}
	DB.Limit(limit).Offset(offset).Preload("Branches").Find(&companies)
	return companies
}

// ReadCompany ...
func ReadCompany(id uint) models.Company {
	company := models.Company{}
	DB.Preload("Branches").Find(&company, id)
	return company
}

// ReadBranches ...
func ReadBranches(limit int, offset int) []models.Branch {
	branches := []models.Branch{}
	DB.Limit(limit).Offset(offset).Preload("Admin").Find(&branches)
	return branches
}

// ReadBranch ...
func ReadBranch(id uint) models.Branch {
	branch := models.Branch{}
	DB.Find(&branch, id)
	return branch
}

// ReadCustomers ...
func ReadCustomers(limit int, offset int) []models.Customer {
	customers := []models.Customer{}
	DB.Limit(limit).Offset(offset).Find(&customers)
	return customers
}

// ReadCustomer ...
func ReadCustomer(id uint) models.Customer {
	customer := models.Customer{}
	DB.Find(&customer, id)
	return customer
}

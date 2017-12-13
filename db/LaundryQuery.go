package db

import "github.com/nikhosagala/soapi/models"

// CreateCompany ...
func CreateCompany(company *models.Company) *models.Company {
	DB.NewRecord(company)
	DB.Create(&company)
	return company
}

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

// UpdateCompany ...
func UpdateCompany(company *models.Company) *models.Company {
	DB.Model(&company).Updates(company)
	return company
}

// DeleteCompany ...
func DeleteCompany(company *models.Company) *models.Company {
	DB.Delete(&company)
	return company
}

// CreateBranch ...
func CreateBranch(branch *models.Branch) *models.Branch {
	DB.NewRecord(branch)
	DB.Create(&branch)
	return branch
}

// ReadAllBranches ...
func ReadAllBranches(limit int, offset int) []models.Branch {
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

// UpdateBranch ...
func UpdateBranch(branch *models.Branch) *models.Branch {
	DB.Model(&branch).Updates(branch)
	return branch
}

// DeleteBranch ...
func DeleteBranch(branch *models.Branch) *models.Branch {
	DB.Delete(&branch)
	return branch
}

// CreateCustomer ...
func CreateCustomer(customer *models.Customer) *models.Customer {
	DB.NewRecord(customer)
	DB.Create(&customer)
	return customer
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

// UpdateCustomer ...
func UpdateCustomer(customer *models.Customer) *models.Customer {
	DB.Model(&customer).Updates(customer)
	return customer
}

// DeleteCustomer ...
func DeleteCustomer(customer *models.Customer) *models.Customer {
	DB.Delete(&customer)
	return customer
}

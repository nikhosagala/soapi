package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/soapi/db"
	"github.com/nikhosagala/soapi/models"
	"github.com/nikhosagala/soapi/utils"
)

// CreateCompany ...
func CreateCompany(c *gin.Context) {
	var company models.Company
	err := c.BindJSON(&company)
	if err != nil {
		payload = utils.Error(err)
	} else {
		company := db.Create(&company).(*models.Company)
		payload = utils.Success(0, 0, company.BaseModel, "Company created.")
	}
	utils.Render(c, payload)
}

// ReadCompanies ...
func ReadCompanies(c *gin.Context) {
	limit, offset := utils.LimitAndOffset(c.Query("limit"), c.Query("offset"))
	payload = utils.Success(limit, offset, db.ReadAllCompanies(limit, offset), "")
	utils.Render(c, payload)
}

// ReadCompany ...
func ReadCompany(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	company := db.ReadCompany(id)
	if company.ID == 0 {
		payload = utils.NotFoundPayload("Company")
	} else {
		payload = utils.Success(0, 0, company, "")
	}
	utils.Render(c, payload)
}

// UpdateCompany ...
func UpdateCompany(c *gin.Context) {
	var company models.Company
	err := c.BindJSON(&company)
	if err != nil {
		payload = utils.Error(err)
	} else {
		company := db.Update(&company).(*models.Company)
		companyUpdated := db.ReadCompany(company.ID)
		if utils.CheckIDAndDeleted(company.BaseModel) {
			payload = utils.NotFoundPayload("Company")
		} else {
			payload = utils.Success(0, 0, companyUpdated.BaseModel, "Company updated.")
		}
	}
	utils.Render(c, payload)
}

// DeleteCompany ...
func DeleteCompany(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	company := db.ReadCompany(id)
	if utils.CheckIDAndDeleted(company.BaseModel) {
		payload = utils.NotFoundPayload("Company")
	} else {
		db.Delete(&company)
		payload = utils.Success(0, 0, company.BaseModel, "Company deleted.")
	}
	utils.Render(c, payload)
}

// CreateBranch ...
func CreateBranch(c *gin.Context) {
	var branch models.Branch
	err := c.BindJSON(&branch)
	if err != nil {
		payload = utils.Error(err)
	} else {
		branch := db.Create(&branch).(*models.Branch)
		payload = utils.Success(0, 0, branch.BaseModel, "Branch created.")
	}
	utils.Render(c, payload)
}

// ReadBranches ...
func ReadBranches(c *gin.Context) {
	limit, offset := utils.LimitAndOffset(c.Query("limit"), c.Query("offset"))
	payload = utils.Success(limit, offset, db.ReadBranches(limit, offset), "")
	utils.Render(c, payload)
}

// ReadBranch ...
func ReadBranch(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	branch := db.ReadBranch(id)
	if branch.ID == 0 {
		payload = utils.NotFoundPayload("Branch")
	} else {
		payload = utils.Success(0, 0, branch, "")
	}
	utils.Render(c, payload)
}

// UpdateBranch ...
func UpdateBranch(c *gin.Context) {
	var branch models.Branch
	err := c.BindJSON(&branch)
	if err != nil {
		payload = utils.Error(err)
	} else {
		branch := db.Update(&branch).(*models.Branch)
		branchUpdated := db.ReadBranch(branch.ID)
		if utils.CheckIDAndDeleted(branch.BaseModel) {
			payload = utils.NotFoundPayload("Branch")
		} else {
			payload = utils.Success(0, 0, branchUpdated.BaseModel, "Branch updated.")
		}
	}
	utils.Render(c, payload)
}

// DeleteBranch ...
func DeleteBranch(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	branch := db.ReadBranch(id)
	if utils.CheckIDAndDeleted(branch.BaseModel) {
		payload = utils.NotFoundPayload("Branch")
	} else {
		db.Delete(&branch)
		payload = utils.Success(0, 0, branch.BaseModel, "Branch deleted.")
	}
	utils.Render(c, payload)
}

// CreateCustomer ...
func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		payload = utils.Error(err)
	} else {
		if customer.Password != "" {
			customer.Password = utils.CreatePassword(customer.Password)
		}
		customer := db.Create(&customer).(*models.Customer)
		payload = utils.Success(0, 0, customer.BaseModel, "Customer created.")
	}
	utils.Render(c, payload)
}

// ReadCustomers ...
func ReadCustomers(c *gin.Context) {
	limit, offset := utils.LimitAndOffset(c.Query("limit"), c.Query("offset"))
	payload = utils.Success(limit, offset, db.ReadCustomers(limit, offset), "")
	utils.Render(c, payload)
}

// ReadCustomer ...
func ReadCustomer(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	customer := db.ReadCustomer(id)
	if customer.ID == 0 {
		payload = utils.NotFoundPayload("Customer")
	} else {
		payload = utils.Success(0, 0, customer, "")
	}
	utils.Render(c, payload)
}

// UpdateCustomer ...
func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		payload = utils.Error(err)
	} else {
		customer := db.Update(&customer).(*models.Customer)
		customerUpdated := db.ReadCustomer(customer.ID)
		if utils.CheckIDAndDeleted(customer.BaseModel) {
			payload = utils.NotFoundPayload("Customer")
		} else {
			payload = utils.Success(0, 0, customerUpdated.BaseModel, "Customer updated.")
		}
	}
	utils.Render(c, payload)
}

// DeleteCustomer ...
func DeleteCustomer(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	customer := db.ReadCustomer(id)
	if utils.CheckIDAndDeleted(customer.BaseModel) {
		payload = utils.NotFoundPayload("Customer")
	} else {
		db.Delete(&customer)
		payload = utils.Success(0, 0, customer.BaseModel, "Customer deleted.")
	}
	utils.Render(c, payload)
}

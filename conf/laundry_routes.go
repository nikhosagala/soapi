package conf

import "github.com/nikhosagala/soapi/controllers"

// InitializeLaundryRoutes ...
func InitializeLaundryRoutes() {

	router.POST("/api/company/new", controllers.CreateCompany)

	router.GET("/api/companies", controllers.ReadCompanies)

	router.GET("/api/company/:id", controllers.ReadCompany)

	router.PUT("/api/company/update", controllers.UpdateCompany)

	router.DELETE("/api/company/:id", controllers.DeleteCompany)

	router.POST("/api/branch/new", controllers.CreateBranch)

	router.GET("/api/branches", controllers.ReadBranches)

	router.GET("/api/branch/:id", controllers.ReadBranch)

	router.PUT("/api/branch/update", controllers.UpdateBranch)

	router.DELETE("/api/branch/:id", controllers.DeleteBranch)

	router.POST("/api/customer/new", controllers.CreateCustomer)

	router.GET("/api/customers", controllers.ReadCustomers)

	router.GET("/api/customer/:id", controllers.ReadCustomer)

	router.PUT("/api/customer/update", controllers.UpdateCustomer)

	router.DELETE("/api/customer/:id", controllers.DeleteCustomer)

}

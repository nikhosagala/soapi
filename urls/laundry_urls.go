package urls

import "github.com/nikhosagala/soapi/controllers"

// InitializeLaundryRoutes ...
func InitializeLaundryRoutes() {

	router.POST("/company/new", controllers.CreateCompany)

	router.GET("/companies", controllers.ReadCompanies)

	router.GET("/company/:id", controllers.ReadCompany)

	router.PUT("/company/update", controllers.UpdateCompany)

	router.DELETE("/company/:id", controllers.DeleteCompany)

	router.POST("/branch/new", controllers.CreateBranch)

	router.GET("/branches", controllers.ReadBranches)

	router.GET("/branch/:id", controllers.ReadBranch)

	router.PUT("/branch/update", controllers.UpdateBranch)

	router.DELETE("/api/branch/:id", controllers.DeleteBranch)

	router.POST("/api/customer/new", controllers.CreateCustomer)

	router.GET("/api/customers", controllers.ReadCustomers)

	router.GET("/api/customer/:id", controllers.ReadCustomer)

	router.PUT("/api/customer/update", controllers.UpdateCustomer)

	router.DELETE("/api/customer/:id", controllers.DeleteCustomer)

}

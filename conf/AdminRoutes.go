package conf

import "github.com/nikhosagala/soapi/controllers"

// InitializeAdminRoutes ...
func InitializeAdminRoutes() {
	router.POST("/api/admin/new", controllers.CreateAdmin)

	router.GET("/api/admins", controllers.ReadAdmins)

	router.GET("/api/admin/:id", controllers.ReadAdmin)

	router.PUT("/api/admin/update", controllers.UpdateAdmin)

	router.DELETE("/api/admin/:id", controllers.DeleteAdmin)
}

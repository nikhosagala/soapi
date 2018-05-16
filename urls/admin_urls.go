package urls

import "github.com/nikhosagala/soapi/controllers"

// InitializeAdminRoutes ...
func InitializeAdminRoutes() {
	router.POST("/admin/new", controllers.CreateAdmin)

	router.GET("/admins", controllers.ReadAdmins)

	router.GET("/admin/:id", controllers.ReadAdmin)

	router.PUT("/admin/update", controllers.UpdateAdmin)

	router.DELETE("/admin/:id", controllers.DeleteAdmin)
}

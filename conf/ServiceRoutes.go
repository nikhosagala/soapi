package conf

import "github.com/nikhosagala/soapi/controllers"

// InitializeServiceRoutes ...
func InitializeServiceRoutes() {
	router.POST("/api/service/new", controllers.CreateService)

	router.GET("/api/services", controllers.ReadServices)

	router.GET("/api/service/:id", controllers.ReadService)

	router.PUT("/api/service/update", controllers.UpdateService)

	router.DELETE("/api/service/:id", controllers.DeleteService)
}

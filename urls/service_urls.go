package urls

import "github.com/nikhosagala/soapi/controllers"

// InitializeServiceRoutes ...
func InitializeServiceRoutes() {
	router.POST("/service/new", controllers.CreateService)

	router.GET("/services", controllers.ReadServices)

	router.GET("/service/:id", controllers.ReadService)

	router.PUT("/service/update", controllers.UpdateService)

	router.DELETE("/service/:id", controllers.DeleteService)
}

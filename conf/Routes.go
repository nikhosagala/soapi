package conf

import "github.com/nikhosagala/soapi/controllers"

// InitializeRoutes function to define routing
func InitializeRoutes() {

	router.GET("/api", controllers.Index)

	router.GET("/api/ping", controllers.Pong)

	router.NoRoute(controllers.NotFound)

	InitializeLaundryRoutes()

	InitializeAdminRoutes()

	InitializeServiceRoutes()

	InitializeOrderRoutes()
}

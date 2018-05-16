package urls

import "github.com/nikhosagala/soapi/controllers"

// InitializeRoutes function to define routing
func InitializeRoutes() {

	//router.GET("/", controllers.Index)

	//router.GET("/ping", controllers.Pong)

	router.NoRoute(controllers.NotFound)

	//InitializeLaundryRoutes()

	//InitializeAdminRoutes()

	//InitializeServiceRoutes()

	//InitializeOrderRoutes()
}

package urls

import "github.com/nikhosagala/soapi/controllers"

// InitializeOrderRoutes ...
func InitializeOrderRoutes() {
	router.POST("/order/new", controllers.CreateOrder)

	router.GET("/orders", controllers.ReadOrders)

	router.GET("/order/:id", controllers.ReadOrder)

	router.PUT("/order/update", controllers.UpdateOrder)

	router.DELETE("/order/:id", controllers.DeleteOrder)
}

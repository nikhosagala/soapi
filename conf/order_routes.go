package conf

import "github.com/nikhosagala/soapi/controllers"

// InitializeOrderRoutes ...
func InitializeOrderRoutes() {
	router.POST("/api/order/new", controllers.CreateOrder)
	
	router.GET("/api/orders", controllers.ReadOrders)

	router.GET("/api/order/:id", controllers.ReadOrder)

	router.PUT("/api/order/update", controllers.UpdateOrder)

	router.DELETE("/api/order/:id", controllers.DeleteOrder)
}

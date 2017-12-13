package models

// Service ...
type Service struct {
	BaseModel
	Name  string  `json:"name" binding:"required"`
	Code  string  `json:"code"`
	Price float32 `json:"price" binding:"required"`
}

// Customer ...
type Customer struct {
	BaseModel
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

// OrderDetails ...
type OrderDetails struct {
	BaseModel
	Weight    float32 `json:"weight" binding:"required"`
	Service   Service `json:"service"`
	ServiceID uint    `json:"service_id"`
	OrderID   uint    `json:"order_id" binding:"required"`
	SubTotal  float32 `json:"sub_total"`
}

// Order ...
type Order struct {
	BaseModel
	Name         string         `json:"name"`
	Customer     Customer       `json:"customer"`
	CustomerID   uint           `json:"customer_id"`
	OrderDetails []OrderDetails `json:"order_details"`
	TotalPrice   float32        `json:"total_price"`
}

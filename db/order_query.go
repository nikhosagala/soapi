package db

import (
	"errors"

	"github.com/nikhosagala/soapi/models"
)

// CreateOrder ...
func CreateOrder(order *models.Order) (*models.Order, error) {
	var err error
	var err1 error
	var err2 error
	var totalPrice float32
	tx := DB.Begin()
	tx.NewRecord(order)
	if ReadCustomer(order.CustomerID).ID == 0 {
		err1 = errors.New("customer not exist")
		err = err1
	}
	for i, v := range order.OrderDetails {
		service := ReadService(v.ServiceID)
		if service.ID == 0 {
			err2 = errors.New("service not exist")
			err = err2
		}
		subTotal := v.Weight * service.Price
		order.OrderDetails[i].Service = service
		order.OrderDetails[i].SubTotal = subTotal
		totalPrice += subTotal
	}
	if err1 != nil && err2 != nil {
		err = errors.New("customer and service not exist")
	}
	order.TotalPrice = totalPrice
	tx.Save(&order)
	if tx.Error != nil || err != nil {
		tx.Rollback()
		return order, err
	}
	tx.Commit()
	return order, nil
}

// ReadOrders ...
func ReadOrders(limit int, offset int) []models.Order {
	orders := []models.Order{}
	DB.Limit(limit).Offset(offset).Preload("Customer").Preload("OrderDetails").Preload("OrderDetails.Service").Find(&orders)
	return orders
}

// ReadOrder ...
func ReadOrder(id uint) models.Order {
	order := models.Order{}
	DB.Preload("Customer").Preload("OrderDetails").Preload("OrderDetails.Service").Find(&order, id)
	return order
}

// UpdateOrder ...
func UpdateOrder(order *models.Order) *models.Order {
	tx := DB.Begin()
	tx.Model(&order).Updates(order)
	if tx.Error != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return order
}

// DeleteOrder ...
func DeleteOrder(order *models.Order) *models.Order {
	tx := DB.Begin()
	tx.Delete(&order)
	tx.Commit()
	return order
}

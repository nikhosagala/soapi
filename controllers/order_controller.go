package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhosagala/soapi/db"
	"github.com/nikhosagala/soapi/models"
	"github.com/nikhosagala/soapi/utils"
)

// CreateOrder ...
func CreateOrder(c *gin.Context) {
	var order models.Order
	err := c.BindJSON(&order)
	if err != nil {
		payload = utils.Error(err)
	} else {
		var orderSave *models.Order
		orderSave, err2 := db.CreateOrder(&order)
		if err2 != nil {
			payload = utils.ErrorMessage(err2, err2.Error())
		} else {
			payload = utils.Success(0, 0, orderSave.BaseModel, "Order created.")
		}
	}
	utils.Render(c, payload)
}

// ReadOrders ...
func ReadOrders(c *gin.Context) {
	limit, offset := utils.LimitAndOffset(c.Query("limit"), c.Query("offset"))
	payload = utils.Success(limit, offset, db.ReadOrders(limit, offset), "")
	utils.Render(c, payload)
}

// ReadOrder ...
func ReadOrder(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	order := db.ReadOrder(id)
	if order.ID == 0 || order.DeletedAt != nil {
		payload = utils.NotFoundPayload("Order")
	} else {
		payload = utils.Success(0, 0, order, "")
	}
	utils.Render(c, payload)
}

// UpdateOrder ...
func UpdateOrder(c *gin.Context) {
	var order *models.Order
	err := c.BindJSON(&order)
	if err != nil {
		payload = utils.Error(err)
	} else {
		order = db.UpdateOrder(order)
		order := db.ReadOrder(order.ID)
		payload = utils.Success(0, 0, order.BaseModel, "Order updated.")
	}
	utils.Render(c, payload)
}

// DeleteOrder ...
func DeleteOrder(c *gin.Context) {
	id := utils.ParseParam2Int(c.Param("id"))
	order := db.ReadOrder(id)
	if order.ID == 0 || order.DeletedAt != nil {
		payload = utils.NotFoundPayload("Order")
	} else {
		db.DeleteOrder(&order)
		payload = utils.Success(0, 0, order.BaseModel, "Order deleted.")
	}
	utils.Render(c, payload)
}

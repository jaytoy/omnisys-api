package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"omnisys.io/core/application"
	"omnisys.io/core/domain"
)

type OrderController struct {
	orderApp application.OrderApp
}

func CreateOrderController(r *gin.Engine, orderApp application.OrderApp) {
	orderController := OrderController{orderApp}

	r.POST("/order", orderController.addOrder)
	r.GET("/order/:id", orderController.viewOrderById)
}

func (e *OrderController) addOrder(c *gin.Context) {
	var order = domain.Order{}
	err := c.Bind(&order)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if len(order.OrderItems) == 0 {
		HandleError(c, http.StatusBadRequest, "empty order")
		return
	}
	newOrder, err := e.orderApp.Create(&order)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	HandleSucces(c, newOrder)
}

func (e *OrderController) viewOrderById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	order, err := e.orderApp.ViewById(id)
	if err != nil {
		HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	HandleSucces(c, order)
}

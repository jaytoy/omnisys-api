package main

import (
	"github.com/gin-gonic/gin"
	"omnisys.io/core/application"
	"omnisys.io/core/infrastructure/database"
	"omnisys.io/core/interfaces/controller"
	"omnisys.io/core/interfaces/repository"
)

func main() {
	// register the database
	db := database.ConnectDB()

	// register the router
	router := gin.Default()

	// register customer module
	customerRepo := repository.CreateCustomerRepository(db)
	customerApp := application.CreateCustomerApp(customerRepo)
	controller.CreateCustomerController(router, customerApp)

	// register product module
	productRepo := repository.CreateProductRepository(db)
	productApp := application.CreateProductApp(productRepo)
	controller.CreateProductController(router, productApp)

	// register order module
	orderRepo := repository.CreateOrderRepository(db)
	orderApp := application.CreateOrderApp(orderRepo)
	controller.CreateOrderController(router, orderApp)

	// start server
	router.Run("localhost:8080")
}

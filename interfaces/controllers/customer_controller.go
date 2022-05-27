package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"omnisys.io/core/application"
	"omnisys.io/core/domain"
)

type CustomerController struct {
	customerPort application.CustomerPort
}

func CreateCustomerController(r *gin.Engine, customerPort application.CustomerPort) {
	customerController := CustomerController{customerPort}

	r.POST("/customer", customerController.addCustomer)
	r.GET("/customer", customerController.viewCustomers)
	r.GET("customer/:id", customerController.viewCustomerById)
	r.PUT("/customer/:id", customerController.editCustomer)
	r.DELETE("/customer/:id", customerController.deleteCustomer)
}

func (e *CustomerController) addCustomer(c *gin.Context) {
	var customer = domain.Customer{}
	err := c.Bind(&customer)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if customer.ID != 0 {
		HandleError(c, http.StatusBadRequest, "input not permitted")
		return
	}

	if customer.Name == "" {
		HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	newPerson, err := e.customerPort.Create(&customer)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	HandleSucces(c, newPerson)
}

func (e *CustomerController) viewCustomers(c *gin.Context) {
	persons, err := e.customerPort.ViewAll()
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*persons) == 0 {
		HandleError(c, http.StatusNotFound, "list customer is empty")
		return
	}
	HandleSucces(c, persons)
}

func (e *CustomerController) viewCustomerById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	customer, err := e.customerPort.ViewById(id)
	if err != nil {
		HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	HandleSucces(c, customer)
}

func (e *CustomerController) editCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	_, err = e.customerPort.ViewById(id)
	if err != nil {
		HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	var tempPerson = domain.Customer{}
	err = c.Bind(&tempPerson)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if tempPerson.ID != 0 {
		HandleError(c, http.StatusBadRequest, "input not permitted")
		return
	}
	if tempPerson.Name == "" {
		HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	customer, err := e.customerPort.Edit(id, &tempPerson)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	HandleSucces(c, customer)
}

func (e *CustomerController) deleteCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	err = e.customerPort.Delete(id)
	if err != nil {
		HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	HandleSucces(c, "success delete data")
}

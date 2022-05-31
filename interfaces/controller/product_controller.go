package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"omnisys.io/core/application"
	"omnisys.io/core/domain"
)

type ProductController struct {
	productApp application.ProductApp
}

func CreateProductController(r *gin.Engine, productApp application.ProductApp) {
	productController := ProductController{productApp}

	r.POST("/product", productController.addProduct)
	r.GET("/products", productController.viewProducts)
	r.GET("/product/:id", productController.viewProductById)
	r.PUT("/product/:id", productController.editProduct)
	r.DELETE("/product/:id", productController.deleteProduct)
}

func (e *ProductController) addProduct(c *gin.Context) {
	var product = domain.Product{}
	err := c.Bind(&product)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if product.ID != 0 {
		HandleError(c, http.StatusBadRequest, "input not permitted")
		return
	}

	if product.Name == "" {
		HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	newProduct, err := e.productApp.Create(&product)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	HandleSucces(c, newProduct)
}

func (e *ProductController) viewProducts(c *gin.Context) {
	products, err := e.productApp.ViewAll()
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*products) == 0 {
		HandleError(c, http.StatusNotFound, "list product is empty")
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")

	// HandleSucces(c, products)
	c.IndentedJSON(http.StatusOK, products)
}

func (e *ProductController) viewProductById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	product, err := e.productApp.ViewById(id)
	if err != nil {
		HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	HandleSucces(c, product)
}

func (e *ProductController) editProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	_, err = e.productApp.ViewById(id)
	if err != nil {
		HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	var tempProduct = domain.Product{}
	err = c.Bind(&tempProduct)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if tempProduct.ID != 0 {
		HandleError(c, http.StatusBadRequest, "input not permitted")
		return
	}
	if tempProduct.Name == "" {
		HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	product, err := e.productApp.Edit(id, &tempProduct)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	HandleSucces(c, product)
}

func (e *ProductController) deleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	err = e.productApp.Delete(id)
	if err != nil {
		HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	HandleSucces(c, "success delete data")
}

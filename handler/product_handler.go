package handler

import (
	"net/http"
	"strconv"

	"order-management/model"
	"order-management/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	productNumber, err := strconv.Atoi(c.Param("productNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product number"})
		return
	}
	productID := c.Param("productID")

	product, err := h.service.GetByID(productNumber, productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
	productNumber, err := strconv.Atoi(c.Param("productNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product number"})
		return
	}
	productID := c.Param("productID")

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ProductNumber = productNumber
	product.ProductID = productID

	if err := h.service.Update(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	productNumber, err := strconv.Atoi(c.Param("productNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product number"})
		return
	}
	productID := c.Param("productID")

	if err := h.service.Delete(productNumber, productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

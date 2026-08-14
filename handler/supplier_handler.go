package handler

import (
	"net/http"
	"strconv"

	"order-management/model"
	"order-management/service"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	service *service.SupplierService
}

func NewSupplierHandler(service *service.SupplierService) *SupplierHandler {
	return &SupplierHandler{service: service}
}

func (h *SupplierHandler) Create(c *gin.Context) {
	var supplier model.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, supplier)
}

func (h *SupplierHandler) GetAll(c *gin.Context) {
	suppliers, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}

func (h *SupplierHandler) GetByID(c *gin.Context) {
	supplierNumber, err := strconv.Atoi(c.Param("supplierNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplier number"})
		return
	}
	supplierID := c.Param("supplierID")

	supplier, err := h.service.GetByID(supplierNumber, supplierID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "supplier not found"})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

func (h *SupplierHandler) Update(c *gin.Context) {
	supplierNumber, err := strconv.Atoi(c.Param("supplierNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplier number"})
		return
	}
	supplierID := c.Param("supplierID")

	var supplier model.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	supplier.SupplierNumber = supplierNumber
	supplier.SupplierID = supplierID

	if err := h.service.Update(&supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

func (h *SupplierHandler) Delete(c *gin.Context) {
	supplierNumber, err := strconv.Atoi(c.Param("supplierNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplier number"})
		return
	}
	supplierID := c.Param("supplierID")

	if err := h.service.Delete(supplierNumber, supplierID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

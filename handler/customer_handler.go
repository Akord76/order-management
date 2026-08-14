package handler

import (
	"net/http"
	"strconv"

	"order-management/model"
	"order-management/service"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(service *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, customer)
}

func (h *CustomerHandler) GetAll(c *gin.Context) {
	customers, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) GetByID(c *gin.Context) {
	custNumber, err := strconv.Atoi(c.Param("custNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer number"})
		return
	}
	custID := c.Param("custID")

	customer, err := h.service.GetByID(custNumber, custID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) Update(c *gin.Context) {
	custNumber, err := strconv.Atoi(c.Param("custNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer number"})
		return
	}
	custID := c.Param("custID")

	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer.CustNumber = custNumber
	customer.CustID = custID

	if err := h.service.Update(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) Delete(c *gin.Context) {
	custNumber, err := strconv.Atoi(c.Param("custNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer number"})
		return
	}
	custID := c.Param("custID")

	if err := h.service.Delete(custNumber, custID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

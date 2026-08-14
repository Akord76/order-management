package handler

import (
	"net/http"
	"strconv"

	"order-management/model"
	"order-management/service"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, employee)
}

func (h *EmployeeHandler) GetAll(c *gin.Context) {
	employees, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	employeeID, err := strconv.Atoi(c.Param("employeeID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return
	}
	cardNumber, err := strconv.Atoi(c.Param("cardNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid card number"})
		return
	}

	employee, err := h.service.GetByID(employeeID, cardNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	employeeID, err := strconv.Atoi(c.Param("employeeID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return
	}
	cardNumber, err := strconv.Atoi(c.Param("cardNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid card number"})
		return
	}

	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee.EmployeeID = employeeID
	employee.EmployeeCardNumber = cardNumber

	if err := h.service.Update(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
	employeeID, err := strconv.Atoi(c.Param("employeeID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return
	}
	cardNumber, err := strconv.Atoi(c.Param("cardNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid card number"})
		return
	}

	if err := h.service.Delete(employeeID, cardNumber); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

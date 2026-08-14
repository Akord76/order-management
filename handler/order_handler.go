package handler

import (
	"net/http"
	"strconv"

	"order-management/model"
	"order-management/service"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var order model.OrderMaster
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetAll(c *gin.Context) {
	orders, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}
	orderNo := c.Param("orderNo")

	order, err := h.service.GetByID(orderID, orderNo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Update(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}
	orderNo := c.Param("orderNo")

	var order model.OrderMaster
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.OrderID = orderID
	order.OrderNo = orderNo

	if err := h.service.Update(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Delete(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}
	orderNo := c.Param("orderNo")

	if err := h.service.Delete(orderID, orderNo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// --- OrderDetail endpoints ---

func (h *OrderHandler) AddDetail(c *gin.Context) {
	var detail model.OrderDetail
	if err := c.ShouldBindJSON(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	detail.OrderNo = c.Param("orderNo")

	if err := h.service.AddDetail(&detail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, detail)
}

func (h *OrderHandler) GetDetails(c *gin.Context) {
	orderNo := c.Param("orderNo")
	details, err := h.service.GetDetails(orderNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}

func (h *OrderHandler) DeleteDetail(c *gin.Context) {
	orderNo := c.Param("orderNo")
	orderDetailNo, err := strconv.Atoi(c.Param("orderDetailNo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order detail no"})
		return
	}

	if err := h.service.DeleteDetail(orderDetailNo, orderNo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

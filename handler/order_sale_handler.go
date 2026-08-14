package handler

import (
	"net/http"
	"strconv"

	"order-management/model"
	"order-management/service"

	"github.com/gin-gonic/gin"
)

type OrderSaleHandler struct {
	service *service.OrderSaleService
}

func NewOrderSaleHandler(service *service.OrderSaleService) *OrderSaleHandler {
	return &OrderSaleHandler{service: service}
}

func (h *OrderSaleHandler) Create(c *gin.Context) {
	var orderSale model.OrderSaleMaster
	if err := c.ShouldBindJSON(&orderSale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&orderSale); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, orderSale)
}

func (h *OrderSaleHandler) GetAll(c *gin.Context) {
	orderSales, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderSales)
}

func (h *OrderSaleHandler) GetByID(c *gin.Context) {
	orderSaleNo, err := strconv.Atoi(c.Param("orderSaleNo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order sale no"})
		return
	}
	orderSaleID := c.Param("orderSaleID")

	orderSale, err := h.service.GetByID(orderSaleNo, orderSaleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order sale not found"})
		return
	}
	c.JSON(http.StatusOK, orderSale)
}

func (h *OrderSaleHandler) Update(c *gin.Context) {
	orderSaleNo, err := strconv.Atoi(c.Param("orderSaleNo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order sale no"})
		return
	}
	orderSaleID := c.Param("orderSaleID")

	var orderSale model.OrderSaleMaster
	if err := c.ShouldBindJSON(&orderSale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderSale.OrderSaleNo = orderSaleNo
	orderSale.OrderSaleID = orderSaleID

	if err := h.service.Update(&orderSale); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderSale)
}

func (h *OrderSaleHandler) Delete(c *gin.Context) {
	orderSaleNo, err := strconv.Atoi(c.Param("orderSaleNo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order sale no"})
		return
	}
	orderSaleID := c.Param("orderSaleID")

	if err := h.service.Delete(orderSaleNo, orderSaleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// --- OrderSaleDetail endpoints ---

func (h *OrderSaleHandler) AddDetail(c *gin.Context) {
	orderSaleNo, err := strconv.Atoi(c.Param("orderSaleNo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order sale no"})
		return
	}

	var detail model.OrderSaleDetail
	if err := c.ShouldBindJSON(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	detail.OrderSaleNo = orderSaleNo

	if err := h.service.AddDetail(&detail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, detail)
}

func (h *OrderSaleHandler) GetDetails(c *gin.Context) {
	orderSaleNo, err := strconv.Atoi(c.Param("orderSaleNo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order sale no"})
		return
	}

	details, err := h.service.GetDetails(orderSaleNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}

func (h *OrderSaleHandler) DeleteDetail(c *gin.Context) {
	orderSaleDetailNo, err := strconv.Atoi(c.Param("orderSaleDetailNo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order sale detail no"})
		return
	}
	orderSaleDetailID := c.Param("orderSaleDetailID")

	if err := h.service.DeleteDetail(orderSaleDetailNo, orderSaleDetailID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

package handler

import (
	"net/http"

	"order-management/model"
	"order-management/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &model.AppUser{
		Username: req.Username,
		FullName: req.FullName,
		Email:    req.Email,
		RoleName: req.RoleName,
		IsActive: true,
	}

	userID, err := h.service.Register(user, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": userID, "username": user.Username})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"user_id":   user.UserID,
			"username":  user.Username,
			"full_name": user.FullName,
			"role_name": user.RoleName,
		},
	})
}

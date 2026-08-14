package handler

import (
	"net/http"
	"strconv"

	"order-management/model"
	"order-management/service"

	"github.com/gin-gonic/gin"
)

// UserHandler exposes the ADMIN "User Management" endpoints:
// Create User, Update User, Delete User, View Users.
type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
}

// Create -> ADMIN: Create User
func (h *UserHandler) Create(c *gin.Context) {
	var req createUserRequest
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

	userID, err := h.service.Create(user, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": userID, "username": user.Username, "role_name": user.RoleName})
}

// GetAll -> ADMIN: View Users
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetByID -> ADMIN: View Users (single)
func (h *UserHandler) GetByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.service.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

type updateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
	IsActive bool   `json:"is_active"`
}

// Update -> ADMIN: Update User
func (h *UserHandler) Update(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &model.AppUser{
		UserID:   userID,
		FullName: req.FullName,
		Email:    req.Email,
		RoleName: req.RoleName,
		IsActive: req.IsActive,
	}

	if err := h.service.Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Delete -> ADMIN: Delete User
func (h *UserHandler) Delete(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := h.service.Delete(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

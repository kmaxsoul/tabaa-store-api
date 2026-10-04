package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmaxsoul/tabaa-store-api/models"
	"github.com/kmaxsoul/tabaa-store-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo *repository.UserRepository
}

// 1. إنشاء هيكل مخصص لاستقبال بيانات التسجيل فقط
type RegisterRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password_hash" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number"`
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid information, please check what you entered.",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process data due to an internal error",
		})
		return
	}

	user := models.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		PhoneNumber:  req.PhoneNumber,
	}

	err = h.repo.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Sign-up failed. This email address is already in use",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account created successfully",
		"user":    user,
	})
}

package controllers

import (
	"go-shop/models"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Login      string `json:"login" binding:"required"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
	Phone      string `json:"phone"`
	Email      string `json:"email" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Login:      input.Login,
		Password:   string(hashedPassword),
		FirstName:  input.FirstName,
		SecondName: input.SecondName,
		Phone:      input.Phone,
		Email:      input.Email,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}


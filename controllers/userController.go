package controllers

import (
	"log"
	"net/http"
	"rms-platform/models"
	"rms-platform/utils"

	"rms-platform/database"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {

		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, please check the provided data."})
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username, email, and password are required."})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {

		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password, please try again."})
		return
	}

	user.Password = hashedPassword

	var existingUser models.User
	if err := database.DB.Where("email = ? OR username = ?", user.Email, user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email or username already exists."})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {

		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user, please try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"fullName": user.FullName,
			"phone":    user.Phone,
			"role":     user.Role,
		},
	})
}

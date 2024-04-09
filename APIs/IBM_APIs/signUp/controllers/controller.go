package controllers

import (
	"fmt"
	"net/http"
	"os"
	"signUp/initializers"
	"signUp/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var db = initializers.DB
var user models.User

func GetUsersv1(c *gin.Context) {

	if err := db.Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {

	var users []models.User

	db = initializers.DB

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

func GetUserByUsernameAndPassword(c *gin.Context) {
	email := c.Param("email")
	password := c.Param("password")
	message := email + " is " + password
	c.String(http.StatusOK, message)

	db = initializers.DB

	if err := db.Where("email = ? AND password >= ?", email, password).Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

func Login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to red body",
		})
		return
	}

	db = initializers.DB

	var user models.User
	db.First(&user, "email = ? AND password >= ?", body.Email, body.Password)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
	}

	// creating a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
		"username": user.Username,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	fmt.Println(tokenString, err)

	if err != nil {
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

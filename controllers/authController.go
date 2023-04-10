package controllers

import (
	"os"
	"time"

	"github.com/baguspanji/go-crud/initializers"
	"github.com/baguspanji/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func AuthUser(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(200, gin.H{
		"message": "success",
		"data":    user,
	})
}

func AuthLogin(c *gin.Context) {
	var body struct {
		Username *string `json:"username" binding:"required"`
		Password string  `json:"password" binding:"required"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"errors":  err.Error(),
		})
		return
	}

	var user models.User
	result := initializers.DB.Where("username = ?", body.Username).First(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if errCompare != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, errToken := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if errToken != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	c.JSON(200, gin.H{
		"message": "success login",
		"token":   tokenString,
		"data":    userResponse,
	})

}

func AuthRegister(c *gin.Context) {
	body := models.UserRequest{}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"errors":  err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "error",
			"error":   result.Error,
		})
		return
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	c.JSON(201, gin.H{
		"message": "data created",
		"data":    userResponse,
	})
}

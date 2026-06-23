package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	token, err := Register(req)

	if err != nil {

		switch err.Error() {

		case "email already exists":
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": "Email already registered",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User registered successfully",
		"token":   token,
	})
}

func LoginHandler(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})

		return
	}

	user, token, err := Login(req)

	if err != nil {

		switch err.Error() {

		case "invalid email or password":

			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid email or password",
			})

		default:

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"bio":   user.Bio,
			"theme": user.Theme,
		},
	})
}

func GetProfileHandler(c *gin.Context) {

	userID := c.MustGet("userID").(string)

	user, err := GetProfile(userID)

	if err != nil {

		switch err.Error() {

		case "user not found":

			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
			})

		default:

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"bio":        user.Bio,
			"theme":      user.Theme,
			"created_at": user.CreatedAt,
		},
	})
}

package auth

import (
	"go-auth/internal/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet("validatedRequest").(RegisterRequest)

		// JWT TODO

		user := models.User{
			Email:    req.Email,
			Password: req.Password,
			Name:     req.Name,
			Role:     req.Role,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

package api

import (
	"go-auth/api/user"
	database "go-auth/internal/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(db *database.Database) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	user.RegisterRoutes(router, db)

	return router
}

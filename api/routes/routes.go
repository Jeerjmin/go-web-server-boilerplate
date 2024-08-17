package routes

import (
	"go-auth/api/handlers"
	"go-auth/api/validator"
	database "go-auth/internal/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(db *database.Database) *gin.Engine {
	router := gin.Default()

	validator.InitValidator()

	// Регистрация маршрутов для пользователя
	router.POST("/users", handlers.CreateUser(db.DB))
	router.GET("/users/:id", handlers.GetUser(db.DB))
	router.GET("/users", handlers.GetUsers(db.DB))
	router.PUT("/users/:id", handlers.UpdateUser(db.DB))
	router.DELETE("/users/:id", handlers.DeleteUser(db.DB))

	return router
}

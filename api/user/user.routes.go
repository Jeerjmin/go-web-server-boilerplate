package user

import (
	"go-auth/api/auth"
	"go-auth/api/validator"
	database "go-auth/internal/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *database.Database) *gin.Engine {
	v := validator.NewValidator()

	router.POST("/users", auth.ProtectedRoute(), validator.ValidateRequest[CreateUserRequest](v), CreateUser(db.DB))
	router.GET("/users/:id", auth.ProtectedRoute(), v.ValidateIDParam(), GetUser(db.DB))
	router.GET("/users", auth.ProtectedRoute(), GetUsers(db.DB))
	router.PUT("/users/:id", auth.ProtectedRoute(), v.ValidateIDParam(), validator.ValidateRequest[UpdateUserRequest](v), UpdateUser(db.DB))
	router.DELETE("/users/:id", auth.ProtectedRoute(), v.ValidateIDParam(), DeleteUser(db.DB))

	return router
}

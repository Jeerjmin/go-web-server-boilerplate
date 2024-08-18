package auth

import (
	"go-auth/api/validator"
	database "go-auth/internal/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *database.Database) *gin.Engine {
	v := validator.NewValidator()

	router.POST("/auth/login", validator.ValidateRequest[LoginRequest](v), Login(db.DB))
	router.POST("/auth/register", validator.ValidateRequest[RegisterRequest](v), Register(db.DB))

	return router

}

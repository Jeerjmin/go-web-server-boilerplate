package image

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) *gin.Engine {

	router.POST("/upload", func(c *gin.Context) {
		UploadHandlerSync(c)
	})

	return router

}

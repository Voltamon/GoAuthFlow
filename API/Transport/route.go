package transport

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *Handler) {
	router.Use(Logger(), ErrorHandler())

	v1 := router.Group("/v1")
	{
		v1.POST("/register", h.Register)
		v1.POST("/login", h.Login)
	}
}

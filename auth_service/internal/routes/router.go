package routes

import (
	"auth_service/internal/handler"
	"auth_service/internal/middleware"
	"auth_service/util"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, handler *handler.UserHandler) {
	userGroup := r.Group("/")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
		userGroup.Use(middleware.AuthMiddleware([]byte(util.SecretKey)))
		userGroup.GET("/me", handler.CurrentUser)
		userGroup.POST("/logout", handler.Logout)
	}
}

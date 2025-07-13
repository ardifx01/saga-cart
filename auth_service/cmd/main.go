package main

import (
	"auth_service/internal/db"
	"auth_service/internal/handler"
	"auth_service/internal/repository"
	"auth_service/internal/routes"
	"auth_service/internal/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// CORS

	// db (postgres)
	db, _ := db.ConnectPostgre()
	// db (postgres)

	// handler
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	// handler

	// route
	routes.Router(r, userHandler)
	// route

	log.Println("[AuthService] auth service run at port :8084")
	r.Run(":8084")
}

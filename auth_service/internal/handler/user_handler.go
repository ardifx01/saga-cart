package handler

import (
	"auth_service/internal/contracts"
	"auth_service/internal/domain"
	"auth_service/util"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService contracts.UserServiceContract
}

func NewUserHandler(service contracts.UserServiceContract) *UserHandler {
	return &UserHandler{userService: service}
}

func (u *UserHandler) CurrentUser(c *gin.Context) {
	id, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var userId int
	switch v := id.(type) {
	case float64:
		userId = int(v)
	case int:
		userId = v
	case string:
		parsed, err := strconv.Atoi(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
			return
		}
		userId = parsed
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unexpected user id type"})
		return
	}

	user, err := u.userService.FindById(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while get current user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) Register(c *gin.Context) {
	var req domain.UserCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	createdUser, err := u.userService.Create(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error create user"})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

func (u *UserHandler) Login(c *gin.Context) {
	var req domain.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	user, err := u.userService.FindByUsername(c, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error get user"})
		return
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "password tidak sama"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(util.SecretKey))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}

func (u *UserHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, "Logout route")
}

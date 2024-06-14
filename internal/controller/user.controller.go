package controller

import (
	"GO-ECOMMERCE-BACKEND-API/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

// Đây là một hàm khởi tạo (constructor) cho UserController.
// Hàm này chịu trách nhiệm tạo và trả về một con trỏ tới một instance của UserController.
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}
func (uc *UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  uc.userService.GetInfoUserService(),
		"users": []string{"Thanh", "m10", "Zidance"},
	})
}

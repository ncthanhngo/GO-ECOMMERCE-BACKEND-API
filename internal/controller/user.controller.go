package controller

import (
	"GO-ECOMMERCE-BACKEND-API/internal/service"
	"GO-ECOMMERCE-BACKEND-API/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{UserService: service}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.UserService.Register(request.Email, request.Password); err != nil {
		response.FailResponse(c, 20004, "invalid param")
		return
	}

	//c.JSON(http.StatusCreated, gin.H{"message": "user created"})
	response.SucessResponse(c, 20001, "User created", request.Email)
}
func (ctrl *UserController) CreateUser(c *gin.Context) {
	// Thực hiện logic tạo người dùng
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	// Thực hiện logic cập nhật người dùng
}

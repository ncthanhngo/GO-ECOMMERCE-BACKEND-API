package routers

import (
	"GO-ECOMMERCE-BACKEND-API/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping:id", controller.NewUserController().GetUserById)
		v1.POST("/ping")
		v1.PATCH("/users/:id")
	}
	return r
}

package routers

import (
	"GO-ECOMMERCE-BACKEND-API/internal/controller"
	_ "GO-ECOMMERCE-BACKEND-API/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controller.UserController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", userController.RegisterUser)
		// Các route khác có thể được thêm ở đây
		v1.POST("/users", userController.CreateUser)      // Giả sử bạn có hàm CreateUser
		v1.PATCH("/users/:id", userController.UpdateUser) // Giả sử bạn có hàm UpdateUser
	}
	return r
}

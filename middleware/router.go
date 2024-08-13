package middleware

import (
	"lookinlabs-test/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController controller.UserController) *gin.Engine {
	router := gin.Default()

	router.GET("api/v1/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("api/v1/users", userController.CreateUser)
	router.GET("api/v1/users", userController.GetUsers)
	router.GET("api/v1/users/:id", userController.GetUserByID)
	router.PATCH("api/v1/users/:id", userController.UpdateUser)

	return router
}

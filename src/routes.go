package src

import (
	"github.com/gin-gonic/gin"

	controller "mn128.com/src/controller"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/all-users", controller.GetAllUsers)
	router.GET("/single/:id", controller.GetSingleUser)
	router.POST("/signup", controller.Singup)
	router.POST("/login", controller.Login)
	router.POST("/create", controller.CreateUser)
	router.PATCH("/update/:id", controller.UpdateUser)
	router.PUT("/forget-password/:id", controller.ForgotPassword)
	router.DELETE("/delete/:id", controller.DeleteUser)
}

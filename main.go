package main

import (
	"github.com/gin-gonic/gin"

	routes "mn128.com/src"
)

func main() {
	router := gin.Default()
	routes.UserRoutes(router)

	router.Run(":8080")
}

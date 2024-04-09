package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "mn128.com/src/model"
)

func GetAllUsers(c *gin.Context) {
	var users = []model.User{
		{
			ID:             "1",
			Name:           "Mahadi",
			Email:          "pGwXO@example.com",
			Password:       "12345",
			Age:            23,
			Img:            "jsdnjwdiwk",
			DrivingLicense: "xyz123",
			Role:           "driver",
			Location:       "airport",
			History:        []string{"1", "2", "3"},
		},
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetSingleUser(c *gin.Context) {

}

func Singup(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
}

func UpdateUser(c *gin.Context) {
}

func DeleteUser(c *gin.Context) {
}

func ForgotPassword(c *gin.Context) {
}

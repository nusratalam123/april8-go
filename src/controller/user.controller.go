package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	client "mn128.com/src/db"
	"mn128.com/src/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/go-playground/validator/v10"
)

// var validate *validator.Validate
// func init() {
// 	validate = validator.New()
// }

// func (u *types.User) Validate(){
// 	return validate.Struct(u)
// }

func GetAllUsers(c *gin.Context) {
	coll := client.GetDbInstance().Instance
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		println(err.Error())
		panic(err)
	}

	var results []types.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}

	c.IndentedJSON(http.StatusOK, results)
}

func GetSingleUser(c *gin.Context) {

	id := c.Param("id")
	coll := client.GetDbInstance().Instance
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var result types.User
	err = coll.FindOne(context.TODO(), bson.D{{Key: "_id",Value: objectId}}).Decode(&result)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, result)
	
}

func Singup(c *gin.Context) {
	var user types.User
	
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	coll := client.GetDbInstance().Instance
	docs := []interface{}{user}

	_, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})


}

func Login(c *gin.Context) {

	var user types.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	coll := client.GetDbInstance().Instance
	var result types.User
	err := coll.FindOne(context.TODO(), bson.D{{Key: "email",Value: user.Email},{Key: "password",Value: user.Password}}).Decode(&result)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, result)

}

func CreateUser(c *gin.Context) {
	var user types.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	coll := client.GetDbInstance().Instance
	docs := []interface{}{user}

	_, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")
	var user types.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	coll := client.GetDbInstance().Instance
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := coll.UpdateOne(context.TODO(), bson.D{{Key: "_id",Value: objectId}}, bson.D{{Key: "$set",Value: user}})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	coll := client.GetDbInstance().Instance
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := coll.DeleteOne(context.TODO(), bson.D{{Key: "_id",Value: objectId}})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK,result )
}

func ForgotPassword(c *gin.Context) {
}

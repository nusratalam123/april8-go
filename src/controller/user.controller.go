package controller

import (
	"github.com/gin-gonic/gin"
)

// type Tea struct {
// 	Type   string
// 	Rating int32
// 	Vendor []string `bson:"vendor,omitempty" json:"vendor,omitempty"`
// }

func GetAllUsers(c *gin.Context) {
	//FIXME: use this for all db instance
	// coll := client.GetDbInstance().Instance

	// filter := bson.D{{"type", "Masala"}}
	// cursor, err := coll.Find(context.TODO(), filter)
	// if err != nil {
	// 	println(err.Error())
	// 	panic(err)
	// }
	// var results []Tea
	// if err = cursor.All(context.TODO(), &results); err != nil {
	// 	panic(err)
	// }
	// for _, result := range results {
	// 	res, _ := bson.MarshalExtJSON(result, false, false)
	// 	fmt.Println(string(res))
	// }

	// c.IndentedJSON(http.StatusOK, results)
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

package types

type User struct {
	//ID             string   `bson:"id,omitempty" json:"id,omitempty"`
	Name           string   `bson:"name,omitempty" json:"name,omitempty" validate:"required,min=3,max=40"`
	Email          string   `bson:"email,omitempty" json:"email,omitempty" validate:"required,email"`
	Password       string   `bson:"password,omitempty" json:"password,omitempty" validate:"required,min=6,max=20"`
	Age            int      `bson:"age,omitempty" json:"age,omitempty"`
	Img            string   `bson:"img,omitempty" json:"img,omitempty"`
	DrivingLicense string   `bson:"drivingLicense,omitempty" json:"drivingLicense,omitempty"`
	Role           string   `bson:"role,omitempty" json:"role,omitempty"`
	Location       string   `bson:"location,omitempty" json:"location,omitempty" validate:"required"`
	History        []string `bson:"history,omitempty" json:"history,omitempty"`
}


// package types

// import (
// 	"github.com/go-playground/validator/v10"
// )

// type User struct {
// 	ID       string `json:"id" validate:"required"`
// 	Name     string `json:"name" validate:"required,min=2,max=100"`
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=6,max=100"`
// }

// var validate *validator.Validate

// func init() {
// 	validate = validator.New()
// }

// // Validate validates the User struct
// func (u *User) Validate() error {
// 	return validate.Struct(u)
// }
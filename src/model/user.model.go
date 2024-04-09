package model

type User struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Password       string   `json:"password"`
	Age            int      `json:"age"`
	Img            string   `json:"img"`
	DrivingLicense string   `json:"drivingLicense"`
	Role           string   `json:"role"`
	Location       string   `json:"location"`
	History        []string `json:"history"`
}

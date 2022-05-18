package dto

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Add struct {
	HouseNo  string `json:"house_no"`
	Floor    int    `json:"floor"`
	Street   string `json:"street"`
	Landmark string `json:"landmark"`
	Area     string `json:"area"`
	City     string `json:"city"`
	District string `json:"district"`
	Pincode  int    `json:"pincode"`
	State    string `json:"state"`
}
type Emp_details struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmpId        string    `json:"emp_id"`
	Role         string    `json:"role"`
	Department   string    `json:"department"`
	JoiningDate  time.Time `json:"joining_date"`
	ReportedTo   string    `json:"reported_to"`
	MobileNo     int       `json:"mobile_no"`
	EmailAddress string    `json:"email_address"`
	Address      Add       `json:"address"`
	Password     string    `json:"password"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

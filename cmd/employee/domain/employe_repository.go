package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ehs/cmd/employee/dto"
	h "github.com/ehs/pkg/dto"
	er "github.com/ehs/pkg/err"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

type HttpRepository struct {
	client *http.Client
}

func (e HttpRepository) EmployeLogin(credentials dto.Credentials) *dto.LoginResponse {
	expirationTime := time.Now().Add(time.Minute * 1)
	claims := &dto.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Println("helloooo")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println("helloooo2345")
	}

	return &dto.LoginResponse{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}
}

func (e HttpRepository) EmployeeRegistration(emp dto.Emp_details) (*dto.Emp_details, *er.AppError) {

	uu, _ := json.Marshal(emp)
	details := h.HttpRequestDetails{
		QueryParams: nil,
		Headers:     nil,
		Body:        uu,
	}
	url := "http://localhost:8008/createuser"
	data, err := h.NewHttpRequest(e.client, http.MethodPost, url, details)

	if err != nil {
		return nil, er.NewBadRequest("")
	}
	var usr *dto.Emp_details
	json.Unmarshal(data, &usr)
	fmt.Println(usr)
	return usr, nil

}

func (e HttpRepository) RefreshToken() (*dto.LoginResponse, *er.AppError) {
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &dto.Claims{}
	claims.ExpiresAt = expirationTime.Unix()
	fmt.Println("helloooo", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println("helloooo2345")
	}
	a := &http.Cookie{
		Name:    "refreshtoken",
		Value:   tokenString,
		Expires: expirationTime,
	}

	return &dto.LoginResponse{
		Name:    "refresh",
		Value:   tokenString,
		Expires: expirationTime,
		Cookie:  *a,
	}, nil

}

func (e HttpRepository) EmpDetails(username string) (*dto.Emp_details, *er.AppError) {

	details := h.HttpRequestDetails{
		QueryParams: nil,
		Headers:     nil,
		Body:        nil,
	}
	url := "http://localhost:8008/getuser/" + "empid" + "/" + username
	data, err := h.NewHttpRequest(e.client, http.MethodGet, url, details)

	if err != nil {
		return nil, er.NewBadRequest("")
	}

	var usr *dto.Emp_details
	json.Unmarshal(data, &usr)

	return usr, nil
}
func NewHttpRepository(client *http.Client) HttpRepository {
	return HttpRepository{client}
}

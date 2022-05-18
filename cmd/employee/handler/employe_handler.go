package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ehs/cmd/employee/dto"
	"github.com/ehs/cmd/employee/service"
)

type Handlers struct {
	Service service.Service
}

var jwtKey = []byte("secret_key")

func Token(w http.ResponseWriter, r *http.Request) *jwt.Token {
	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	tokenstr := cookie.Value
	fmt.Println("in refreshtoken", tokenstr)
	claims := &dto.Claims{}
	fmt.Println("claims in refreshtoken", claims)
	tkn, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	fmt.Println("tkn in refresh", tkn)
	if !tkn.Valid {

		writeResponse(w, http.StatusUnauthorized, "")
	}
	return tkn

}

func (d *Handlers) RefreshToken(w http.ResponseWriter, r *http.Request) {

	Token(w, r)
	if time.Unix(60, 0).Sub(time.Now()) <= 30*time.Second {

		data, er := d.Service.RefreshToken()
		http.SetCookie(w,
			&data.Cookie)
		if er != nil {
			writeResponse(w, http.StatusBadGateway, er)
		} else {
			writeResponse(w, http.StatusOK, data)
		}

	}

}
func (d *Handlers) EmployeeRegistration(w http.ResponseWriter, r *http.Request) {
	var emp dto.Emp_details
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "error while decoding request body at handler layer: "+err.Error())
		return
	}

	emp_det, apperr := d.Service.EmployeeRegistration(emp)
	if apperr != nil {
		writeResponse(w, apperr.Messages.ResponseCode, apperr.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, emp_det)
	}
}
func (d *Handlers) EmployeLogin(w http.ResponseWriter, r *http.Request) {

	var credentials dto.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	details, er := d.Service.EmployeLogin(credentials)
	if er != nil {
		writeResponse(w, http.StatusBadRequest, er)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   details.Value,
			Expires: details.Expires,
		})

	writeResponse(w, http.StatusOK, details)
}

func (d *Handlers) Home(w http.ResponseWriter, r *http.Request) {

}
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		//panic(err)
	}
}

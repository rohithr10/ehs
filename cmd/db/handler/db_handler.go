package handler

import (
	"encoding/json"
	"net/http"

	// "practise/User/dto"

	"github.com/ehs/cmd/db/service"
	"github.com/ehs/cmd/employee/dto"
	"github.com/gorilla/mux"
)

type Dbhandlers struct {
	Service service.DbUserService
}

func (d Dbhandlers) CreateUser(w http.ResponseWriter, r *http.Request) {

	var emp dto.Emp_details
	_ = json.NewDecoder(r.Body).Decode(&emp)
	user, err := d.Service.CreateUser(emp)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, nil)
	} else {
		writeResponse(w, http.StatusOK, user)
	}

}

func (d Dbhandlers) GetUserByFilter(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	field := params["field"]
	username := params["id"]
	
	user, err := d.Service.GetUserByFilter(field, username)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, nil)
	} else {
		writeResponse(w, http.StatusOK, user)
	}

}
func (d Dbhandlers) CreateRoom(w http.ResponseWriter, r *http.Request){

}
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

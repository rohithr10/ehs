package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ehs/cmd/employee/domain"
	"github.com/ehs/cmd/employee/service"
	"github.com/gorilla/mux"
)

func Start() {

	r := mux.NewRouter()
	client := &http.Client{}

	//wiring
	var d = Handlers{
		Service: service.NewService(domain.NewHttpRepository(client)),
	} // define routes

	fmt.Println(d)
	fmt.Println("Started Service")
	r.HandleFunc("/create", d.RefreshToken).Methods(http.MethodPost)
	r.HandleFunc("/token/{email}", d.Home).Methods(http.MethodGet)
	r.HandleFunc("/login", d.EmployeLogin).Methods(http.MethodPost)
	r.HandleFunc("/register", d.EmployeeRegistration).Methods(http.MethodPost)
	r.HandleFunc("/refresh", d.RefreshToken).Methods(http.MethodPost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8010"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))

}

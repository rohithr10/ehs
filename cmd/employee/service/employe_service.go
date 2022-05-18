package service

import (
	"github.com/ehs/cmd/employee/domain"
	"github.com/ehs/cmd/employee/dto"
	d "github.com/ehs/pkg/dto"

	"fmt"

	t "github.com/ehs/pkg/err"
)

type Service interface {
	EmployeLogin(credentials dto.Credentials) (*dto.LoginResponse, *t.AppError)
	EmployeeRegistration(emp dto.Emp_details) (*d.Response, *t.AppError)
	RefreshToken() (*dto.LoginResponse, *t.AppError)
}

type DefaultService struct {
	Repo domain.Repository
}

func (e DefaultService) EmployeLogin(credentials dto.Credentials) (*dto.LoginResponse, *t.AppError) {

	// call getempdetails----save in a variable based on username
	// check the password
	fmt.Println("in service layer;;;;;;;;")
	data, er := e.Repo.EmpDetails(credentials.Username)
	if er != nil {
		return nil, t.NewBadRequest("error occurred while creating document from db ")

	}
	fmt.Println("data fron db", data)
	//pass to handlers
	if credentials.Username == data.EmpId {
		if credentials.Password == data.Password {
			// call repo
			det := e.Repo.EmployeLogin(credentials)
			return &dto.LoginResponse{
				Name:    "token",
				Value:   det.Value,
				Expires: det.Expires,
			}, nil

		} else {
			fmt.Println("qwewewee")
			return nil, t.NewUnexpectedError("password is incorrect")

		}
	}
	return nil, t.NewUnAuthorizedError("password is incorrect")

}
func (e DefaultService) EmployeeRegistration(emp dto.Emp_details) (*d.Response, *t.AppError) {
	fmt.Println("1111111111111111111111111")
	_, err := e.Repo.EmployeeRegistration(emp)
	if err != nil {
		return nil, err.AsMessage()
	}
	a := d.Response{
		Message:      "Successfully Registered",
		ResponseCode: 200,
	}
	return &a, nil

}

func (e DefaultService) RefreshToken() (*dto.LoginResponse, *t.AppError) {
	return e.Repo.RefreshToken()
}

func NewService(Repo domain.Repository) DefaultService {

	return DefaultService{Repo}
}

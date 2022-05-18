package domain

import (
	"github.com/ehs/cmd/employee/dto"
	e "github.com/ehs/pkg/err"
)

type Repository interface {
	EmployeLogin(credentials dto.Credentials) *dto.LoginResponse
	EmployeeRegistration(emp dto.Emp_details) (*dto.Emp_details, *e.AppError)
	EmpDetails(username string) (*dto.Emp_details, *e.AppError)
	RefreshToken() (*dto.LoginResponse, *e.AppError)
}

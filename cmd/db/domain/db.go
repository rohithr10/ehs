package domain

import (
	"github.com/ehs/cmd/employee/dto"
	"github.com/ehs/pkg/err"
)

type DbUserRepository interface {
	CreateUser(emp dto.Emp_details) (*dto.Emp_details, *err.AppError)
	GetUserByFilter(field, username string) (*dto.Emp_details, *err.AppError)
	CreateRoom() *err.AppError
}

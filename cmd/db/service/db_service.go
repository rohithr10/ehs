package service

import (
	// "practise/User/dto"
	//"context"

	"fmt"

	"github.com/ehs/cmd/db/domain"
	"github.com/ehs/cmd/employee/dto"
	"github.com/ehs/pkg/err"
)

type DbUserService interface {
	CreateUser(emp dto.Emp_details) (*dto.Emp_details, *err.AppError)
	GetUserByFilter(field, username string) (*dto.Emp_details, *err.AppError)
	CreateRoom() *err.AppError
}

type DefaultDbService struct {
	httpRepo domain.DbUserRepository
}

func (s DefaultDbService) CreateUser(emp dto.Emp_details) (*dto.Emp_details, *err.AppError) {
	fmt.Println("4444444444444444444444")
	return s.httpRepo.CreateUser(emp)

}
func (s DefaultDbService) GetUserByFilter(field, username string) (*dto.Emp_details, *err.AppError) {
	fmt.Println("4444444444444444444444qweerr")

	return s.httpRepo.GetUserByFilter(field, username)

}
func (s DefaultDbService) CreateRoom() *err.AppError {
	return nil
}
func Newdbuserservice(Repo domain.DbUserRepository) DefaultDbService {
	return DefaultDbService{Repo}
}

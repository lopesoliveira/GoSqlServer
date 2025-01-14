package service

import (
	"github.com/lopesoliveira/GoSqlServer/data/request"
	"github.com/lopesoliveira/GoSqlServer/data/response"
)

type GormSqlServerService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []response.UserResponse
}

package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/lopesoliveira/GoSqlServer/data/request"
	"github.com/lopesoliveira/GoSqlServer/data/response"
	"github.com/lopesoliveira/GoSqlServer/helper"
	"github.com/lopesoliveira/GoSqlServer/model"
	"github.com/lopesoliveira/GoSqlServer/repository"
)

type GormSqlServerServiceImpl struct {
	GormSqlServerRepository repository.GormSqlServerRepository
	Validate                *validator.Validate
}

func NewGormSqlServerServiceImpl(gormSqlServerRepository repository.GormSqlServerRepository, validate *validator.Validate) GormSqlServerService {
	return &GormSqlServerServiceImpl{
		GormSqlServerRepository: gormSqlServerRepository,
		Validate:                validate,
	}
}

func (u *GormSqlServerServiceImpl) Create(user request.CreateUserRequest) {
	err := u.Validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
	}
	u.GormSqlServerRepository.Save(userModel)
}

func (u *GormSqlServerServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.GormSqlServerRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	userData.Name = user.Name
	userData.Email = user.Email
	userData.Password = user.Password
	userData.Age = user.Age
	u.GormSqlServerRepository.Update(userData)
}

func (u *GormSqlServerServiceImpl) Delete(userId int) {
	u.GormSqlServerRepository.Delete(userId)
}

func (u *GormSqlServerServiceImpl) FindById(userId int) response.UserResponse {
	userData, err := u.GormSqlServerRepository.FindById(userId)
	helper.ErrorPanic(err)

	userResponse := response.UserResponse{
		Id:       userData.Id,
		Name:     userData.Name,
		Email:    userData.Email,
		Password: userData.Password,
		Age:      userData.Age,
	}
	return userResponse
}

func (u *GormSqlServerServiceImpl) FindAll() []response.UserResponse {
	result := u.GormSqlServerRepository.FindAll()

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:       value.Id,
			Name:     value.Name,
			Email:    value.Email,
			Password: value.Password,
			Age:      value.Age,
		}
		users = append(users, user)
	}

	return users
}

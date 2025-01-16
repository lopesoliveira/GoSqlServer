package service

import (
	"github.com/lopesoliveira/GoSqlServer/data/request"
	"github.com/lopesoliveira/GoSqlServer/data/response"
	"github.com/stretchr/testify/mock"
)

type GormSqlServerServiceMock struct {
	mock.Mock
}

func (m *GormSqlServerServiceMock) Create(user request.CreateUserRequest) {
	//TODO implement me
	panic("implement me")
}

func (m *GormSqlServerServiceMock) Update(user request.UpdateUserRequest) {
	//TODO implement me
	panic("implement me")
}

func (m *GormSqlServerServiceMock) Delete(userId int) response.DeleteResponse {
	//TODO implement me
	panic("implement me")
}

func (m *GormSqlServerServiceMock) FindAll() []response.UserResponse {
	//TODO implement me
	panic("implement me")
}

func (m *GormSqlServerServiceMock) FindById(userId int) response.UserResponse {
	args := m.Called(userId)
	return args.Get(0).(response.UserResponse)
}

// 1 - Create a mock for the GormSqlServerService

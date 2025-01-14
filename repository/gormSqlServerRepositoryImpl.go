package repository

import (
	"errors"
	"github.com/lopesoliveira/GoSqlServer/data/request"
	"github.com/lopesoliveira/GoSqlServer/helper"
	"github.com/lopesoliveira/GoSqlServer/model"
	"gorm.io/gorm"
)

type GormSqlServerRepositoryImpl struct {
	Db *gorm.DB
}

func NewGormSqlServerRepositoryImpl(Db *gorm.DB) GormSqlServerRepository {
	return &GormSqlServerRepositoryImpl{Db: Db}
}

func (u GormSqlServerRepositoryImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (u GormSqlServerRepositoryImpl) Update(user model.User) {
	var updatedUser = request.UpdateUserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
	}
	result := u.Db.Model(&user).Updates(updatedUser)
	helper.ErrorPanic(result.Error)
}

func (u GormSqlServerRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

func (u GormSqlServerRepositoryImpl) FindById(userId int) (model.User, error) {
	var user model.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (u GormSqlServerRepositoryImpl) FindAll() []model.User {
	var users []model.User
	results := u.Db.Find(&users)
	helper.ErrorPanic(results.Error)
	return users
}

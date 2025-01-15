package repository

import (
	"github.com/lopesoliveira/GoSqlServer/model"
)

type GormSqlServerRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int) (bool, error)
	FindById(userId int) (model.User, error)
	FindAll() []model.User
}

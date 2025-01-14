package repository

import (
	"github.com/lopesoliveira/GoSqlServer/model"
)

type GormSqlServerRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (user model.User, err error)
	FindAll() []model.User
}

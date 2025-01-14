package model

type User struct {
	Id       int    `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"type:varchar(500);not null"`
	Email    string `gorm:"type:varchar(500);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Age      int    `gorm:"type:int(11);not null"`
}

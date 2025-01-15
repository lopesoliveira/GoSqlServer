package model

type User struct {
	Id       int    `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"type:varchar(500);not null"`
	Email    string `gorm:"type:varchar(500);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Age      int    `gorm:"not null"`
}

// TableName overrides the table name used by User to `User`
// Gorm will pluralize table names. If I don't add this it will try to access table Users
func (User) TableName() string {
	return "User"
}

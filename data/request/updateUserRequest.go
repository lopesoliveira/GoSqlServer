package request

type UpdateUserRequest struct {
	Id       int    `validate:"required" json:"id"`
	Name     string `validate:"required,min=1,max=500" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=1,max=500" json:"password"`
	Age      int    `validate:"required" json:"age"`
}

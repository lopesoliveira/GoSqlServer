package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lopesoliveira/GoSqlServer/data/request"
	"github.com/lopesoliveira/GoSqlServer/data/response"
	"github.com/lopesoliveira/GoSqlServer/helper"
	"github.com/lopesoliveira/GoSqlServer/service"
	"net/http"
	"strconv"
)

type UsersController struct {
	gormSqlServerService service.GormSqlServerService
}

func NewUsersController(service service.GormSqlServerService) *UsersController {
	return &UsersController{
		gormSqlServerService: service,
	}
}

// CreateUser		godoc
// @Summary			Create User
// @Description		Save user data in Db.
// @Param			user body request.CreateUserRequest true "Create user"
// @Produce			application/json
// @Tags			user
// @Success			200 {object} response.UserResponse{}
// @Router			/user [post]
func (controller *UsersController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.gormSqlServerService.Create(createUserRequest)

	webResponse := response.UserResponse{}

	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTags		godoc
// @Summary			Update user
// @Description		Update user data.
// @Param			userId path string true "update user by id"
// @Param			user body request.UpdateUserRequest true  "Update user"
// @Tags			user
// @Produce			application/json
// @Success			200 {object} response.UserResponse{}
// @Router			/user/{userId} [patch]
func (controller *UsersController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	updateUserRequest.Id = id

	controller.gormSqlServerService.Update(updateUserRequest)

	webResponse := response.UserResponse{}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteUser		godoc
// @Summary			Delete user
// @Description		Remove user data by id.
// @Param           userId path int true "User Id"
// @Produce			application/json
// @Tags			user
// @Success			200 {object} response.DeleteResponse{}
// @Router			/user/{userId} [delete]
func (controller *UsersController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	success := controller.gormSqlServerService.Delete(id)

	//success := response.UserResponse{}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, success)

}

// FindByIdTags 		godoc
// @Summary				Get Single user by id.
// @Param				userId path string true "find user by id"
// @Description			Return the user whose userId value matches id.
// @Produce				application/json
// @Tags				user
// @Success				200 {object} response.UserResponse{}
// @Router				/user/{userId} [get]
func (controller *UsersController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.gormSqlServerService.FindById(id)

	/*webResponse := response.UserResponse{
		Id:       userResponse.Id,
		Name:     userResponse.Name,
		Email:    userResponse.Email,
		Password: userResponse.Password,
		Age:      userResponse.Age,
	}*/

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, userResponse)
}

// FindAllTags 		godoc
// @Summary			Get All users.
// @Description		Return list of users.
// @Tags			users
// @Success			200 {object} response.UserResponse{}
// @Router			/user [get]
func (controller *UsersController) FindAll(ctx *gin.Context) {
	userResponse := controller.gormSqlServerService.FindAll()

	/*webResponse := response.UserResponse{
		Id: userResponse[0].Id,
	}*/
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, userResponse)
}

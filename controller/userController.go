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

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *UsersController) Create(ctx *gin.Context) {
	createTagRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	controller.gormSqlServerService.Create(createTagRequest)

	webResponse := response.UserResponse{}

	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTags		godoc
// @Summary			Update tags
// @Description		Update tags data.
// @Param			tagId path string true "update tags by id"
// @Param			tags body request.CreateTagsRequest true  "Update tags"
// @Tags			tags
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [patch]
func (controller *UsersController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	updateUserRequest.Id = id

	controller.gormSqlServerService.Update(updateUserRequest)

	webResponse := response.UserResponse{}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTags		godoc
// @Summary			Delete tags
// @Description		Remove tags data by id.
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagID} [delete]
func (controller *UsersController) Delete(ctx *gin.Context) {
	userId := ctx.Param("tagId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.gormSqlServerService.Delete(id)

	webResponse := response.UserResponse{}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindByIdTags 		godoc
// @Summary				Get Single tags by id.
// @Param				tagId path string true "update tags by id"
// @Description			Return the tahs whoes tagId valu mathes id.
// @Produce				application/json
// @Tags				tags
// @Success				200 {object} response.Response{}
// @Router				/tags/{tagId} [get]
func (controller *UsersController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.gormSqlServerService.FindById(id)

	webResponse := response.UserResponse{
		Id: userResponse.Id,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @Success			200 {obejct} response.Response{}
// @Router			/tags [get]
func (controller *UsersController) FindAll(ctx *gin.Context) {
	userResponse := controller.gormSqlServerService.FindAll()

	webResponse := response.UserResponse{
		Id: userResponse[0].Id,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

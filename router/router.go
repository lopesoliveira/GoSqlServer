package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lopesoliveira/GoSqlServer/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter(userController *controller.UsersController) *gin.Engine {

	router := gin.Default()

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Go with Gin Gonic, Gorm, and Sql Server")
	})

	baseRouter := router.Group("/api")
	userRouter := baseRouter.Group("/user")
	userRouter.GET("", userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)

	return router
}

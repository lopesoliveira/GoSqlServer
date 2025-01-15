package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/lopesoliveira/GoSqlServer/config"
	"github.com/lopesoliveira/GoSqlServer/controller"
	_ "github.com/lopesoliveira/GoSqlServer/docs"
	"github.com/lopesoliveira/GoSqlServer/helper"
	"github.com/lopesoliveira/GoSqlServer/model"
	"github.com/lopesoliveira/GoSqlServer/repository"
	"github.com/lopesoliveira/GoSqlServer/router"
	"github.com/lopesoliveira/GoSqlServer/service"
	"net/http"
)

// @title 	Tag Service API
// @version	1.0
// @description A User service API in Go using Gin framework

// @host 	localhost:8082
// @BasePath /api
func main() {

	fmt.Println("Hello World")
	//var myDictionary = getDictionary()
	//config.CheckDbConnection()
	db := config.DatabaseConnection()
	//db := config.SetupDatabase("GormSqlServer")
	validate := validator.New()

	db.Table("User").AutoMigrate(&model.User{})

	// Repository
	gormSqlServerRepository := repository.NewGormSqlServerRepositoryImpl(db)

	// Service
	gormSqlServerService := service.NewGormSqlServerServiceImpl(gormSqlServerRepository, validate)

	// Controller
	userController := controller.NewUsersController(gormSqlServerService)

	// Router
	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:    ":8082",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

	/*
		port := ":8082"
		router := gin.Default()

		router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, myDictionary)
		})

		// Set the HTTP port to 8080 and handle the error
		if err := router.Run(port); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	*/
}

/*
func getDictionary() map[int]string {
	dict := map[int]string{1: "Joao", 2: "Carlos", 3: "Soares", 4: "Lopes", 5: "de", 6: "Oliveira"}
	return dict
}
*/

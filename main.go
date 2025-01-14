package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lopesoliveira/GoSqlServer/config"
	"github.com/lopesoliveira/GoSqlServer/model"
	"github.com/lopesoliveira/GoSqlServer/repository"
	"net/http"
)

func main() {

	fmt.Println("Hello World")
	var myDictionary = getDictionary()
	//config.CheckDbConnection()
	config.DatabaseConnection()
	db := config.SetupDatabase("GormSqlServer")
	validate := validator.New()

	db.Table("User").AutoMigrate(&model.User{})

	// Repository
	gormSqlServerRepository := repository.NewGormSqlServerRepositoryImpl(db)

	// Service

	port := ":8082"
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, myDictionary)
	})

	// Set the HTTP port to 8080 and handle the error
	if err := router.Run(port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func getDictionary() map[int]string {
	dict := map[int]string{1: "Joao", 2: "Carlos", 3: "Soares", 4: "Lopes", 5: "de", 6: "Oliveira"}
	return dict
}

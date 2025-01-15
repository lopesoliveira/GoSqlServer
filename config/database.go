package config

import (
	"fmt"
	"github.com/lopesoliveira/GoSqlServer/helper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "5CD114LJ30"
	port     = 1433
	user     = "sonarsqlauth"
	password = "Eqs2024."
	dbName   = "GormSqlServer"
)

func DatabaseConnection() *gorm.DB {
	//sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	sqlInfo := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, host, port, dbName)

	db, err := gorm.Open(sqlserver.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err) //If sqlInfo is ok and connects it proceeds; otherwise shows error and program stops
	return db
}

func SetupDatabase(dbName string) *gorm.DB {
	// Connect to the master database to create a new database
	masterDsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=master", user, password, host, port)
	masterDb, err := gorm.Open(sqlserver.Open(masterDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to master database: %v", err)
	}

	// Create the new database
	createDbQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	if err := masterDb.Exec(createDbQuery).Error; err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	fmt.Println("Database created successfully")

	// Connect to the new database
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, host, port, dbName)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to new database: %v", err)
	}
	return db
}

/*
func CheckDbConnection() {
	dsn := "sqlserver://sonarsqlauth:Eqs2024@5CD114LJ30:1433?database=sonar"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Db connection failed")
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Db connection failed")
		os.Exit(1)
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("Db connection failed")
		os.Exit(1)
	}

	fmt.Println("Connected to Db")
}
*/

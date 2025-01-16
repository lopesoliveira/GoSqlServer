package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lopesoliveira/GoSqlServer/data/response"
	"github.com/lopesoliveira/GoSqlServer/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindById(t *testing.T) {
	// Create a mock service
	mockService := new(service.GormSqlServerServiceMock)
	userController := NewUsersController(mockService)

	// Define the expected response
	expectedUser := response.UserResponse{
		Id:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password",
		Age:      30,
	}

	// Set up the mock to return the expected response
	mockService.On("FindById", 1).Return(expectedUser)

	// Create a new Gin context
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user/:userId", userController.FindById)

	// Create a new request to pass to the handler
	req, _ := http.NewRequest(http.MethodGet, "/user/1", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":1,"name":"John Doe","email":"john.doe@example.com","password":"password","age":30}`, w.Body.String())

	// Assert that the mock was called as expected
	mockService.AssertExpectations(t)

}

// 2 - Write the unit test for UsersController

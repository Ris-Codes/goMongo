package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ris-Codes/goMongo/models"
	"github.com/gin-gonic/gin"
)

// Mock UserService
type MockUserService struct {
	CreateUserFunc func(user *models.User) error
	GetUserFunc    func(name *string) (*models.User, error)
	GetAllFunc     func() ([]*models.User, error)
	UpdateUserFunc func(user *models.User) error
	DeleteUserFunc func(name *string) error
}

func (m *MockUserService) CreateUser(user *models.User) error {
	return m.CreateUserFunc(user)
}

func (m *MockUserService) GetUser(name *string) (*models.User, error) {
	return m.GetUserFunc(name)
}

func (m *MockUserService) GetAll() ([]*models.User, error) {
	return m.GetAllFunc()
}

func (m *MockUserService) UpdateUser(user *models.User) error {
	return m.UpdateUserFunc(user)
}

func (m *MockUserService) DeleteUser(name *string) error {
	return m.DeleteUserFunc(name)
}

func TestCreateUser(t *testing.T) {
	mockService := &MockUserService{
		CreateUserFunc: func(user *models.User) error {
			return nil
		},
	}
	userController := New(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/user/create", userController.CreateUser)

	user := models.User{Name: "John Doe", Age: 30, Address: models.Address{
		State: "Any State",
		City: "Anytown", 
		Pincode: 673631}}

	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	expected := `{"message":"Success"}`
	if resp.Body.String() != expected {
		t.Fatalf("Expected response body %s but got %s", expected, resp.Body.String())
	}
}

func TestGetUser(t *testing.T) {
	mockService := &MockUserService{
		GetUserFunc: func(name *string) (*models.User, error) {
			return &models.User{Name: *name, Age: 30, Address: models.Address{
				State: "Any State",
				City: "Anytown", 
				Pincode: 673631}}, nil
		},
	}
	userController := New(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user/get/:name", userController.GetUser)

	req, _ := http.NewRequest(http.MethodGet, "/user/get/John Doe", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	expectedUser := models.User{Name: "John Doe", Age: 30, Address: models.Address{
		State: "Any State",
		City: "Anytown", 
		Pincode: 673631}}
	expected, _ := json.Marshal(expectedUser)
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	if resp.Body.String() != string(expected) {
		t.Fatalf("Expected response body %s but got %s", string(expected), resp.Body.String())
	}
}

func TestGetAllUsers(t *testing.T) {
	mockService := &MockUserService{
		GetAllFunc: func() ([]*models.User, error) {
			return []*models.User{
				{Name: "John Doe", Age: 30, Address: models.Address{
					State: "Any State",
					City: "Anytown", 
					Pincode: 673631}},
				{Name: "Jane Doe", Age: 25, Address: models.Address{
					State: "Any State",
					City: "Anytown", 
					Pincode: 673631}},
			}, nil
		},
	}
	userController := New(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user/getall", userController.GetAll)

	req, _ := http.NewRequest(http.MethodGet, "/user/getall", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	expectedUsers := []*models.User{
		{Name: "John Doe", Age: 30, Address: models.Address{
			State: "Any State",
			City: "Anytown", 
			Pincode: 673631}},
		{Name: "Jane Doe", Age: 25, Address: models.Address{
			State: "Any State",
			City: "Anytown", 
			Pincode: 673631}},
	}
	expected, _ := json.Marshal(expectedUsers)
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	if resp.Body.String() != string(expected) {
		t.Fatalf("Expected response body %s but got %s", string(expected), resp.Body.String())
	}
}

func TestUpdateUser(t *testing.T) {
	mockService := &MockUserService{
		UpdateUserFunc: func(user *models.User) error {
			return nil
		},
	}
	userController := New(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PATCH("/user/update", userController.UpdateUser)

	user := models.User{Name: "John Doe", Age: 30, Address: models.Address{
		State: "Any State",
		City: "Anytown", 
		Pincode: 673631}}

	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPatch, "/user/update", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	expected := `{"message":"Success"}`
	if resp.Body.String() != expected {
		t.Fatalf("Expected response body %s but got %s", expected, resp.Body.String())
	}
}

func TestDeleteUser(t *testing.T) {
	mockService := &MockUserService{
		DeleteUserFunc: func(name *string) error {
			return nil
		},
	}
	userController := New(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/user/delete/:name", userController.DeleteUser)

	req, _ := http.NewRequest(http.MethodDelete, "/user/delete/John Doe", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	expected := `{"message":"Success"}`
	if resp.Body.String() != expected {
		t.Fatalf("Expected response body %s but got %s", expected, resp.Body.String())
	}
}
package handler_test

import (
	"app/internal/sections/handler"
	"app/internal/sections/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDepartmentRepository struct {
	mock.Mock
}

func (m *MockDepartmentRepository) GetAllDepartments() ([]models.GetDepartmentsRequest, error) {
	args := m.Called()
	return args.Get(0).([]models.GetDepartmentsRequest), args.Error(1)
}

func (m *MockDepartmentRepository) CreateDepartment(input models.CreateDepartmentRequest) (models.DepartmentModel, error) {
	args := m.Called(input)
	return args.Get(0).(models.DepartmentModel), args.Error(1)
}

func (m *MockDepartmentRepository) EditDepartment(id string, input models.EditDepartmentRequest) (models.DepartmentModel, error) {
	args := m.Called(input)
	return args.Get(0).(models.DepartmentModel), args.Error(1)
}

func (m *MockDepartmentRepository) DeactivateDepartment(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockDepartmentRepository) ReactivateDepartment(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupRouter(handler *handler.DepartmentHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/departments", handler.GetAllDepartments)
	r.POST("/departments", handler.CreateDepartment)
	r.PUT("/departments/:id", handler.EditDepartment)
	r.DELETE("/departments/:id/deactivate", handler.DeactivateDepartment)
	r.PUT("/departments/:id/reactivate", handler.ReactivateDepartment)
	return r
}

func TestGetAllDepartment(t *testing.T) {
	mockRepo := new(MockDepartmentRepository)
	mockRepo.On("GetAllDepartments").Return([]models.GetDepartmentsRequest{{ID: "1", Name: "TI"}}, nil)

	h := handler.NewDepartmentHandler(mockRepo)
	r := setupRouter(h)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/departments", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateDepartment(t *testing.T) {
	mockRepo := new(MockDepartmentRepository)
	input := models.CreateDepartmentRequest{Name: "TI", Code: "101"}
	created := models.DepartmentModel{ID: "1", Name: "TI", Code: "101", IsActive: true}

	mockRepo.On("CreateDepartment", input).Return(created, nil)

	h := handler.NewDepartmentHandler(mockRepo)
	r := setupRouter(h)

	body, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/departments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestEditDepartment(t *testing.T) {
	mockRepo := new(MockDepartmentRepository)
	id := "123"
	input := models.EditDepartmentRequest{Name: ptrString("Novo nome")}
	updated := models.DepartmentModel{ID: "123", Name: "Novo nome"}

	mockRepo.On("EditDepartment", id, input).Return(updated, nil)

	h := handler.NewDepartmentHandler(mockRepo)
	r := setupRouter(h)

	body, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/departments/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeactivateDepartment(t *testing.T) {
	mockRepo := new(MockDepartmentRepository)
	id := "123"
	mockRepo.On("DeactivateDepartment", id).Return(nil)

	h := handler.NewDepartmentHandler(mockRepo)
	r := setupRouter(h)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/departments/"+id+"/deactivate", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestReactivateDepartment(t *testing.T) {
    mockRepo := new(MockDepartmentRepository)
    id := "123"
    
    mockRepo.On("ReactivateDepartment", id).Return(nil)

    h := handler.NewDepartmentHandler(mockRepo)
    r := setupRouter(h)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodPut, "/departments/"+id+"/reactivate", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.JSONEq(t, `{"message":"Department reactivated successfully"}`, w.Body.String())
    mockRepo.AssertExpectations(t)
}

func ptrString(s string) *string {
	return &s
}

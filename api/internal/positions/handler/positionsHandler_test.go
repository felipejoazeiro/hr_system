package handler_test

import (
	"app/internal/positions/handler"
	"app/internal/positions/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPositionRepository struct {
	mock.Mock
}

func (m *MockPositionRepository) CreatePosition(input models.CreatePosition) (models.PositionModel, error) {
	args := m.Called(input)
	return args.Get(0).(models.PositionModel), args.Error(1)
}

func (m *MockPositionRepository) EditPosition(id string, input models.EditPosition) (models.PositionModel, error) {
	args := m.Called(id, input)
	return args.Get(0).(models.PositionModel), args.Error(1)
}

func (m *MockPositionRepository) GetAllPositions() ([]models.PositionModel, error) {
	args := m.Called()
	return args.Get(0).([]models.PositionModel), args.Error(1)
}

func (m *MockPositionRepository) DeactivatePosition(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPositionRepository) ReactivePosition(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupRouter(handler *handler.PositionHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/positions", handler.CreatePosition)
	r.GET("/positions", handler.GetAllPositions)
	r.PUT("/positions/:id", handler.EditPosition)
	r.PUT("/positions/:id", handler.DeactivatePosition)
	r.PUT("/positions/:id", handler.ReactivePosition)
	return r
}

func TestCreatePosition(t *testing.T) {
	mockRepo := new(MockPositionRepository)
	handler := handler.NewPositionHandler(mockRepo)
	router := setupRouter(handler)

	input := models.CreatePosition{Name: "TI", Code: "1"}
	expected := models.PositionModel{ID: "1", Name: "TI", IsActive: true, Code: 1}

	mockRepo.On("CreatePosition", input).Return(expected, nil)

	body, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/positions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetAllPositions(t *testing.T) {
	mockRepo := new(MockPositionRepository)
	handler := handler.NewPositionHandler(mockRepo)
	router := setupRouter(handler)

	mockRepo.On("GetAllPositions").Return([]models.PositionModel{
		{ID: "1", Name: "TI", Code: 101, IsActive: true},
	}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/positions", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestEditPosition(t *testing.T) {
	mockRepo := new(MockPositionRepository)
	handler := handler.NewPositionHandler(mockRepo)
	router := setupRouter(handler)

	id := "1"
	newName := "RH"
	input := models.EditPosition{Name: &newName}
	expected := models.PositionModel{ID: "1", Name: "RH", Code: 1, IsActive: true}

	mockRepo.On("EditPosition", id, input).Return(expected, nil)

	body, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "positions/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeactivatePosition(t *testing.T) {
	mockRepo := new(MockPositionRepository)
	handler := handler.NewPositionHandler(mockRepo)
	router := setupRouter(handler)

	id := "1"
	mockRepo.On("DeactivatePosition", id).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/positions/"+id+"/deactivate", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestReactivatePosition(t *testing.T) {
	mockRepo := new(MockPositionRepository)
	handler := handler.NewPositionHandler(mockRepo)
	router := setupRouter(handler)

	id := "1"
	mockRepo.On("ReactivatePosition", id).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/positions/"+id+"/reactivate", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

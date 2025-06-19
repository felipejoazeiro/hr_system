package hr

import(
	//"errors"
	"sync"
	//"time"
	"app/internal/hr/models"
)

var (
	funcionarios []models.EmployeeModel
	mu sync.Mutex
)
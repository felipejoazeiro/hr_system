package hr

import(
	"errors"
	"sync"
	"time"
)

var (
	funcionarios []Funcionario
	mu sync.Mutex
)
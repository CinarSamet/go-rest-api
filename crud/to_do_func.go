package to_do_func

import (
	"go-rest-api/models"
	"sync"
)

var (
	userTodos      = make(map[string]map[int]models.ToDo)
	userIDCounters = make(map[string]int)
	mutex          = &sync.Mutex{}
)

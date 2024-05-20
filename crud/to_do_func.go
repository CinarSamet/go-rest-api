package to_do_func

import (
	"encoding/json"
	"go-rest-api/models"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

// create maps
var (
	userTodos      = make(map[string]map[int]models.ToDo)
	userIDCounters = make(map[string]int)
	mutex          = &sync.Mutex{}
)

func ListTodos(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	mutex.Lock()
	todos := userTodos[username]
	mutex.Unlock()
	filteredTodos := []models.ToDo{}
	for _, todo := range todos {
		if todo.DeletedOn.IsZero() {
			filteredTodos = append(filteredTodos, todo)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\n")
	encoder.Encode(filteredTodos)
}

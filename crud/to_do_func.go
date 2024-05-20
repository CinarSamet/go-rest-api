package to_do_func

import (
	"encoding/json"
	"go-rest-api/models"
	"net/http"
	"sync"
	"time"

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
func ListAllTodos(w http.ResponseWriter, r *http.Request) {
	allTodos := []models.ToDo{}
	mutex.Lock()
	for _, todos := range userTodos {
		for _, todo := range todos {
			if !todo.DeletedOn.IsZero() {
				// If the ToDo is deleted, append ' (deleted)' to the description.
				todo.Description += " (deleted)"
			}
			allTodos = append(allTodos, todo)
		}
	}
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\n")
	encoder.Encode(allTodos)
}
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	var todo models.ToDo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	if userTodos[username] == nil {
		userTodos[username] = make(map[int]models.ToDo)
	}
	todo.ID = userIDCounters[username] + 1
	userIDCounters[username]++
	todo.CreatedOn = time.Now()
	todo.User = username
	userTodos[username][todo.ID] = todo
	json.NewEncoder(w).Encode(todo)
}
func AdminCreateOwnTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	username := "admin"
	mutex.Lock()
	defer mutex.Unlock()
	if userTodos[username] == nil {
		userTodos[username] = make(map[int]models.ToDo)
	}
	todo.ID = userIDCounters[username] + 1
	userIDCounters[username]++
	todo.CreatedOn = time.Now()
	todo.User = username
	userTodos[username][todo.ID] = todo
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
	w.Write([]byte("ToDo successfully Created"))
}

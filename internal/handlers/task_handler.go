package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"go_server/internal/services"

	"go.mongodb.org/mongo-driver/bson"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler{
	return &TaskHandler{
		service: service,
	}
}

// CreateTaskHandler godoc
// @Summary Create a new task
// @Description Creates a task with a title
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body object true "Task Payload"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]string
// @Router /tasks [post]
func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		Title string `json:"title"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(body.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GetTasksHandler godoc
// @Summary Get all tasks
// @Description Retrieves all tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func (h *TaskHandler)GetTasksHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, "Metho not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := h.service.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// UpdateTaskHandler godoc
// @Summary Update a task
// @Description Updates task fields
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body object true "Update Payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")

	var updateData bson.M

	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.UpdateTask(id, updateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task Updated Successfully",
	})
}

// DeleteTaskHandler godoc
// @Summary Delete a task
// @Description Deletes a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")

	err := h.service.DeleteTask(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task Deleted Successfully",
	})
}

func (h *TaskHandler) TaskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		h.CreateTaskHandler(w, r)

	case http.MethodGet:
		h.GetTasksHandler(w, r)

	case http.MethodPut:
		h.UpdateTaskHandler(w, r)

	case http.MethodDelete:
		h.DeleteTaskHandler(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

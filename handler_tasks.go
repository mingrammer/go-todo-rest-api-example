package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func (app *App) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := GetProjectOr404(app.DB, projectTitle, w, r)
	if project == nil {
		return
	}

	tasks := []Task{}
	if err := app.DB.Model(&project).Related(&tasks).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tasks)
}

func (app *App) CreateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := GetProjectOr404(app.DB, projectTitle, w, r)
	if project == nil {
		return
	}

	task := Task{ProjectID: project.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := app.DB.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, task)
}

func (app *App) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := GetProjectOr404(app.DB, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := GetTaskOr404(app.DB, id, w, r)
	if task == nil {
		return
	}
	respondJSON(w, http.StatusOK, task)
}

func (app *App) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := GetProjectOr404(app.DB, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := GetTaskOr404(app.DB, id, w, r)
	if task == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := app.DB.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, task)
}

func (app *App) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := GetProjectOr404(app.DB, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := GetTaskOr404(app.DB, id, w, r)
	if task == nil {
		return
	}

	if err := app.DB.Delete(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func (app *App) CompleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := GetProjectOr404(app.DB, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := GetTaskOr404(app.DB, id, w, r)
	if task == nil {
		return
	}

	task.Complete()
	if err := app.DB.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, task)
}

func (app *App) UndoTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectTitle := vars["title"]
	project := GetProjectOr404(app.DB, projectTitle, w, r)
	if project == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	task := GetTaskOr404(app.DB, id, w, r)
	if task == nil {
		return
	}

	task.Undo()
	if err := app.DB.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, task)
}

// GetTaskOr404 gets a task instance if exists, or respond the 404 error otherwise
func GetTaskOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *Task {
	task := Task{}
	if err := db.First(&task, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &task
}

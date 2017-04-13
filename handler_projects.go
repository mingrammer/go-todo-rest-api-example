package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func (app *App) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := []Project{}
	app.DB.Find(&projects)
	respondJSON(w, http.StatusOK, projects)
}

func (app *App) CreateProject(w http.ResponseWriter, r *http.Request) {
	project := Project{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := app.DB.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

func (app *App) GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(app.DB, title, w, r)
	if project == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func (app *App) UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(app.DB, title, w, r)
	if project == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := app.DB.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func (app *App) DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(app.DB, title, w, r)
	if project == nil {
		return
	}
	if err := app.DB.Delete(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func (app *App) ArchiveProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(app.DB, title, w, r)
	if project == nil {
		return
	}
	project.Archive()
	if err := app.DB.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func (app *App) RestoreProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := GetProjectOr404(app.DB, title, w, r)
	if project == nil {
		return
	}
	project.Restore()
	if err := app.DB.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// GetProjectOr404 gets a project instance if exists, or respond the 404 error otherwise
func GetProjectOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *Project {
	project := Project{}
	if err := db.First(&project, Project{Title: title}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &project
}

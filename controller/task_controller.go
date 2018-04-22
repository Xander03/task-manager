package controller

import (
	. "../dao"
	. "../model"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type TaskController struct {
	Controller Controller
	Repository TaskRepository
}

func (c *TaskController) AddTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		c.Controller.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := c.Repository.AddTask(task); err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, map[string]string{"payload": "added successfully"})
}

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]
	tasks, err := c.Repository.GetAllTasks(userId)
	if err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, tasks)
}

func (c *TaskController) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	task, err := c.Repository.GetTaskById(id)
	if err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, task)
}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		c.Controller.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := c.Repository.UpdateTask(task); err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, map[string]string{"payload": "updated successfully"})
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := c.Repository.DeleteTask(id); err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, map[string]string{"payload": "deleted successfully"})
}
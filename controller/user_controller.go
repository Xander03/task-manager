package controller

import (
	. "../dao"
	. "../model"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type UserController struct {
	Controller Controller
	Repository UserRepository
}

func (c *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		c.Controller.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := c.Repository.AddUser(user); err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, map[string]string{"payload": "Added successfully"})
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Repository.GetAllUsers()
	if err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, users)
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := c.Repository.GetUserById(id)
	if err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, user)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		c.Controller.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if err := c.Repository.UpdateUser(user); err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, map[string]string{"payload": "updated successfully"})
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := c.Repository.DeleteUser(id); err != nil {
		c.Controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.Controller.RespondWithSuccess(w, map[string]string{"payload": "deleted successfully"})
}
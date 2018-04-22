package router

import (
	"github.com/gorilla/mux"
	. "../dao"
	. "../controller"
)

var taskController = &TaskController{
	Controller: Controller{},
	Repository: TaskRepository{
		Repository: Repository{},
	},
}

func NewTaskRouter() *mux.Router {
	router := mux.NewRouter()
	router.Path("/").HandlerFunc(taskController.AddTask).Methods("POST")
	router.Path("/{user_id}").HandlerFunc(taskController.GetAllTasks).Methods("POST")
	router.Path("/{id}").HandlerFunc(taskController.GetTaskById).Methods("GET")
	router.Path("/").HandlerFunc(taskController.UpdateTask).Methods("PUT")
	router.Path("/{id}").HandlerFunc(taskController.DeleteTask).Methods("DELETE")

	return router
}
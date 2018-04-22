package router

import (
	"github.com/gorilla/mux"
	. "../controller"
	. "../dao"
)

var userController = &UserController{
	Controller: Controller{},
	Repository: UserRepository{
		Repository: Repository{},
	},
}

func NewUserRouter() *mux.Router {
	router := mux.NewRouter()
	router.Path("/").HandlerFunc(userController.AddUser).Methods("POST")
	router.Path("/").HandlerFunc(userController.GetAllUsers).Methods("GET")
	router.Path("/{id}").HandlerFunc(userController.GetUserById).Methods("GET")
	router.Path("/").HandlerFunc(userController.UpdateUser).Methods("PUT")
	router.Path("/{id}").HandlerFunc(userController.DeleteUser).Methods("DELETE")

	return router
}
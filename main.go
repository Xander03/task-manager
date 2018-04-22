package main

import (
	"./router"

	"net/http"
	"github.com/gorilla/mux"
)

func main() {


	rootRouter := mux.NewRouter()
	router.NewUserRouter(rootRouter.PathPrefix("/users").Subrouter())
	router.NewTaskRouter(rootRouter.PathPrefix("/tasks").Subrouter())

	http.ListenAndServe(":3000", rootRouter)
}
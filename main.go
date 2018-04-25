package main

import (
	"./router"

	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {

	rootRouter := mux.NewRouter()
	router.NewUserRouter(rootRouter.PathPrefix("/users").Subrouter())
	router.NewTaskRouter(rootRouter.PathPrefix("/tasks").Subrouter())

	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"PUT", "GET", "OPTIONS", "DELETE", "POST"})

	http.ListenAndServe(":8080", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(rootRouter))
}
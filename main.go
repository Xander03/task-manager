package main

import (
	"net/http"
	"./router"
)

func main() {
	userRouter := router.NewTaskRouter()

	http.ListenAndServe(":3000", userRouter)
}
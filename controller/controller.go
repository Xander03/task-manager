package controller

import (
	"net/http"
	"encoding/json"
)

type Controller struct {}

func (c *Controller) RespondWithSuccess(w http.ResponseWriter, payload interface{}) {
	json.NewEncoder(w).Encode(payload)
}

func (c *Controller) RespondWithError(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	c.RespondWithSuccess(w, payload)
}

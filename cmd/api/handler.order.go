package main

import (
	"net/http"
	"theaveasso/uwo/pkg"
)

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
    body := map[string]interface{}{"message": "create order"}
    pkg.WriteJSON(w, http.StatusCreated, body)
}

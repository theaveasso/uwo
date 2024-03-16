package main

import (
	"net/http"
	"theaveasso/uwo/pkg"
)

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
    body := map[string]interface{}{"message": "create order"}
    pkg.WriteJSON(w, http.StatusCreated, body)
}

func listOrdersHandler(w http.ResponseWriter, r *http.Request) {
    body := map[string]interface{}{"message": "list orders"}
    pkg.WriteJSON(w, http.StatusOK, body)
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
    body := map[string]interface{}{"message": "get order"}
    pkg.WriteJSON(w, http.StatusOK, body)
}

func deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
    body := map[string]interface{}{"message": "delete order"}
    pkg.WriteJSON(w, http.StatusOK, body)
}

func updateOrderHandler(w http.ResponseWriter, r *http.Request) {
    body := map[string]interface{}{"message": "update order"}
    pkg.WriteJSON(w, http.StatusOK, body)
}

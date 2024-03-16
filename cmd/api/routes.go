package main

import "net/http"

func setupRoutes() http.Handler {
    r := http.NewServeMux()

    r.HandleFunc("GET /", homepageHandler)
    r.HandleFunc("GET /healthcheck", healthcheckHandler)
    r.HandleFunc("POST /orders", createOrderHandler)
    r.HandleFunc("GET /orders", listOrdersHandler)
    r.HandleFunc("GET /orders/{id}", getOrderHandler)
    r.HandleFunc("DELETE /orders/{id}", deleteOrderHandler)
    r.HandleFunc("PUT /orders/{id}", updateOrderHandler)

    return r
}

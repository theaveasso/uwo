package main

import "net/http"

func setupRoutes() http.Handler {
    r := http.NewServeMux()

    r.HandleFunc("GET /", homepageHandler)
    r.HandleFunc("GET /healthcheck", healthcheckHandler)

    return r
}

package main

import "net/http"

func setupRoutes() http.Handler {
    r := http.NewServeMux()

    r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })

    return r
}

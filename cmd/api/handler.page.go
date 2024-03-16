package main

import "net/http"

func homepageHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello from the UWO\n"))
}

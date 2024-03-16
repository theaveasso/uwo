package main

import (
	"net/http"
	"theaveasso/uwo/pkg"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    pkg.WriteResponse(w, http.StatusOK, "Hello from the UWO\n")
}

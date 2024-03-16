package pkg

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, body string) {
    w.WriteHeader(code)
    w.Write([]byte(body))
}

func WriteJSON(w http.ResponseWriter, code int, body interface{}) {
    w.Header().Set("Content-Type", "application/json")

    jsonBody, err := json.Marshal(body)
    if err != nil {
        WriteResponse(w, http.StatusInternalServerError, "Error: encoding JSON")
        return
    }

    w.WriteHeader(code)
    w.Write(jsonBody)
}

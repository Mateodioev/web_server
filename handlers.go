package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from root\n")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home\n")
}

func HandleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the api endpoint\n")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var newUser User 

    err := decoder.Decode(&newUser)
    if err != nil {
        fmt.Fprintf(w, "error: %v", err)
        return
    }

    response, err := newUser.toJson()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "%s", err)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}


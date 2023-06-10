package main

import (
    "net/http"
    "fmt"
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


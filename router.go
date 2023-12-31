package main

import (
    "fmt"
    "net/http"
)

type Router struct {
    rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
    return &Router{
        rules: make(map[string]map[string]http.HandlerFunc),
    }
}

func (r *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
    _, exist := r.rules[path]
    handler, methodExist := r.rules[path][method]

    return handler, methodExist, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
    handler, methodExist, exists := r.FindHandler(request.Method, request.URL.Path)

    if !exists {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Not found\n")
        return
    }

    if !methodExist {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "Method no allowed")
        return
    }

    handler(w, request)
}


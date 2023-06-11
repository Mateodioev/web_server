package main

import (
    "net/http"
)

type Server struct {
    port string
    router *Router
}

func NewServer(port string) *Server {
    return &Server{
        port: port,
        router: NewRouter(),
    }
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
    _, exists := s.router.rules[path]
    if !exists {
        s.router.rules[path] = make(map[string]http.HandlerFunc)
    }
    s.router.rules[path][method] = handler
}

func (s *Server) All(path string, handler http.HandlerFunc) {
    methods := [4]string{"GET", "POST", "PUT", "DELETE"}

    for _, method := range methods {
        s.Handle(method, path, handler)
    }
}

func (s *Server) Get(path string, handler http.HandlerFunc) {
    s.Handle("GET", path, handler)
}

func (s *Server) Post(path string, handler http.HandlerFunc)  {
    s.Handle("POST", path, handler)
}

func (s *Server) Put(path string, handler http.HandlerFunc) {
    s.Handle("PUT", path, handler)
}

func (s *Server) Delete(path string, handler http.HandlerFunc) {
    s.Handle("DELETE", path, handler)
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, m := range middlewares {
        f = m(f)
    }

    return f
}

func (s *Server) Listen() error {
    http.Handle("/", s.router)

    return http.ListenAndServe(s.port, nil)
}


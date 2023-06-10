package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flag := true // only for test reasons
		log.Println("Checking authentication")

		if !flag {
			log.Println("Auth failed")
			return
		}

		log.Println("Auth success")
		next(w, r)
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Println(getLogMessage(r, start))
		}()

		next(w, r)
	}
}

// TODO: add in next version
func RestrictMethod(method string, next http.HandlerFunc) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        if r.Method != method {
            w.WriteHeader(http.StatusMethodNotAllowed)
            fmt.Fprintf(w, "Method no allowed")
            return
        }

        next(w, r)
    }
}

func getLogMessage(r *http.Request, start time.Time) string {
	return fmt.Sprintf("~ %s : path: %s | Time: %s", r.Method, r.URL.Path, time.Since(start))
}

package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	handler.ServeHTTP(w, r)
    })
}

func handler(w http.ResponseWriter, r *http.Request) {
    querystring, _ := json.Marshal(r.URL.Query())
    s := string(querystring[:])
    fmt.Fprintf(w, s)
}

func main() {
    fmt.Println("Serving echo on :8080")
    http.HandleFunc("/", handler)
    http.HandleFunc("/echo", handler)
    http.ListenAndServe(":8080", Log(http.DefaultServeMux))
}

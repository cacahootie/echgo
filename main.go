package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
    querystring, _ := json.Marshal(r.URL.Query())
    s := string(querystring[:])
    fmt.Fprintf(w, s)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/echo", handler)
    http.ListenAndServe(":8080", nil)
}

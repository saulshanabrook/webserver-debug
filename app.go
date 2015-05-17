package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
)

func port() string{
    port := os.Getenv("PORT")
    if port == "" {
        return "8080"
    }
    return port
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Printf("got request")
    fmt.Fprintf(w, "env variable TEXT=%s at path %s with port %s", os.Getenv("TEXT"), r.URL.Path, port())
}

func main() {


    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + port(), nil)
}

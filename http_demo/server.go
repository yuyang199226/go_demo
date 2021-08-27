package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

func slowHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("receive request")
    time.Sleep(1 * time.Second)
    io.WriteString(w, "I am slow!\n")
}

func main() {
    srv := http.Server{
        Addr:    ":8888",
        Handler: http.HandlerFunc(slowHandler),
    }

    if err := srv.ListenAndServe(); err != nil {
        fmt.Printf("Server failed: %s\n", err)
    }
}

package main

import (
    "fmt"
    "errors"
    "net/http"
    "time"
    "log"
)
var processing chan int
func Q() []byte{
     c := &http.Client{Timeout: 2 * time.Second}

     res, err := c.Get("http://localhost:8888")
     if err != nil {
         log.Println(err)
     }
     var r []byte

     _, err = res.Body.Read(r)
     processing <- 1
     return r

}
func request () error{
    go Q()
    select {
        case <- time.After(1 * time.Second):
            return errors.New("timeout")
        case <- processing:
            return nil

    }

}
func main() {
    processing = make(chan int)
    res := request()
    fmt.Println(res)
}

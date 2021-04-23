package main

import (
    "fmt"
    "time"
)

func show(name string) {
    fmt.Println(name)

}
func main() {
    select {
    case <- time.After(10*time.Second):
        show("name")
    
    }
}

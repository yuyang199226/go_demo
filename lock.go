package main
import (
    "sync"
    "fmt"
    "time"
)
var mu sync.RWMutex

func read() {
    mu.RLock()
    fmt.Println("get read lock")
    defer mu.RUnlock()
}
func main() {
    go func() {
    mu.Lock()
    defer mu.Unlock()
    read()
}()
time.Sleep(300*time.Second)

}

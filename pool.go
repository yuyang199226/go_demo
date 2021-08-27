package main
import (
    "fmt"
    "sync"

)
var wg = sync.WaitGroup{}
func main() {

    userCount := 20
    ch := make(chan bool, 5)
    for i := 0; i < userCount; i++ {
        wg.Add(1)
        go Read(ch, i)
    }
    wg.Wait()

    //time.Sleep(time.Second)
}

func Read(ch chan bool, i int) {
    defer wg.Done()
    ch <- true
    fmt.Printf("go func: %d\n", i)
    fmt.Printf("go func: %d\n", i)
    <- ch
}

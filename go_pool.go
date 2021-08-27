package main

import (
    "fmt"
    "sync"
    "time"

)


var wg sync.WaitGroup

func getUrl(ch chan int, num int, m map[string]string) {
    defer wg.Done()
    fmt.Println(num,m)
    time.Sleep(time.Second)
    <- ch
}
func main() {
    c := make(chan int, 4)
    m := map[string]string{"1":"asdas", "2": "ddff"}
    for i :=0;i<10;i++ {
        wg.Add(1)
        c <- 1
        go getUrl(c, i,m)
    }
    wg.Wait()
    //time.Sleep(10*time.Second)

}

package main

import (
    "fmt"
    "time"
    "runtime"
    "net/http"
)
func getSum(n int) int {
    sum:=0
    for i:=0;i<=n;i++ {
        sum+= n
        //_, err := http.Get("http://example.com/")
        http.Get("http://example.com/")
        //if err != nil {
        //    fmt.Println(err)
        //}
    }
    return sum
}

func main() {
    runtime.GOMAXPROCS(2)
    for i:=0;i<1000;i++ {
        go getSum(100)
    }
    for i:=0;i<5;i++ {
        time.Sleep(time.Second)
        fmt.Println("Hello World")
    }
    time.Sleep(20*time.Second)

}

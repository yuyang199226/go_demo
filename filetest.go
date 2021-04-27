package main

import (
    "fmt"
    "io/ioutil"
    "time"

)

func Print(p *[]int) {
    old := *p
    old = old[:len(old)-1]
    fmt.Printf("%p, %v\n", old, old)
    fmt.Printf("%p, %v\n", p, *p)

}

func main() {
    a := []int{1,2,3}
    b := &a
    Print(b)
    ioutil.WriteFile("a_1.json", []byte("hello World"), 0644)
    time.Sleep(1*time.Second)
}

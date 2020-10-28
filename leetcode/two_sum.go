package main

import "fmt"

func main() {
    a := []int{2,3,5,1,9}
    k := 7
    m := make(map[int]int)
    for i, v := range a {
        target := k - v
        if j, ok := m[target];ok {
            fmt.Println(i,j)
        } else {
            m[v] = i
        }
    }

}

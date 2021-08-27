package main

import (
    "fmt"
)

func findLongestS(s string ) string {
    //n := len(s)
    l :=0
    resL,resR :=0,0
    res := 1
    m := make(map[byte]int)
    for j:= range s {
        m[s[j]]++
        for m[s[j]] > 1 {
            m[s[l]]--
            l++
        }
        if (j-l+1) > res {
            res = j-l+1
            resL = l
            resR = j
        }
    }
    return string(s[resL:resR+1])

}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    s := "abbcdac"
    fmt.Println(s)
    res := findLongestS(s)
    fmt.Println(res)

}

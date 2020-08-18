package main
import "fmt"

func main() {

    b := map[int]int{10:7, 11:24, 12:46, 13:63,14:85,15:111,16:143,17:143}
    fmt.Println(b[55])
    a := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 24, 46, 63, 85, 111, 143, 143}
    for i:=1;i<len(a);i++ {
        a[i] = a[i] - b[i-1]
        fmt.Println(a)
    
    }

}

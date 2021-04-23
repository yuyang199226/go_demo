package main
import "fmt"
import "time"

func main() {
    m := map[int]int{1:2}
    fmt.Println(m[1])
    for i:=0;i<100;i++ {
    
    go func(){
    m[1] = 3
    //fmt.Println(m[1])
    }()
}

    time.Sleep(1*time.Second)

}

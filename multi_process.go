package main
import (
    //"fmt"
    "runtime"
    "log"

)

func bigBytes() *[]byte {
    s := make([]byte, 100000000)
    return &s
}

func main() {
    var mem runtime.MemStats
    runtime.ReadMemStats(&mem)
    //ls := make([][]byte,100000000)
    for i:=0;i<10;i++ {
                s := bigBytes()
        if s == nil {

        //ls = append(ls, []byte("this is a test demo, hello word, sunday, if the number converted back into!"))
        }
    }
    runtime.ReadMemStats(&mem)
    log.Println(mem.Alloc)
    log.Println(mem.TotalAlloc)
    log.Println(mem.HeapAlloc)
    log.Println(mem.HeapSys)






}

package main

import "fmt"
//import "time"
import "sync"
//import "sync/atomic"


func main(){
    var mutex sync.Mutex
    sum := 0
    var wg sync.WaitGroup
    var total int64
    for i := 1; i <= 10; i++ {
        sum += i
        wg.Add(1)
  go func(i int) {
      mutex.Lock()
      defer wg.Done()
      total += int64(i)
      mutex.Unlock()
  }(i)
}

wg.Wait()
fmt.Println(sum, total)
}

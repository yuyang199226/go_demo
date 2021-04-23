package main

import "fmt"

type Context struct {

    done chan int
}

func (c Context) Done() {

    <- c.done
}

func (c Context) Set(v int) {

    c.done <- v
    fmt.Println("set finish")
}
func main() {
    ctx := Context{done: make(chan int)}
    go func(ctx Context) {
        for i:=1; i< 1000; i++ {
          if i == 999 {
            fmt.Println(i)
            ctx.Done()
            fmt.Println("Done")
        }
    }
  }(ctx)
  ctx.Set(1)
}

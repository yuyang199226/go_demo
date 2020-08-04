package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 400 * time.Millisecond)
    defer cancel()

    select {
    case <- time.After(1 * time.Second):
        fmt.Println("overslept")
    case <- ctx.Done():
        fmt.Println(ctx.Err())
    
    
    }


}

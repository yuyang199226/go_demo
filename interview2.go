
package main
 
import (
	"fmt"
	"time"
)
 
func main() {
	c := make(chan int,0)
	go func() {
		for i := 1; i < 11; i+=2 {
			<- c
            fmt.Println(i)
			//奇数
			c <- 1
		}
	}()
	go func() {
		c <- 1
		for i := 2; i < 11; i+=2 {
			<- c
			//偶数
				fmt.Println(i)
		c <- 1
		}
	}()
	time.Sleep(1 * time.Second)
}

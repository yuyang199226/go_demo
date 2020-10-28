package main

 
import (
    "fmt"
    "context"
    "time"
)

func main() {
    	_, cancel := context.WithTimeout(
		context.Background(),
		(time.Second*2),
	)
	defer  cancel()
	time.Sleep(time.Second * 3)
	fmt.Println("---")


}

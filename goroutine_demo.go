package main
import "fmt"
import "sync"
import "runtime"
func main() {
	runtime.GOMAXPROCS(10)
	wg := sync.WaitGroup{}
	wg.Add(30)
	for i := 0; i < 20; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(">>i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

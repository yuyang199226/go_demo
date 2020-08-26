package main
import "time"
import "fmt"
func main() {

loc, _ := time.LoadLocation("Asia/Shanghai")
dateTime, err := time.ParseInLocation("20060102", "20060405", loc)
if err != nil {
		fmt.Printf("%v time ParseInLocation failed", )
	}
    fmt.Println(dateTime)

}

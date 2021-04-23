package main
import "fmt"
import "time"
import "github.com/jakecoffman/cron"

var cronJob = cron.New()
func conDemo() {
    spec := "*/5 * * * * ?" //每5s执行一次
    //cronJob.RemoveJob()//要删除任务使用这个方法
    cronJob.AddFunc(spec, conFun, "cronFun")
}

func HelloWorld(){
    fmt.Println("hello world")
}

func conFun() {
    time.Sleep(time.Second*3)
    fmt.Println("this is conFun Test")
}

func main() {
    //conDemo()
    cronJob.Start()
    cronJob.AddFunc("*/2 * * * * ?", HelloWorld, "hello")
    select {}
}

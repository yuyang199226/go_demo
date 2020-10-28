package main
import "fmt"
import "github.com/jakecoffman/cron"

var cronJob = cron.New()
func conDemo() {
    spec := "*/5 * * * * ?" //每5s执行一次
    //cronJob.RemoveJob()//要删除任务使用这个方法
    cronJob.AddFunc(spec, conFun, "cronFun")
    cronJob.Start()
}

func conFun() {
    fmt.Println("this is conFun Test")
}

func main() {
    conDemo()
    go cronJob.Start()
    select {}
}

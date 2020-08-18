## gin 集成 pprof

`go get github.com/DeanThompson/ginpprof`

代码里 

```
r := gin.Default()
ginpprof.Wrap(r)
```

启动服务，访问 http://127.0.0.1:8080/debug/pprof/

## mac 安装graphviz

`brew install graphviz`

点击网页上的profile 会下载一个文件名叫profile 的文件

在命令行执行 `go tool pprof -http=:8081 profile` 会浏览器打开 ，选择flame graph 

查看使用内存top20
执行命令行 `go tool pprof -inuse_space http://127.0.0.1:8080/debug/pprof/heap`。进入交互界面 输入 `top20`


将分配内存 生成svg

`go tool pprof -alloc_space -cum -svg http://127.0.0.1:8080/debug/pprof/heap?debug=1 >heap_4G.svg`

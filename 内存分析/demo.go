package main

import (
    _ "net/http/pprof"
    //"github.com/EDDYCJY/go-pprof-example/data"
)


var datas []string

func Add(str string) string {
    data := []byte(str)
    sData := string(data)
    datas = append(datas, sData)

    return sData
}

//func main() {
//    go func() {
//        for {
//            log.Println(Add("https://github.com/EDDYCJY"))
//        }
//    }()
//
//    http.ListenAndServe("0.0.0.0:6060", nil)
//}

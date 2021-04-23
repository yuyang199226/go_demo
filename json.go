package  main
import "encoding/json"
import "io/ioutil"
type Student struct {
    Name string `json:name`
    Age int `json:"age"`
}
func main() {
    res := make([]Student, 1000)
    for i:=0;i<1000;i++ {
        res[i] = Student{"asdfg", 12}
    }
    //data, errs := json.MarshalIndent(res, "", "    ")
    data, errs := json.Marshal(res)
    if errs != nil {
        panic(errs)
    }
    ioutil.WriteFile("out.json", data, 0644)

}




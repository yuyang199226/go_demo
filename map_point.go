package main

import "fmt"

type student struct {
    Name string
    Age  int
}

func pase_student() map[string]*student {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 2},
        {Name: "li", Age: 8},
        {Name: "wang", Age: 1},
    }
    for i, _ := range stus {
        stu := stus[i]
        m[stu.Name] = &stu
    }
    return m
}

func main() {
    students := pase_student()
    for k,v := range students {
        fmt.Printf("key=%v, value=%v\n", k, v)
    }
}


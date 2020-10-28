package main

import "fmt"

func isHuiwen(ls []byte) bool{
	if len(ls) > 0 {
		if ls[0] == ls[len(ls)-1] {
			return isHuiwen(ls[1:len(ls)-1])
		} else {
			return false
		}
	} else {
		return true
	}

}
func main() {
	s := []byte("09ZAazsd,dsz aAZ90")
	res := make([]byte, 0)
	for _, v:= range s {
		fmt.Println(v)
		if v >=48 && v <= 57 {
			res = append(res, v)
		}
		if v >= 65 && v <= 122 {
			res = append(res, v)
		}
	}
	fmt.Println(string(res))
	r := isHuiwen((res))
	fmt.Println(r)
}
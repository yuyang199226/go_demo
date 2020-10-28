package main

import "fmt"

func mergeSort(l1, l2 []int) {
	n1 := len(l1)
	n2 := len(l2)
	i := 0
	j := 0
	res := make([]int, 0)
	for i<n1 && j <n2 {
		if l1[i] <= l2[j] {
			res = append(res, l1[i])
			i++
		} else {
			res = append(res, l2[j])
			j++
		}
	}
	if i < n1 {
		res = append(res, l1[i:n1]...)
	}
	if j < n2 {
		res = append(res, l1[j:n2]...)
	}
	fmt.Println(res)
}
func main() {
	l1 := []int{1,2,2, 4,7,10}
	l2 := []int{0,3,4,5,6,9}
	mergeSort(l1, l2)
}

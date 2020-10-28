package main

import "fmt"
func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := m-1
    j:= n-1
    index := len(nums1)-1
    for i>=0 && j >=0 {
        if nums2[j] > nums1[i] {
            nums1[index] = nums2[j]
            index--
            j--
        } else {
            nums1[index] = nums1[i]
            index--
            i--
        }
        fmt.Println(nums1)
        fmt.Println(nums1)
    }
        for j >= 0 {
            nums1[index] = nums2[j]
            index--
            j--
        }
}


func main() {
    a := []int{0}
    b:= []int{1}
    merge(a,0,b,1)
}

package main

import "fmt"


func partition(nums []int, l, r int) int {
    x := nums[r]
    i := l-1
    for j:=l;j<r;j++ {
        if nums[j] < x {
            i++
            nums[i],nums[j] = nums[j],nums[i]
        }
    }
    i++
    nums[i],nums[r] = nums[r], nums[i]
    fmt.Println(nums)
    return i
}


func quickSort(nums []int, l, r) {
    if l < r {
        mid := partition(nums, l, r)
        quickSort(nums, l, mid)
        quickSort(nums, mid+1, r)
    }

}
func main() {
 nums := []int{3,4,9,2,1,4,7,6}
 r := len(nums)-1
 l := 0
 quickSort(nums, l, r)
 fmt.Println(nums)

}

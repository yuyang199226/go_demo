package main
import "fmt"


func partition(nums []int, l,r, k int) int {
    i:=l-1
    for j:=l;j<r;j++ {
        if nums[j] < nums[r] {
            i++
            nums[i],nums[j] = nums[j],nums[i]
        }
    }
    nums[i+1],nums[r] = nums[r],nums[i+1]
    return i+1
}
func main(){
    a := []int{2,1,5,4,2,3}
    partition(a,0,len(a)-1,3)
    fmt.Println(a)

}

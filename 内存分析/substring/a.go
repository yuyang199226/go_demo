package main

import "fmt"

// )((adaswe(as(da)))(ads))


func getLongestSubString(s string) int {
	ans := 0
	stack := make([]byte, 0)
	n:= len(s)
	//balanace := 0
	for i:=0;i<n;i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		}else if s[i] == ')' {
			fmt.Println(string(stack))
			count := 0
			for len(stack)>0 && stack[len(stack)-1] != '(' && stack[len(stack)-1] != ')'{
				stack = stack[:len(stack)-1]
				count++
			}
			if len(stack)>0 && stack[len(stack)-1] == ')' {
				count = 0
			}
			ans = max(ans, count)
			stack = append(stack, s[i])
			fmt.Println(string(stack),count)
		}else {
			stack = append(stack, s[i])
		}
	}
	fmt.Println(string(stack))
	return ans

}

func max(a, b int)int {
	if a > b {
		return a
	}
	return b
}


func main() {
	s := ")((adaswe(asdd)(da))assda)(a))"
	num := getLongestSubString(s)
	fmt.Println(num)
}


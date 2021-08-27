package main

// )((adaswe(as(da)))(qwert))


func getLongestSubString(s string) int {
	ans := 0
	stack := make([]byte, 0)
	n:= len(s)
	balanace := 0
	for i:=0;i<n;i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		}else if s[i] == ')' {
			count := 0
			for len(stack)>0 && stack[len(stack)-1] != '(' {
				stack = stack[:len(stack)-1]
				count++
			}
			ans = max(ans, count)
			stack = append(stack, s[i])
		}else {
			stack = append(stack, s[i])
		}
	}
	return ans

}

func max(a, b int)int {
	if a > b {
		return a
	}
	return b
}


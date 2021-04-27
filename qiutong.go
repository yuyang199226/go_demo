package main 



import (
    "fmt"

)

func find(n,k,m int) int{
    dp := make([][]int, k+1)
    for i:= range dp {
        dp[i] = make([]int, n+1)
    }
    for i:=0;i<= m;i++ {
        dp[1][i]=1
    }
    for i:=2;i<=k;i++ {
        for j:=0;j<=n;j++ {
            for k:=0;k<=m;k++ {
                if j >=k {
                    dp[i][j] += dp[i-1][j-k]
                }
            }
        }
    }
fmt.Println(dp)
return dp[k][n]
}

func main() {
    ans :=find(30,10,10)
    fmt.Println(ans)

}

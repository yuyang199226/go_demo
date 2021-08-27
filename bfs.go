package main

import (
    "fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode

}

func bfs(root *TreeNode) [][]int{
    if root == nil {
        return nil
    }
    res := make([][]int,0)
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue)>0 {
        n := len(queue)
        path := make([]int,0)
        for i:=0;i<n;i++ {
            node := queue[0]
            queue = queue[1:]
            path = append(path, node.Val)
            if node.Left != nil {
                queue = append(queue,node.Left)
            }
            if node.Right != nil {
                queue = append(queue,node.Right) 

            }
        }
        res = append(res, path)

    }
    return res

}

func main() {
    n1 := &TreeNode{Val: 1}
    n2 := &TreeNode{Val: 2}
    n3 := &TreeNode{Val: 3}
    n4 := &TreeNode{Val: 4}
    n1.Left = n2
    n1.Right = n3
    n2.Left = n4
    res := bfs(n1)
    fmt.Println(res)
}

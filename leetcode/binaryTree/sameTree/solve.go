package sameTree

import "ProblemSolving/leetcode/binaryTree"


func IsSameTree(p *binaryTree.TreeNode, q *binaryTree.TreeNode) bool {
    if p==nil && q==nil { return true}
    if p == nil || q == nil  || p.Val != q.Val {
        return false
    }
    return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
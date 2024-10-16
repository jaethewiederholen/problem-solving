package binaryTree

func InorderTraversal(root *TreeNode) []int {
    inorder := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode){
        if node == nil {
            return
        }
        dfs(node.Left)
        inorder = append(inorder, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return inorder
}
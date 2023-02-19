二叉树问题的通用思考过程：

1、是否可以通过遍历一遍二叉树得到答案？如果可以用一个 traverse 函数配合外部变量来实现。

2、是否可以定义一个递归函数，通过子问题（子树）的答案推导出原问题的答案？如果可以，写出这个递归函数并充分利用返回值。

3、无论使用哪一种模式，你需要明白二叉树的每一个节点需要做什么，需要在什么时候做（前中后序）。





比如前序遍历的写法：

#### 递归模式

```
func preorderTraversal(root *TreeNode) []int {
    var ans []int
    var traversal func(root *TreeNode)
    traversal = func(root *TreeNode) {
        if root == nil {
            return
        }
        ans = append(ans, root.Val)
        traversal(root.Left)
        traversal(root.Right)
    }
    traversal(root)
    return ans
}
```



#### 动态规划模式

将问题拆分为子问题：

```
func preorderTraversal(root *TreeNode) []int {
    var ans []int
    if root == nil {
        return ans
    }
    ans = append(ans, root.Val)
    ans = append(ans, preorderTraversal(root.Left)...)
    ans = append(ans, preorderTraversal(root.Right)...)
    return ans
}
```


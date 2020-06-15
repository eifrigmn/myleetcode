## [从二叉搜索树到更大和树](<https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/>)

### 题目

给出二叉**搜索**树的根节点，该二叉树的节点值各不相同，修改二叉树，使每个节点`node`的新值等于原树种大于或等于`node.val`的值之和。

提醒一下，二叉搜索树满足下列约束条件：

+ 节点的左子树仅包含键**小于**节点键的节点
+ 节点的右子树仅包含键**大于**节点键的节点
+ 左右子树也必须是二叉搜索树

**示例：**

![tree](../assets/tree.png)

~~~
输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
~~~

**提示：**

1. 树种的节点数介于`1`和`100`之间
2. 每个节点的值介于`0`和`100`之间
3. 给定的树为二叉搜索树

### 思路

将二叉搜索树的节点进行中序遍历存放于数组中

每个节点的值等于其后所有节点值及其自身之和
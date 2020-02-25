## [最小栈](<https://leetcode-cn.com/problems/min-stack/>)

### 题目

设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。

**示例：**

~~~
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
~~~

### 思路

+ 第一种

  两个栈，`l1`按序压入元素，`l2`在压入元素的同时，将栈内元素从小到大排列，每次`l1`出栈时，在`l2`中找到对应值的元素并删除，这样`l2`的栈顶元素即为`l1`栈中的最小元素

+ 第二种

  两个栈，`l1`按序压入元素，`l2`每次压入的元素与自身头结点元素相比较小的值，`l1`每次出栈时，`l2`也出栈，这样，`l2`的栈顶元素即为`l1`栈内元素的最小值
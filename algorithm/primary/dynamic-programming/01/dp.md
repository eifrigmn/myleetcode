动态规划法

根据`n`的值，由少到多寻找公式：

| n    | result                                     |                    |
| ---- | ------------------------------------------ | ------------------ |
| 1    | 1                                          | f(1) = 1           |
| 2    | 2((1,1),(2))                               | f(2) = 2           |
| 3    | 3((1,2),(2,1),(1,1,1))                     | f(3) = f(1)+f(2)=3 |
| 4    | 5((1,1,1,1),(1,1,2),(1,2,1),(2,2),(2,1,1)) | f(4) = f(2)+f(3)=5 |

由此可得状态转化方程：

~~~
f(n) = f(n-1)+f(n-2)
~~~

其中，边界条件为：

~~~
f(1) = 1
f(2) = 2
~~~

优化：

可以看出，以上状态转换方程是一个类斐波那契方程，随着`n`值的增加，该方程的计算量会非常大，而每次计算结果，仅与相邻的前两个状态值有关，因此采用尾递归方式，构造一个中间函数`ff(n, an1, an)`

其中`an1`表示上一个状态值，`an`表示当前状态值

即状态转化方程更改为：

~~~
f(n, an1, an) = f(n-1, an, an1+an)
其中，起始状态为
f(n,1,1)
n==1时，an即为所求值
~~~

~~~go
func climbStairs_dp(n int) int {
	if n < 3 {
		return n
	}
	return f(n, 1, 1)
}

func f(n, a, b int) int {
	if n == 1 {
		return b
	}
	return f(n-1, b, a+b)
}
~~~

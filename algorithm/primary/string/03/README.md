# [字符串中的第一个唯一字符](https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/34/)

## 题目

给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

### 示例

```
s = "leetcode"
返回 0

s = "loveleetcode"
返回 2
```

### 提示

你可以假定该字符串只包含小写字母。

### 思路

#### 思路一

把字符串中的所有字符存储于map中：
    其中key为字符串值，value为该字符在字符串中出现的次数
再遍历字符串，找到第一个出现次数为1的字符并输出其游标

#### 思路二

参考代码**firstUniqChar1**
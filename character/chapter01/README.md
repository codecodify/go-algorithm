# 格式化字符串
## 需求
给定一个混合数字和字母的字符串，尝试编写程序对字符串进行格式化，使得新的字符串满足如下条件：
* 相同类型的字母不能相连，即数字的左右必须是字母，字母的左右必须是数字。
* 如果无法满足上述条件，则返回空字符串。

## 思路
根据需求，字符被分为`数字型`和`字母型`这两类，那么第一步就需要对字符进行分类。
分类完成后需要重新组合，如果这两类字符的数量大于1，则不能满足条件，返回空字符串。满足条件则取出字符重新排列成新的字符串。
需要注意的是：首个字符需要从数量较多的一类取出。
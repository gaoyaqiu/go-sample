package main

import (
	"fmt"
)

/*
	位运算符，比较的是二进制
	使用数字6、11， 来测试下结果
	6:  0110
	11: 1011
---------------
	&: 如果两个数都是1，就是1，有一个不是1就是0， 那么结果是： 0010 转换为10进制为： 2
	|: 如果两个数有一个是1就是1，否则是0， 那么结果是：1111 转换为10进制为：15
	^: 如果两个数，有一个是1，才是1，否则是0， 那么结果是： 1101 转换为10进制为：13
	&^：如果两个数，第二个数是1，需要帮第一个数强制改为0，否则是1，那么结果是：0100 转换为10进制为：4
 */
func main() {
	// 一元运算符 ^ 和 二元运算符 ^ 区别： 如果两边都是数字就是二元运算符， 否则是一元运算符
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
}

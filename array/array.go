package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a)
	fmt.Println(a[len(a)-1])
	fmt.Println("------")
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("------")
	for k, v := range a {
		fmt.Println(k, v)
	}
	b := [...]int{1, 2, 3, 4}
	fmt.Println(b)
	r := [...]int{99: -1} // 数组赋值的一种方式,99是索引，-1是赋的值。给第一百号元素赋值-1，其余的元素赋值为0
	fmt.Println(r)

	c := [2]int{1, 2}
	d := [...]int{1, 2}
	e := [2]int{1, 3}
	fmt.Println(c == d, c == e, d == e) // 数组相等的判别方式，长度&&类型&&值都相等，两个数组才能相等
	// f := [3]int{1, 2}
	// fmt.Println(c == f)  编译不过，invalid operation: c == f (mismatched types [2]int and [3]int)

}

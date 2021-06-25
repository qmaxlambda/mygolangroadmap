package main

import "fmt"

func main() {
	print("132")
	fmt.Println("hw")
	//定义变量
	var i int
	i = 10

	var j = 200
	k := 100
	fmt.Println(i, j, k)
	//枚举
	const (
		A = iota + 1
		B = 11
	)
	fmt.Println(A, B)
}

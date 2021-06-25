package main

import (
	"fmt"
)

func main() {
	//数组
	//a:=[...]int{5:1}
	//fmt.Println(a)

	//for _, v := range a {
	//	fmt.Println(v)
	//}

	//数组长度和类型相同的值类型相同
	//var j = 200
	array1 := [1]int{1}
	array2 := [2]int{2, 1}
	fmt.Printf("%T\n", array1)
	fmt.Printf("%T\n", array2)
}

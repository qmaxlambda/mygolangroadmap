package main

import "fmt"

func main() {
	//切片，动态数组
	var slice []string = []string{"a", "b"}
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)
	a := make([]string, 5)
	fmt.Printf("%T\n", a)
	//数组转切片
	array := [2]string{"A", "B"}
	slice1 := array[:len(array)]
	fmt.Println(slice1)
	fmt.Printf("%T\n", slice1)

	//切片赋值
	slice2 := slice1[1:]
	fmt.Printf("%T\n", slice2)
	//切片新增
	slice2 = append(slice2, "2")
	fmt.Println(slice2)

	//拷贝
	slice3 := make([]string, len(slice2))
	copy(slice3, slice2)
	fmt.Println(slice3)
	//合并
	slice4 := []string{}
	slice4 = append(slice1, slice2...)
	fmt.Println(slice4)
	fmt.Printf("%d,%d", len(slice4), cap(slice4))

}

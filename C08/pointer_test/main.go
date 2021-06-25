package main

import "fmt"

func Swap(a *int, b *int) {
	fmt.Printf("a的地址%p,值%d\n", a, *a)
	var c = *a
	*a = *b
	*b = c
	fmt.Printf("地址%p,值%d\n", &c, c)
}

func readStruct(s *struct1) {
	fmt.Printf("a的地址%p,值%d\n", s, *s)
}

type struct1 struct {
	a int
	b int
	c *int
	d *struct1
	struct2
}

type struct2 struct {
	a int
	b int
}

func main() {
	a := 10
	b := 20
	Swap(&a, &b)
	fmt.Println(a, b)
	//快速交换的方法
	a, b = b, a
	fmt.Println(a, b)
	fmt.Printf("%p\n", &a)
	var p *int = &a
	fmt.Println(p)
	//指针初始化赋值,指针类型只保存地址，*ip是从地址获取指针存放的值
	var ip *int
	fmt.Println(ip)
	ip = &a
	fmt.Println(ip)
	*ip = 10
	fmt.Printf("指针内存地址%p,值%d\n", ip, *ip)

	struct11 := new(struct1)
	struct12 := new(struct2)
	struct12 = &struct2{}
	struct11 = &struct1{1, 2, &a, struct11, *struct12}
	readStruct(struct11)
}

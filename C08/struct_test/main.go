package main

import (
	"basic/model_pkg"
	"encoding/json"
	"fmt"
	"time"
	"unsafe"
)

//func say(a *Member){
//	fmt.Print(Member)
//}

type Class struct {
	Name string
	code int
}

type Student struct {
	Name string
	Class
}

func (c *Student) Info() {
	fmt.Println(c.Name)
	fmt.Printf("%s %d\n", c.Class.Name, c.Class.code)
}

//方法绑定结构体，且结构体和方法只能在同一个包中
func (c *Class) SetName(Name string) {
	c.Name = Name
}

func main() {

	//type
	//1.类型别名
	type myByte = byte
	type myByte1 = byte
	var b myByte
	fmt.Printf("%T\n", b)

	//2.定义一个新类型
	type myInt int
	var i myInt
	fmt.Printf("%T\n", i)

	//3.定义结构体

	//面向对象
	//基本特征:1.封装 2、继承 3、多态 4、方法重载 5、基类抽象
	var c model_pkg.Member = model_pkg.Member{
		Name: "django",
		//age:12,//age匿名变量
	}
	//指针
	class := &Class{"课堂1", 111}
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", *class)
	fmt.Println(class.Name)
	fmt.Println(unsafe.Sizeof("class"))
	//大小写 public private
	class.SetName("王牌")
	fmt.Println(class.Name)

	//继承
	student := &Student{"小黑", *class}
	student.Info()
	//student.Class.SetName("语文")
	student.SetName("语文")
	student.Info()

	//结构体标签
	type Goods struct {
		Id         int       `json:"id"`
		CreateTime time.Time `json:"createtime_at"`
	}
	goods := Goods{1, time.Now()}
	bytes, _ := json.Marshal(goods)
	fmt.Printf("%s", bytes)
}

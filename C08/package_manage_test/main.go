package main

import (
	"fmt"
	model_pkg2 "mygolangroad/C08/package_manage_test/model_pkg"
)

func main() {
	// 调用add方法
	fmt.Println(model_pkg2.Add(10, 25))
	// 调用 books结构体
	var threeKindom model_pkg2.Book
	threeKindom.Title = "threekindom"
	threeKindom.Author = "andy"
	// threekindom.bookid = 1003 这里不能定义，会报错，小写开头只供包内调用
}

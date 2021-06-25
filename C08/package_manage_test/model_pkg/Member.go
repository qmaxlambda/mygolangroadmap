package model_pkg

type Member struct {
	Name string
	age  int
}

func Add(x, y int) int { // 方法名Add首字母大写
	return x + y
}

type Book struct { // 方法名Books首字母大写
	Title  string // 对象首字母大写
	Author string
	bookid int // 首字母不大写不能被引用, 被调用声明后会有一个默认值，但是不能被再定义
}

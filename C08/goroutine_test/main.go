package main

import (
	"fmt"
	"time"
)

func Print1(a string) {
	i := 0
	for {
		fmt.Printf("%s:%d\n", a, i)
		i++
	}
}

func main() {
	//主死从随
	go Print1("a")
	go Print1("b")
	time.Sleep(time.Second * 2)
	fmt.Print("hw")
}

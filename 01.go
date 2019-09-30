package main

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func fun1() {
	fmt.Println("This is a ")
	// 格式化输出，会自动换行。
}

func test() {
	fmt.Println("This is function test!")
}

func fun2() {
	fmt.Printf("This is fun2")
	// 输出之后不会自动换行。
}

var ds int = 10

func variable1() {
	var s string = "This is string"
	// var s string = "This is"
	// f := "Runbool"
	fmt.Println(s)

	// fmt.Println(f)
	var a, b int = 12, 14
	fmt.Println(a, b)
}

func condition() {
	var a int = 15
	if a > 20 {
		fmt.Print("a:")
	} else {
		fmt.Print("a:", a)
	}
}

func looped() {
	for a := 1; a < 101; a++ {
		if a == 99 {
			continue
		}
		fmt.Println("This is loop", a)
	}
}

func switchde() {
	a := 100
	switch a {
	case 10:
		fmt.Println("a", 10)
	case 20:
		fmt.Println("a", 20)
	default:
		fmt.Println("100")
	}
}

func connectionwithmysql() {

}

func main() {
	// condition()
	// variable1()
	// switchde()
	looped()
}

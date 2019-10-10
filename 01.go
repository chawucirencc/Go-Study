package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	_"unsafe"
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
	/*
	循环...
	*/
	for a := 1; a < 11; a++ {
		fmt.Println("This is loop ", a)
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

func pointer()  {
	//  Go语言中的指针
	var a int = 10
	var ptr *int
	ptr = &a
	fmt.Println(ptr, *ptr)
}

func array()  {
	var a = [10] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[0:5]
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	fmt.Println(b)
	
}

func connectionwithmysql() {
	//  数据库连接...
	const(
		userName = "root"
		password = "password"
		ip = "127.0.0.1"
		port = "3306"
		dbName = "world"
	)

	db, err := sql.Open("mysql", "root:password@/world?charset=utf8")
	if (err != nil){
		fmt.Println("错误", err)
		return
	}
	c := db.Ping()
	if c != nil {
		fmt.Println("错误", err)
	}else{
		fmt.Println("连接成功")
	}
	rows, _ := db.Query("select * from dytt_new")
	// fmt.Println(rows)
	for rows.Next() {
		
	}
}

func factorial(n uint64) (result uint64) {
	if (n > 0){
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func fib(n uint64) uint64 {
	// 递归的性能太差。
	var a, b uint64 = 0, 1
	var i uint64
	for i=0; i < n; i++{
		a, b = b, a + b
		fmt.Println(a)
	}
	return 1
}

func main() {
	connectionwithmysql()
}

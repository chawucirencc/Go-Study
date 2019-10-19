package main

import (
	"database/sql"
	"fmt"
	_ "math"
	_ "unsafe"

	_ "github.com/go-sql-driver/mysql"
)

func helloworld() {
	// 文本输出，hello world
	fmt.Println("This is a text")
	fmt.Println("Hello World")
	// 格式化输出，会自动换行。
}

func variable() {
	// Go语言变量的使用。
	var a string = "SAMSUNG"
	fmt.Println(a)

	b := 12 //简易形式
	fmt.Println(b)

	var c, d int = 1, 2
	fmt.Println(c, ":", d)

	var e int //不主动赋值，默认值。
	var f string
	fmt.Println(e, ":", f)
}

func constantuse() {
	//常量的使用。
	const n = 5e8
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
}

func forloop() {
	//for循环使用。
	for i := 0; i < 20; i++ {
		if i == 15 {
			fmt.Println("i=", i)
			continue
		}
		fmt.Println("Number is ", i)
	}
}

func choseifelse() {
	// if...else...选择语句
	var i int = 21
	if i == 20 {
		fmt.Println("i =", i)
	} else {
		fmt.Println("i not", 20)
	}
}

func switchde() {
	// switch语句的使用。
	a := 100
	switch a {
	case 10:
		fmt.Println("a", 10)
	case 20:
		fmt.Println("a", 20)
	default: // 没有匹配的情况下输出default的内容。
		fmt.Println("100")
	}
}

func slicesfun() {
	// 切片的使用，创建非0长度的空切片，使用内置函数make(),len()返回切片的长度。
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println("s[0] =", s[0]) //取第一个元素
	fmt.Println("长度：", len(s))  // 获取长度
	fmt.Println("s[:2] =", s[:2])
}

func rangegofun() {
	// Go语言范围的使用。
	num := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := "This is a test text!"
	for _, i := range num { // range返回两个值，第一个为数组中元素的下标，第二个为数组中的元素。
		fmt.Println(i)
	}
	for _, j := range s { // 这种对于字符串的遍历方法称为Unicode遍历。
		fmt.Println(j)
	}
	for k := 0; k < len(s); k++ { // 这种遍历方法称为UTF-8遍历。
		fmt.Println(s[k])
	}
}

func functionuseadd(a int, b int) (int, int, int) { // 函数的使用，加入参数。加法函数。
	// 函数可以返回多个值。
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)
	return a + b, a, b // 必须有return
	// a, b , c := functionuseadd(100, 200)
	// fmt.Println(a, b, c)
}

func changefunctionuse(nums ...int) int { // ...int参数会变成一个可变数组，通过循环将数组内的数字相加。
	fmt.Println("nums", nums)
	total := 0
	for _, i := range nums {
		total += i
	}
	fmt.Println("total", total)
	return total
	// num := []int{1, 2, 3, 4, 5}
	// changefunctionuse(num...)
}

func closepackfunctionuse() func() int { // 匿名函数，闭包函数。
	// 会返回一个函数，每次调用都会让结果+1.
	// nextInt := closepackfunctionuse()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	i := 0
	return func() int {
		i++
		return i
	}
}

func fib(n uint64) uint64 { // 函数递归的使用。
	// 递归的性能太差。
	var a, b uint64 = 0, 1
	var i uint64
	for i = 0; i < n; i++ {
		a, b = b, a+b
		fmt.Println(a)
	}
	return 1
}

func array() {
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[0:5]
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	fmt.Println(b)
	for k, v := range a {
		fmt.Println(k, ":", v)
	}

}

func connectionwithmysql() {
	//  数据库连接...
	const (
		userName = "root"
		password = "password"
		ip       = "127.0.0.1"
		port     = "3306"
		dbName   = "world"
	)

	db, err := sql.Open("mysql", "root:password@/world?charset=utf8")
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	c := db.Ping()
	if c != nil {
		fmt.Println("错误", err)
	} else {
		fmt.Println("连接成功")
	}
	rows, _ := db.Query("select * from dytt_new")
	// fmt.Println(rows)
	for rows.Next() {

	}
}

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func findstruct() {
	type student struct {
		id     int
		name   string
		mobile int
	}
	var stuname student
	var a *student
	stuname.name = "Jonh"
	// var b int
	// b = 10
	// fmt.Println(&a)
	fmt.Println(a)
}

func main() {
	
}

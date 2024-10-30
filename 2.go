package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	fmt.Println("输入两个数字")
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Println("输入一个运算符")
	var str string
	fmt.Scan(&str)
	switch str {
	case "+":
		println(n1 + n2)
	case "-":
		println(n1 - n2)
	case "*":
		println(n1 * n2)
	case "/":
		println(n1 - n2)

	}

}

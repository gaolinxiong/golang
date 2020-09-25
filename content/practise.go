package main

import "fmt"

func main() {
	//变量的两种命名方式
	var a int32 = 123
	x := int32(321)

	//if else
	if x > 123 {
		fmt.Println(a)
	} else if x == 123 {
		print(1)
	} else {
		print(0)
	}

	switch false {
		case x < a:
			println("错的")
		case x > a:
			println("对的")
		default:
			println("我错了")

	}

	for i := 0; i < 5; i++ {
		println("i", i)
	}

	for a > 100 {
		println('c', a)
		a--
	}

	for {
		println("while", a)
		a++
		if a == 123 {
			break
		}
	}
	fmt.Println(a+x)
}

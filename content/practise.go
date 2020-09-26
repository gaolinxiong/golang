package main

import (
	"errors"
	"fmt"
)

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

	b := []int{100, 101, 102}

	for i, n :=range b {
		println(i, "i:", n)
	}
	fmt.Println(a+x)
	c, d := "10", 0
	e, f := div(c, d)

	fmt.Print(e, f)
	test(1, 8)
	sliceContent()
	structure()
}

func div(a string, b int) (string, error) {
	if b == 0 {
		return a, errors.New("division by zero")
	}

	return a, nil
}

func test(a, b int) {
	defer print("高林雄")

	defer print("喜欢")

	defer print("谁")

	println(a / b, "gaolinxiong")
}

//切片及map
func sliceContent() {
	x := make([]int, 0, 5)

	for i := 0; i < 11; i++ {
		x = append(x, i)
	}
	println(x)

	m := make(map[string]int)

	m["a"] = 1

	c, ok := m["b"]

	fmt.Println(c, ok)

	fmt.Println(m["a"])

	delete(m, "a")
}

type user struct {
	name string
	age byte
}

type manager struct {
	user
	title string
}

func structure()  {
	var m manager
	m.name = "Tom"
	m.age = 29
	m.title = "CTO"

	fmt.Println(m.user)

	var x X
	x.inc()
	println("func", x)

	var mo managerOne
	mo.name = "gaolinxiong"
	mo.age = 99
	mo.title = "很帅"

	println(mo.ToString())
}

type X int

func (x *X) inc()  {
	*x++
}

type userOne struct {
	name string
	age byte
}

func (u userOne) ToString() string {
	return fmt.Sprint("%+v", u)
}

type managerOne struct {
	userOne
	title string
}


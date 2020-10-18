package main

import (
	"errors"
	"fmt"
	"time"
)

//var x = 100
func main() {
	
	//x,_ := strconv.Atoi("123")
	//println(_)
	//变量的两种命名方式
	//var a int32 = 123
	//x := int32(321)
	//
	////if else
	//if x > 123 {
	//	fmt.Println(a)
	//} else if x == 123 {
	//	print(1)
	//} else {
	//	print(0)
	//}
	//
	//switch false {
	//	case x < a:
	//		println("错的")
	//	case x > a:
	//		println("对的")
	//	default:
	//		println("我错了")
	//
	//}
	//
	//for i := 0; i < 5; i++ {
	//	println("i", i)
	//}
	//
	//for a > 100 {
	//	println('c', a)
	//	a--
	//}
	//
	//for {
	//	println("while", a)
	//	a++
	//	if a == 123 {
	//		break
	//	}
	//}
	//
	//b := []int{100, 101, 102}
	//
	//for i, n :=range b {
	//	println(i, "i:", n)
	//}
	//fmt.Println(a+x)
	//c, d := "10", 0
	//e, f := div(c, d)
	//
	//fmt.Print(e, f)
	//test(1, 8)
	//sliceContent()
	//structure()
	//
	//var u userTwo
	//u.name = "Tom"
	//u.age = 20
	//
	//var p Printer = u
	//p.PrintUser()
	//
	//fmt.Println("-------------------")
	//go task(1)
	//
	//go task(2)
	//
	//time.Sleep(time.Second * 5)
	//
	//ch2 := make(chan string, 1)
	//// 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
	//// 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
	//// 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
	//// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
	//go func() {
	//	ch2 <- ("已到达")
	//}()
	//var value string = "数据"
	//value = value + (<-ch2)
	//fmt.Println(value)

	//fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>")
	//testFun()

	//x := 345
	//fmt.Println(&x)
	//const x = 123
	//println(&x)
	//
	//const y = 1.23
	//
	//{
	//	x = 99
	//	println(x)
	//}
	//const (
	//	x, y int = 99, 1
	//	b byte = byte(x)
	//	n = uint8(y)
	//)
	//
	//fmt.Println(b, y)
	//
	//s1 := "hello, world"
	//s2 := `hello, world.
	//send using Golang`
	//fmt.Println(s1, s2)

	//const (
	//	f = 1 << iota
	//	g
	//	h
	//)
	//fmt.Println(f, g, h)

	//start := time.Now()
	//ch := make(chan int, 20)
	//go testChan(ch)
	//for i := range ch {
	//	fmt.Println("接收到的数据:", i)
	//}
	//end := time.Now()
	//consume := end.Sub(start).Seconds()
	//fmt.Println("程序执行耗时(s)：", consume)

	//pi := 3.14159
	//// 按数值本身的格式输出
	//variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	//
	//fmt.Println(variant)
	// 匿名结构体声明, 并赋予初值
	//profile := &struct {
	//	Name string
	//	HP   int
	//}{
	//	Name: "rat",
	//	HP:   150,
	//}
	//
	//fmt.Printf("使用'%%+v' %+v\n", profile)
	//fmt.Printf("使用'%%#v' %#v\n", profile)
	//fmt.Printf("使用'%%T' %T\n", profile)
}

func testChan(ch chan<- int) {
	for i := 0; i < 100; i++ {
		chda <- i
	}
	close(ch)
}

func consumer(data chan int, done chan bool)  {
	println("data", len(data))
	println("done", done)
	for x := range data {
		println("recv", x)
	}

	done <- true
}

func producer(data chan int)  {
	for i:= 0; i < 4; i++ {
		data <- i
		time.Sleep(time.Second * 3)
	}

	close(data)
}

func testFun()  {
	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)

	go producer(data)

	<-done
}


func task(id int) {
	fmt.Println("测试", id)
	for i :=0; i < 10;i++ {
		fmt.Printf("%d: %d\n", id, i)
		time.Sleep(time.Second)
	}
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

type userTwo struct {
	name string
	age byte
}

func (u userTwo) PrintUser()  {
	fmt.Println("userTwo", u)
}

type Printer interface {
	PrintUser()
}



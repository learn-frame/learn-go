package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

// 变量除了可以在全局声明中初始化, 也可以在 init 函数中初始化. 这是一类非常特殊的函数, 它不能够被人为调用, 而是在每个包完成初始化后自动执行, 并且执行优先级比 main 函数高.
// init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine
func init() {
	var Pi = 4 * math.Atan(1) // init() function computes Pi
	println(Pi)
}

func main() {
	const PI = 3.14
	const start, end = 0, 1
	const (
		Monday, Tuesday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	// 关于 iota: https://juejin.cn/post/7087796114585780232
	const (
		a = iota
		b = iota
		c = iota
	)

	type Color int

	const (
		RED    Color = iota // 0
		ORANGE              // 1
		YELLOW              // 2
		GREEN               // ..
		BLUE
		INDIGO
		VIOLET // 6
	)

	// int8 int16 int32 int64
	// uint8 uint16 uint32 uint64
	// float32 float64, float32 精确到小数点后 7 位, float64 精确到小数点后 15 位
	var varA int
	var varB bool
	var varC string = ""
	println(varA, varB, varC)

	// 函数内的局部变量用简短声明语法
	someVar := 1
	println(someVar)

	e, f, g := 5, 7, "abc"
	println(e, f, g)

	// 交换两个值
	e, f = f, e

	fmt.Println("hello, world")
	fmt.Printf("hello, world %s %v %t", varC, varC, varB)

	type BitFlag int
	const (
		Active  BitFlag = 1 << iota // 1 << 0 == 1
		Send                        // 1 << 1 == 2
		Receive                     // 1 << 2 == 4
	)

	flag := Active | Send // == 3
	println(flag)

	rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / \n", 100*rand.Float32())
	}

	// 类型别名
	type TZ int

	println(strings.Replace("hello thank you thank you very much", "thank", "thx", 1))
	println(strings.Fields("hello thank you thank you very much"))
	println(strings.Split("hello thank you thank you very much", " "))

	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	str3 := strings.Join(sl, ";")
	fmt.Printf("sl joined by ;: %s\n", str3)

	// - `strconv.Itoa(i int) string` 返回数字 i 所表示的字符串类型的十进制数.
	//
	// - `strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string` 将 64 位浮点型的数字转换为字符串,
	// 其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'）, prec 表示精度, bitSize 则使用 32 表示 float32, 用 64 表示 float64.
	//
	// `strconv.Atoi(s string) (i int, err error)` 将字符串转换为 int 型.
	//
	// `strconv.ParseFloat(s string, bitSize int) (f float64, err error)` 将字符串转换为 float64 型.

	t := time.Now()
	fmt.Println(t.Format("02 Jan 2006 15:04"))

	// if initialization; condition {
	// 	// do something
	// }

	const max = 1
	val := 10
	if val > max {
		// do something
	}
	if val := 10; val > max {
		// do something
	}

	/*
		 	if err := file.Chmod(0664); err != nil {
				fmt.Println(err)
				return err
			}
	*/

	if t, ok := mySqrt(25.0); ok == false {
		fmt.Println(t)
	}

	// for 循环
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			println(j)
		}
	}

	// 类似于 while
	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}

	// 无限循环
	// for { }

	for pos, char := range "Go is a beautiful language!" {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}

	useGetX2AndX3()
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	} // error case

	return math.Sqrt(f), true
}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Stack [1]int

func (st *Stack) Pop() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}

	return 0
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

func useGetX2AndX3() {
	a, b := getX2AndX3(5)
	println(a, b) // 10, 15

	c, d := getX2AndX3_2(5)
	println(c, d) // 10, 15
}

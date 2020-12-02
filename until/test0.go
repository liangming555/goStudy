package until

import (
	"fmt"
	"math"
	"sort"
)

var e struct { //声明一个结构体变量,拥有一个整型的x字段
	x int
	y string
}

type name struct {
	x int
	y string
}

var d func() bool //声明一个返回值为布尔类型的函数变量,这种形式一般用于回调函数,即将函数以变量的形式保存下来,在需要的时候重新调用这个函数

func goDate() name {
	return name{x: 123, y: "21123"}
}

func TypeOfDate() {
	fmt.Println("Hello world!")
	var a int
	a = 123
	fmt.Println("a=", a)
	b := 123
	fmt.Println("b=", b)
	fmt.Println("maxInt8", math.MinInt8, math.MaxInt8) //-1<<7-1~1<<7-1
	var a1 int8
	a1 = 127
	fmt.Println(a1)
	fmt.Println(math.MaxInt64)
	a3 := 9223372036854775807
	fmt.Println(a3)
	fmt.Println(math.MaxUint8) //1<<8-1
	fmt.Println(math.MaxInt32)
	var f1 float32
	f1 = 0.1
	fmt.Println(f1, "float32 最大", math.MaxFloat32)
	str := "ewewew"
	fmt.Println("str=", str, "这是一个字符串", str+"123")
	var name []int
	name = append(name, 123, 789, 456)
	fmt.Println(name)
	fmt.Println("name的长度", len(name), name[0], name[len(name)-1])
	sort.Ints(name)
	fmt.Println(name)

	//var d func() bool//声明一个返回值为布尔类型的函数变量,这种形式一般用于回调函数,即将函数以变量的形式保存下来,在需要的时候重新调用这个函数
	//fmt.Println(d())

	if e.x == 0 {
		e.x = 1
	}
	fmt.Println(e.x, len(e.y))
	resX := goDate().x
	resY := goDate().y
	fmt.Println(resX, resY)
}

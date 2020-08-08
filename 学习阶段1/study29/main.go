package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	Name string
	Age  int
}

//反射获取任意变量的类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	// v.Name() //类型名称
	// v.Kind() //种类
	fmt.Printf("类型：%v  类型名称：%v 类型种类：%v\n", v, v.Name(), v.Kind())
	// fmt.Println(v)
}

func reflectValue(x interface{}) {
	// var num = 10 + x  //报错：mismatched types int and interface {}
	// b, _ := x.(int)
	// var num = 10 + b
	// fmt.Println(num)

	//反射来实现这个功能
	v := reflect.ValueOf(x)
	kind := v.Kind()
	// var m = v.Int() + 12
	// fmt.Println(m)
	switch kind {
	case reflect.Int:
		fmt.Printf("int类型的原始值%v\n", v.Int())
	case reflect.Float32:
		fmt.Printf("float32类型的原始值%v\n", v.Float())
	case reflect.Float64:
		fmt.Printf("float64类型的原始值%v\n", v.Float())
	case reflect.String:
		fmt.Printf("string类型的原始值%v\n", v.String())
	default:
		fmt.Printf("zzz")
	}

}

func main() {
	a := 10
	b := 23.4
	c := true
	d := "你好"
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)
	var e myInt = 34
	var f = Person{
		Name: "张三",
		Age:  20,
	}
	reflectFn(e) //main.myInt
	reflectFn(f) //main.Person

	var g = 24
	reflectFn(&g) //*int

	fmt.Println("----------")

	var h = 13
	reflectValue(h)
	var str = "nihao"
	reflectValue(str)
}

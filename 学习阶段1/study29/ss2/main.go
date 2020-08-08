package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄:%v 成绩:%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法")
}

//打印字段
func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传人的参数不是一个结构体")
		return
	}

	//1.通过类型变量里面的Filed可以获取结构体的字段
	filed0 := t.Field(0)
	fmt.Println(filed0)
	fmt.Println("字段名称", filed0.Name)
	fmt.Println("字段类型", filed0.Type)
	fmt.Println("字段Tag:", filed0.Tag.Get("json"))

	//2.通过类型变量里面的FieldbyName可以获取结构体的字段
	filed1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println("--字段名称：", filed1.Name)           //Age
		fmt.Println("字段类型", filed1.Type)              //int
		fmt.Println("字段Tag:", filed1.Tag.Get("json")) //age
	}

	//3.通过类型变量里面的NumField获取到该结构体有几个字段
	var fieldCount = t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性")

	//4.通过值变量获取结构体属性对应的值
	fmt.Println(v.FieldByName("Name")) //张三
}

func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传人的参数不是一个结构体")
		return
	}
	//1.通过类型变量里面的Method可以获取结构体的方法
	Method0 := t.Method(0) //和结构体方法的顺序没关系，和结构体方法的AscII有关系
	fmt.Println(Method0)   //{GetInfo  func(main.Student) string <func(main.Student) string Value> 0}

	//2.通过类型变量获取这个结构体有多少个方法
	Method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(Method1.Name)
		fmt.Println(Method1.Type)
	}

	//3.通过值变量执行方法（注意需要使用值变量，并且要注意参数）v.Method(0).Call(nil)或者v.MethodByName
	// v.Method(1).Call(nil)
	v.MethodByName("Print").Call(nil)
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)
	//4.执行方法传入参数（注意需要使用值变量名，并且要注意参数，接收的参数是[）reflect.Value的切片
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(20))
	params = append(params, reflect.ValueOf(90))
	v.MethodByName("SetInfo").Call(params) //执行方法，传入参数
	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

	//5.修改结构体属性

	//6.获取方法数量
	fmt.Println("方法数量", t.NumMethod())

}

func main() {
	s1 := Student{
		Name:  "张三",
		Age:   23,
		Score: 100,
	}
	// PrintStructField(s1)
	PrintStructFn(&s1)
}

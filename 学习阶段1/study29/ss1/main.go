package main

import (
	"fmt"
	"reflect"
)

func reflectSetvalue1(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())        //ptr
	fmt.Println(v.Elem().Kind()) //int64
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	}
}

func reflectSetValue2(x interface{}) {

}

func main() {
	var a int64 = 100
	reflectSetvalue1(&a)
	fmt.Println(a)
}

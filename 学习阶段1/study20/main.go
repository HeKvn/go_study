package main

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/tidwall/gjson"
)

//上面引入的这个包是解决float精度的
//使用go mod download
//使用go mod vendor命令将第三方包复制到项目文件夹，就可以免去第一次执行的时候去寻找包了，推荐
//(此项目已经go mod vendor 了，所以显示了那2个文件夹，不复制也行，但推荐这么做)

func main() {
	// var num1 float64 = 3.1
	// var num2 float64 = 4.2
	// fmt.Println(num1 + num2) //7.300000000000001

	//加法
	var num1 float64 = 3.1
	var num2 float64 = 4.2
	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	fmt.Println(d1)

	//剑法
	m1 := 8.2
	m2 := 3.8
	m3 := decimal.NewFromFloat(m1).Sub(decimal.NewFromFloat(m2))
	fmt.Println(m3)
	//乘除 Mul Div 与上类似

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())
}

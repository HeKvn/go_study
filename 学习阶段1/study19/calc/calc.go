package calc

var aaa = "私有变量"
var AAA = "公有变量"

func Add(x, y int) int { //首字母大写，公有方法
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

package main

import "fmt"

// 常量是一个包级别的声明，这里有一个小知识点：包级别的声明，不仅仅包含其声明的源文件可见，对于同一个包内的文件，也是可见的。
const boilingF = 212.0

func main() {
	// f 和 c 属于 main函数的局部变量
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("%gf or %g c", f, c)
	write()
}

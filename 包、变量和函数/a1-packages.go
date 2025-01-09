package main

// 如果一个名字以大写开头，就是已导出的，在导入一个包时，只能引用已经导出的名字


import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("我最喜欢的数字是",rand.Intn(10))
	fmt.Println("现在你有了 %g 个问题。\n",math.Sqrt(7))
	fmt.Println(math.Pi)
	
}
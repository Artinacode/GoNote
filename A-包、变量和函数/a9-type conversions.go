package main

import (
	"fmt"
	"math"
)
//go中不同类型的项之间的赋值需要显示转换。
func main() {
	var x,y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x,y,z)
}
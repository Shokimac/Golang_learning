package main

import "fmt"

// 浮動小数点型

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 暗黙的な定義の場合は、自動でfloat64で定義される（環境依存ではない
	fl := 3.2
	fmt.Println(fl64 + fl)

	fmt.Printf("%T, %T\n", fl64, fl)

	// 基本的にはfloat32は使わない
	var fl32 float32 = 1.2

	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	// ゼロで割るとプラス or マイナスの無限大や
	// ゼロをゼロで割ると NaN(Not a Number)という非数を保持する
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

}

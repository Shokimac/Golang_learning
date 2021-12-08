package main

import (
	"fmt"
	"strconv"
)

// 型変換

func main() {
	// // int から floatへ
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// // float から intへ
	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3)
	// fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	// 文字列を数値へ
	// strconvのAtoiは2つ変数を返すが、今回使うのは1つのみ
	// 定義した変数は使わないといけないのがGoのルールだが
	// アンダーバーに格納させることで、その値は使わないことを明示的にできる
	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	// 数値を文字列へ
	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	// 文字列とバイト配列の変換
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	// バイト配列を文字列へ
	h2 := string(b)
	fmt.Println(h2)
}

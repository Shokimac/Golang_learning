package main

import "fmt"

// byte(uint8)型

func main() {
	// 鍵括弧を先頭につけることで配列を定義できる
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// 配列を渡すことで文字列を表示できる
	fmt.Println(string(byteA))

	// 文字列をバイトの配列に変換できる
	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))
}

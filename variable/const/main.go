package main

import "fmt"

// 定数

// 頭文字を大文字にすると他のパッケージからも呼び出せる
const Pi = 3.14

const (
	URL      = "http://xxx.com"
	SiteName = "test"
)

// 値を省略して定義
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	// iota で連番を代入することができる
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(c0, c1, c2)
}

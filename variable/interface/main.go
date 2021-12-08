package main

import "fmt"

func main() {
	// interface型は{}まで含めて型の名前になる
	// 他の型との互換性がある
	// 初期値は nil
	// 演算の対象としては使えない
	// 全ての型を表すことができる
	var x interface{}
	fmt.Println(x)

	// 数値も入る
	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	x = 2
	fmt.Println(x + 3)
}

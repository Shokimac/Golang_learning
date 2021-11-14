package main

import "fmt"

// 型
// 整数型

func main() {
	var i int = 100

	/*
		int のみの指定だと環境依存となる（32bitなら int32となる
		種類は以下の通り
		int8
		int16
		int32
		int64
	*/

	//  bit指定も可能
	var i2 int64 = 200

	fmt.Println(i + 50)

	// int指定の変数とint64指定の変数で計算は出来ない（「invalid operation: i + i2 (mismatched types int and int64)」表示
	// 環境依存でint64となった変数と明示的にint64で定義した変数は全く別物扱いとなる
	// fmt.Println(i + i2)

	// fmtのPrintfで %T とすることで型表記ができる
	fmt.Printf("%T\n", i2)

	// 型変換について
	// int32()で変数を指定すると型変換ができる
	fmt.Printf("%T\n", int32(i2))

	// intのみで定義された i も型変換することで明示的にint64定義されたi2と計算ができるようになる
	fmt.Println(i2 + int64(i))
}

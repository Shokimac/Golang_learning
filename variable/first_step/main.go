package main

import "fmt"

// 変数について

// 関数外では暗黙的な定義は出来ない
// i5  := 500

// 明示的な定義は関数外でも出来る
var i5 int = 500

// 基本的には明示的な定義を使う（型指定を行ってバグを防ぐ
// 他者が見た時に明示的な方が型を把握しやすい

// Goでは定義した変数は必ずプログラム内で使用しなければいけない決まり（破るとエラー発生

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// 変数定義には、明示的な定義と暗黙的な定義がある

	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	// 変数名をカンマで区切ると複数定義できる
	var t, f bool = true, false
	fmt.Println(t, f)

	// 丸括弧でも複数定義を行える
	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 値を入れずに変数のみ定義する場合
	var i3 int    // 初期値で0が入る
	var s3 string // 初期値で空文字が入る
	fmt.Println(i3, s3)

	// 宣言済みの変数に値を格納
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義 = 型指定の必要がない
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 暗黙的な定義は、再定義できない
	// i4 := 500

	// int型で定義されているため、String型も入れられない
	// i4 = "Hello"

	fmt.Println(i5)

	outer()

	// 定義した変数はその関数内でのみ使用可能
	// fmt.Println(s4)

	// 実行すると「s5 declared but not used」とエラーが出る
	var s5 string = "Not Use"
	fmt.Println(s5)
}

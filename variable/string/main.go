package main

import "fmt"

func main() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	// バッククォーテーションで囲むとそのまま出力できる
	fmt.Println(`test
	test
		test`)

	//　エスケープはバックスラッシュ
	fmt.Println("\"")
	// バッククォーテーションでも表示可能
	fmt.Println(`"`)

	// バイト型に対して下記の表記で1文字目を取得できる
	fmt.Println(string(s[0]))
}

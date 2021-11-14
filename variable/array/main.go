package main

import "fmt"

// 配列型

// 後から要素数を追加、削除が出来ない(Goの特徴)

func main() {
	// 初期値が入る
	var arr1 [3]int
	fmt.Println(arr1)

	// [3]int = Goでは要素数も含めて型指定となるため、要素数を変更することが出来ない
	fmt.Printf("%T\n", arr1)

	// 3つ目には空文字がデフォルトで入る
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	// 暗黙的な配列定義（要素数指定
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 暗黙的な配列定義（要素数未指定
	arr4 := [...]string{"A", "B"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// 配列操作について
	// 値の取り出し
	fmt.Println(arr2[0])

	// 値の更新 追加
	arr2[2] = "C"
	fmt.Println(arr2)

	// 要素数が異なる配列同士では割り当てられずエラー
	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 要素数の出力方法
	fmt.Println(len(arr1))
}

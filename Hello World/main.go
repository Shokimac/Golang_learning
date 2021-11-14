package main

// 複数の場合は丸括弧でまとめることもできる
import (
	"fmt"
	"time"
)

// Hello World

// パッケージの作成→関数の作成→必要ならパッケージのインポート
// パッケージの宣言は必ず1つ

/*
複数行のコメントはこのように書く
*/

func main() {
	// 渡したデータを改行付きで出力する関数
	fmt.Println("Hello World")
	// 時間の表示
	fmt.Println(time.Now())
}

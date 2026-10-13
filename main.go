package main

import (
	"fmt"
)

type Player struct {
	Checker [24]int
	Bar     int
	Off     int
}

type Board struct {
	Black Player
	White Player
}

func main() {
	//チェッカー初期配置
	initCheckers := [24]int{0, 0, 0, 0, 0, 5, 0, 3, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}

	//黒と白の各プレイヤーを定義
	black := Player{
		Checker: initCheckers,
		Bar:     0,
		Off:     0,
	}
	white := Player{
		Checker: initCheckers,
		Bar:     0,
		Off:     0,
	}

	//ボードを定義
	board := Board{black, white}
	board.printBoard()

	fmt.Println("\n黒のチェッカームーブです。ダイス目：3 1, ムーブ：8/5 6/5")
	board.moveBlack(3, 8)
	board.moveBlack(1, 6)
	board.printBoard()

	fmt.Println("\n白のチェッカームーブです。ダイス目：6 5, ムーブ：24/13")
	board.moveWhite(6, 24)
	board.moveWhite(5, 18)
	board.printBoard()
	moves, dice := diceRoll()

	fmt.Println("\nここからはダイスロールして動かせる出目を取り出す処理を行います。")

	fmt.Println("出目：", dice[0], dice[1])

	fmt.Println("\n動かせる数を表示します。")
	for _, move := range moves {
		fmt.Println(move)
	}

	//fmt.Println("終了するにはEnterキーを押してください...")
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func (b *Board) printBoard() {
	// 盤面表示

	fmt.Println("ボード表示")
	fmt.Println(*b)

	fmt.Println("ここから個別表示")

	fmt.Println("黒のチェッカー")
	fmt.Println(b.Black.Checker)
	fmt.Println("黒のオンザバー")
	fmt.Println(b.Black.Bar)
	fmt.Println("黒のベアオフ")
	fmt.Println(b.Black.Off)

	fmt.Println("白のチェッカー")
	fmt.Println(b.White.Checker)
	fmt.Println("白のオンザバー")
	fmt.Println(b.White.Bar)
	fmt.Println("白のベアオフ")
	fmt.Println(b.White.Off)
}

// チェッカームーブ(エンター/ベアオフは別処理)
func (boa *Board) moveBlack(dice int, s int) {
	to := (s - 1) - dice
	boa.Black.Checker[s-1]--
	boa.Black.Checker[to]++
}

func (boa *Board) moveWhite(dice int, s int) {
	to := (s - 1) - dice
	boa.White.Checker[s-1]--
	boa.White.Checker[to]++
}

// 合法手実装はこれから

package main

import (
	"crypto/rand"
	"math/big"
)

func diceRoll() ([]int, [2]int) {
	dice := [2]int{}
	mov := []int{}

	for i := range dice {
		n, err := rand.Int(rand.Reader, big.NewInt(6))
		if err != nil {
			panic(err)
		}
		dice[i] = int(n.Int64()) + 1
	}

	// ゾロ目の場合は4回ムーブ
	if dice[0] == dice[1] {
		mov = append(mov, dice[0], dice[0], dice[0], dice[0])

		// ダイス目の順番を降順に並べ替える
	} else if dice[0] < dice[1] {
		dice[0], dice[1] = dice[1], dice[0]
		mov = append(mov, dice[0], dice[1])
	} else {
		mov = append(mov, dice[0], dice[1])
	}
	return mov, dice
}

package main

import (
	"math/rand"
	"time"
)

func main() {
	// 現在時刻を数値で取得する
	t := time.Now().UnixNano()
	// 乱数のたねを設定
	rand.Seed(t)
	// xは0-10の間の値になる
	s := rand.Intn(6)
	// println(s)
	switch s + 1 {
		case 6:
			println("大吉")
		case 5, 4:
			println("中吉")
		case 3, 2:
			println("吉")
		case 1:
			println("凶")
	}
}

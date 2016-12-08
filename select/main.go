package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	//ch1から受信した整数を2バイしてch2へ送信
	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()
	// ch2から受信した整数を1減算してch3へ送信
	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()

	n := 1
LOOP:
	for {
		select { //selectは以下のcaseをランダムに評価する
		// 整数を増分させつつch1へ送信
		case ch1 <- n:
			n++
			// ch3から受信した整数を出力
		case i := <-ch3:
			fmt.Println("recived", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}

}

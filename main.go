package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	normalResult := make(chan string)
	endResult := make(chan string)

	go call(normalResult, endResult)

	for {
		// どちらかを受信するまで待機する
		select {
		case result := <-normalResult:
			fmt.Println("Roop " + result)

		case <-endResult:
			fmt.Println("Roop end!")
			return
		}
	}
}

func call(normalResult chan string, endResult chan string) {
	var word string

	for {
		// 文字入力を読み取る
		fmt.Scan(&word)

		if word == "end" {
			endResult <- word
			break
		} else {
			normalResult <- word
		}
	}
}

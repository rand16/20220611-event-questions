package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		str := ""

		// ここから記述
		if i%15 == 0 {
			// 15の倍数であれば FizzBuzz に変換
			str = "FizzBuzz"
		} else if i%5 == 0 {
			// 5の倍数であれば Buzz に変換
			str = "Buzz"
		} else if i%3 == 0 {
			// 3の倍数であれば Fizz に変換
			str = "Fizz"
		} else {
			// 3, 5, 15 の倍数でなければ数値型を文字列型に変換
			str = strconv.Itoa(i)
		}
		// ここまで記述

		fmt.Println(str)
	}
}

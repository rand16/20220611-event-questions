package main

import "fmt"

func main() {
	// ランダムに並べられた重複のない整数の配列
	array := []int{5, 4, 6, 2, 1, 9, 8, 3, 7, 10}
	// ソート実行
	sortedArray := sort(array)
	// 結果出力
	for _, v := range sortedArray {
		fmt.Println(v)
	}
}

func sort(array []int) []int {
	// 要素が一つの場合はソートの必要がないので、そのまま返却
	if len(array) == 1 {
		return array
	}

	// 配列の先頭を基準値とする
	pivot := array[0]

	// ここから記述

	// 最終的なソートの結果を格納する配列
	var ans []int

	// 再帰的にソートを行う関数を定義
	var sortRecursion func(array []int) []int
	sortRecursion = func(array []int) []int {
		var emp []int
		if len(array) == 1 {
			flag := 0
			// 対象の配列の要素数が１で，かつ結果を格納する配列にこの要素が格納されていなければ，配列に格納する
			for _, v := range ans {
				if array[0] == v {
					flag = 1
					break
				}
			}
			if flag == 0 {
				ans = append(ans, array[0])
			}
			return emp
		}

		pivot = array[0]
		arrayLen := len(array)
		// 指定のアルゴリズムに則ってソート
		for start, end := 0, arrayLen-1; start < end; {
			if (array[start] >= pivot) && (array[end] < pivot) {
				val := array[start]
				array[start] = array[end]
				array[end] = val
				start++
				end--
			} else if array[start] >= pivot {
				end--
			} else if array[end] < pivot {
				start++
			} else {
				start++
				end--
			}
		}

		flag2 := 0
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				flag2 = 1
			}
		}
		// 配列 array が完全にソートされていれば結果の配列に格納
		if flag2 == 0 {
			for _, v := range array {
				ans = append(ans, v)
			}
		}

		var small, large []int
		// 完全にソートされていない場合，ピボットより小さい側と大きい側で配列を分離
		for _, val := range array {
			if val < pivot {
				small = append(small, val)
			} else {
				large = append(large, val)
			}
		}

		// 完全なソートが完了するまで sortRecursion() 関数を実行
		if len(small) >= 1 {
			sortRecursion(small)
		} else {
			return emp
		}
		if len(large) >= 1 {
			sortRecursion(large)
		} else {
			return emp
		}

		return ans
	}
	return sortRecursion(array)

	// ここまで記述

}

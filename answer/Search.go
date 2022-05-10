package main

import "fmt"

func main() {
	// 昇順にソートされた配列
	sortedArray := []int{1, 2, 3, 5, 12, 7890, 12345}
	// 探索対象の番号
	targetNumber := 7890
	// 探索実行
	targetIndex := serchIndex(sortedArray, targetNumber)
	// 結果出力
	fmt.Println(targetIndex)
}

func serchIndex(sortedArray []int, targetNumber int) int {

	// ここから記述

	// 探索対象の配列の長さを取得
	sortedArrayLen := len(sortedArray)
	// 配列の探索範囲の視点と終点のindexを設定
	start := 0
	end := sortedArrayLen - 1

	// 指定のアルゴリズムに則って探索
	for count := sortedArrayLen; count > 0; {
		count = count / 2
		// 配列の中間の値
		target := (start + end) / 2
		if targetNumber < sortedArray[target] {
			// 終点をずらす
			end = end - count - 1
		} else if targetNumber > sortedArray[target] {
			if (start+end)%2 == 1 {
				// 配列の探索範囲に存在する要素数が偶数個の場合の始点のずらし
				start = start + count
			} else {
				// 配列の探索範囲に存在する要素数が奇数個の場合の始点のずらし
				start = start + count + 1
			}
		} else {
			return target
		}
	}

	// ここまで記述

	// 探索対象が存在しない場合、-1を返却
	return -1
}

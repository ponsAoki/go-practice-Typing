package main

import (
	"fmt"
)

func main() {
	totalScore := 0
	ask(1, "ディレクトリ (フォルダ) 移動", "cd", &totalScore)
	ask(2, "フォルダの内容をリスト形式で表示", "ls", &totalScore)
	ask(3, "ファイルを移動", "mv", &totalScore)
	ask(4, "ファイルやフォルダをコピーする", "cp", &totalScore)
	ask(5, "フォルダを作成する", "mkdir", &totalScore)
	ask(6, "ファイルを作成する", "touch", &totalScore)
	ask(7, "空のフォルダを削除する", "rmdir", &totalScore)
	ask(8, "ファイルをフォルダを削除する", "rm", &totalScore)
	ask(9, "現在いる位置を表示する", "pwd", &totalScore)
	ask(10, "ファイルの検索をする", "find", &totalScore)
	// ask(1, "ディレクトリ (フォルダ) 移動", "cd", &totalScore)
	// ask(1, "ディレクトリ (フォルダ) 移動", "cd", &totalScore)
	// ask(1, "ディレクトリ (フォルダ) 移動", "cd", &totalScore)

	fmt.Println("スコア", totalScore)
	if totalScore <= 30 {
		fmt.Println("頑張りましょう!")
	} else if totalScore <= 60 && totalScore >= 31 {
		fmt.Println("もう少しです!")
	} else if totalScore <= 80 && totalScore >= 61 {
		fmt.Println("なかなかいいですね!")
	} else if totalScore >= 81 {
		fmt.Println("素晴らしい!!")
	}
}

func ask(number int, question string, answer string, scorePtr *int) {
	var ans string
	fmt.Printf("[問%d] %s\n", number, question)
	fmt.Scan(&ans)
	if answer == ans {
		fmt.Println("正解!!")
		*scorePtr += 10
	} else {
		fmt.Println("不正解!!")
	}
}

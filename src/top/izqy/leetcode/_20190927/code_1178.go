package _20190927

/* 1178. 猜字谜 （为解决，超时） */
func findNumOfValidWords(words []string, puzzles []string) []int {
	// 使用位标记的方式设置字符串中包含的字符
	wordsX := make([]uint32, len(words))
	for i, word := range words {
		for _, ch := range word {
			wordsX[i] = wordsX[i] | 1<<(uint32)(ch-'a')
		}
		//fmt.Printf("%b\n", wordsX[i])
	}
	puzzlesX := make([]uint32, len(puzzles))
	for i, puzzle := range puzzles {
		for _, ch := range puzzle {
			puzzlesX[i] = puzzlesX[i] | 1<<(uint32)(ch-'a')
		}
		//fmt.Printf("%b\n", puzzlesX[i])
	}
	answer := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		answer[i] = 0
		for j, word := range words {
			if isMatch(word, puzzle, wordsX[j], puzzlesX[i]) {
				answer[i]++
			}
		}
	}
	return answer
}

func isMatch(word, puzzle string, wordX, puzzleX uint32) bool {
	if (1<<(uint32)(puzzle[0]-'a'))&wordX == 0 {
		return false
	}
	if wordX&puzzleX|puzzleX == 0 {
		return false
	}
	//for _, ch := range word {
	//	if (1<<(uint32)(ch-'a'))&puzzleX == 0 {
	//		return false
	//	}
	//}
	return true
}
